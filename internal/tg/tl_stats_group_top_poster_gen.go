// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"fmt"

	"github.com/ernado/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)

// StatsGroupTopPoster represents TL type `statsGroupTopPoster#18f3d0f7`.
type StatsGroupTopPoster struct {
	// UserID field of StatsGroupTopPoster.
	UserID int
	// Messages field of StatsGroupTopPoster.
	Messages int
	// AvgChars field of StatsGroupTopPoster.
	AvgChars int
}

// StatsGroupTopPosterTypeID is TL type id of StatsGroupTopPoster.
const StatsGroupTopPosterTypeID = 0x18f3d0f7

// Encode implements bin.Encoder.
func (s *StatsGroupTopPoster) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode statsGroupTopPoster#18f3d0f7 as nil")
	}
	b.PutID(StatsGroupTopPosterTypeID)
	b.PutInt(s.UserID)
	b.PutInt(s.Messages)
	b.PutInt(s.AvgChars)
	return nil
}

// Decode implements bin.Decoder.
func (s *StatsGroupTopPoster) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode statsGroupTopPoster#18f3d0f7 to nil")
	}
	if err := b.ConsumeID(StatsGroupTopPosterTypeID); err != nil {
		return fmt.Errorf("unable to decode statsGroupTopPoster#18f3d0f7: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode statsGroupTopPoster#18f3d0f7: field user_id: %w", err)
		}
		s.UserID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode statsGroupTopPoster#18f3d0f7: field messages: %w", err)
		}
		s.Messages = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode statsGroupTopPoster#18f3d0f7: field avg_chars: %w", err)
		}
		s.AvgChars = value
	}
	return nil
}

// Ensuring interfaces in compile-time for StatsGroupTopPoster.
var (
	_ bin.Encoder = &StatsGroupTopPoster{}
	_ bin.Decoder = &StatsGroupTopPoster{}
)
