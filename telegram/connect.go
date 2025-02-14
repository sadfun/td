package telegram

import (
	"context"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/go-faster/errors"
	"go.uber.org/multierr"
	"go.uber.org/zap"

	"github.com/gotd/td/internal/exchange"
	"github.com/gotd/td/internal/tdsync"
	"github.com/gotd/td/telegram/auth"
)

func (c *Client) runUntilRestart(ctx context.Context) error {
	g := tdsync.NewCancellableGroup(ctx)
	g.Go(c.conn.Run)

	// If we don't need updates, so there is no reason to subscribe for it.
	if !c.noUpdatesMode {
		g.Go(func(ctx context.Context) error {
			// Call method which requires authorization, to subscribe for updates.
			// See https://core.telegram.org/api/updates#subscribing-to-updates
			c.log.Debug("Calling c.Self to subscribe for updates")
			self, err := c.Self(ctx)
			if err != nil {
				// Ignore unauthorized errors.
				if !auth.IsUnauthorized(err) {
					c.log.Warn("Got error on self", zap.Error(err))
				}
				return nil
			}

			c.log.Info("Got self", zap.String("username", self.Username))
			return nil
		})
	}

	g.Go(func(ctx context.Context) error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-c.restart:
			c.log.Debug("Restart triggered")
			// Should call cancel() to cancel group.
			g.Cancel()

			return nil
		}
	})

	return g.Wait()
}

func (c *Client) isPermanentError(err error) bool {
	return errors.Is(err, exchange.ErrKeyFingerprintNotFound)
}

func (c *Client) reconnectUntilClosed(ctx context.Context) error {
	b := tdsync.SyncBackoff(backoff.WithContext(c.connBackoff(), ctx))

	return backoff.RetryNotify(func() error {
		if err := c.runUntilRestart(ctx); err != nil {
			if ctxErr := ctx.Err(); ctxErr != nil {
				c.log.Error("Stopping reconnection attempts due to parent context error",
					zap.Error(err),
					zap.Errors("error_parent", []error{ctxErr}),
				)
				return errors.Wrap(err, "parent context closed")
			}
			if c.isPermanentError(err) {
				return backoff.Permanent(err)
			}
			return errors.Wrap(err, "runUntilRestart")
		}

		return nil
	}, b, func(err error, timeout time.Duration) {
		c.log.Info("Restarting connection",
			zap.Error(err),
			zap.Duration("backoff", timeout),
		)

		c.connMux.Lock()
		c.conn = c.createPrimaryConn(nil)
		c.connMux.Unlock()
	})
}

func (c *Client) onReady() {
	c.log.Debug("Ready")
	c.ready.Signal()
}

func (c *Client) resetReady() {
	c.ready.Reset()
}

// Run starts client session and blocks until connection close.
// The f callback is called on successful session initialization and Run
// will return on f() result.
//
// Context of callback will be canceled if fatal error is detected.
// The ctx is used for background operations like updates handling or pools.
//
// See `examples/bg-run` and `contrib/gb` package for classic approach without
// explicit callback, with Connect and defer close().
func (c *Client) Run(ctx context.Context, f func(ctx context.Context) error) (err error) {
	if c.ctx != nil {
		select {
		case <-c.ctx.Done():
			return errors.Wrap(c.ctx.Err(), "client already closed")
		default:
		}
	}

	// Setting up client context for background operations like updates
	// handling or pool creation.
	c.ctx, c.cancel = context.WithCancel(ctx)

	c.log.Info("Starting")
	defer c.log.Info("Closed")
	// Cancel client on exit.
	defer c.cancel()
	defer func() {
		c.subConnsMux.Lock()
		defer c.subConnsMux.Unlock()

		for _, conn := range c.subConns {
			if closeErr := conn.Close(); !errors.Is(closeErr, context.Canceled) {
				multierr.AppendInto(&err, closeErr)
			}
		}
	}()

	c.resetReady()
	if err := c.restoreConnection(ctx); err != nil {
		return err
	}

	g := tdsync.NewCancellableGroup(ctx)
	g.Go(c.reconnectUntilClosed)
	g.Go(func(ctx context.Context) error {
		select {
		case <-ctx.Done():
			c.cancel()
			return ctx.Err()
		case <-c.ctx.Done():
			return c.ctx.Err()
		}
	})
	g.Go(func(ctx context.Context) error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-c.ready.Ready():
			if err := f(ctx); err != nil {
				return errors.Wrap(err, "callback")
			}
			// Should call cancel() to cancel ctx.
			// This will terminate c.conn.Run().
			c.log.Debug("Callback returned, stopping")
			g.Cancel()
			return nil
		}
	})
	if err := g.Wait(); !errors.Is(err, context.Canceled) {
		return err
	}

	return nil
}
