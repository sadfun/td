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

// Error represents TL type `error#c4b9f9bb`.
type Error struct {
	// Code field of Error.
	Code int
	// Text field of Error.
	Text string
}

// ErrorTypeID is TL type id of Error.
const ErrorTypeID = 0xc4b9f9bb

// Encode implements bin.Encoder.
func (e *Error) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode error#c4b9f9bb as nil")
	}
	b.PutID(ErrorTypeID)
	b.PutInt(e.Code)
	b.PutString(e.Text)
	return nil
}

// Decode implements bin.Decoder.
func (e *Error) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode error#c4b9f9bb to nil")
	}
	if err := b.ConsumeID(ErrorTypeID); err != nil {
		return fmt.Errorf("unable to decode error#c4b9f9bb: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode error#c4b9f9bb: field code: %w", err)
		}
		e.Code = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode error#c4b9f9bb: field text: %w", err)
		}
		e.Text = value
	}
	return nil
}

// Ensuring interfaces in compile-time for Error.
var (
	_ bin.Encoder = &Error{}
	_ bin.Decoder = &Error{}
)
