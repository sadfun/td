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

// MessagesBotCallbackAnswer represents TL type `messages.botCallbackAnswer#36585ea4`.
type MessagesBotCallbackAnswer struct {
	// Flags field of MessagesBotCallbackAnswer.
	Flags bin.Fields
	// Alert field of MessagesBotCallbackAnswer.
	Alert bool
	// HasURL field of MessagesBotCallbackAnswer.
	HasURL bool
	// NativeUI field of MessagesBotCallbackAnswer.
	NativeUI bool
	// Message field of MessagesBotCallbackAnswer.
	//
	// Use SetMessage and GetMessage helpers.
	Message string
	// URL field of MessagesBotCallbackAnswer.
	//
	// Use SetURL and GetURL helpers.
	URL string
	// CacheTime field of MessagesBotCallbackAnswer.
	CacheTime int
}

// MessagesBotCallbackAnswerTypeID is TL type id of MessagesBotCallbackAnswer.
const MessagesBotCallbackAnswerTypeID = 0x36585ea4

// Encode implements bin.Encoder.
func (b *MessagesBotCallbackAnswer) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode messages.botCallbackAnswer#36585ea4 as nil")
	}
	buf.PutID(MessagesBotCallbackAnswerTypeID)
	if err := b.Flags.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode messages.botCallbackAnswer#36585ea4: field flags: %w", err)
	}
	if b.Flags.Has(0) {
		buf.PutString(b.Message)
	}
	if b.Flags.Has(2) {
		buf.PutString(b.URL)
	}
	buf.PutInt(b.CacheTime)
	return nil
}

// SetAlert sets value of Alert conditional field.
func (b *MessagesBotCallbackAnswer) SetAlert(value bool) {
	if value {
		b.Flags.Set(1)
	} else {
		b.Flags.Unset(1)
	}
}

// SetHasURL sets value of HasURL conditional field.
func (b *MessagesBotCallbackAnswer) SetHasURL(value bool) {
	if value {
		b.Flags.Set(3)
	} else {
		b.Flags.Unset(3)
	}
}

// SetNativeUI sets value of NativeUI conditional field.
func (b *MessagesBotCallbackAnswer) SetNativeUI(value bool) {
	if value {
		b.Flags.Set(4)
	} else {
		b.Flags.Unset(4)
	}
}

// SetMessage sets value of Message conditional field.
func (b *MessagesBotCallbackAnswer) SetMessage(value string) {
	b.Flags.Set(0)
	b.Message = value
}

// GetMessage returns value of Message conditional field and
// boolean which is true if field was set.
func (b *MessagesBotCallbackAnswer) GetMessage() (value string, ok bool) {
	if !b.Flags.Has(0) {
		return value, false
	}
	return b.Message, true
}

// SetURL sets value of URL conditional field.
func (b *MessagesBotCallbackAnswer) SetURL(value string) {
	b.Flags.Set(2)
	b.URL = value
}

// GetURL returns value of URL conditional field and
// boolean which is true if field was set.
func (b *MessagesBotCallbackAnswer) GetURL() (value string, ok bool) {
	if !b.Flags.Has(2) {
		return value, false
	}
	return b.URL, true
}

// Decode implements bin.Decoder.
func (b *MessagesBotCallbackAnswer) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode messages.botCallbackAnswer#36585ea4 to nil")
	}
	if err := buf.ConsumeID(MessagesBotCallbackAnswerTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.botCallbackAnswer#36585ea4: %w", err)
	}
	{
		if err := b.Flags.Decode(buf); err != nil {
			return fmt.Errorf("unable to decode messages.botCallbackAnswer#36585ea4: field flags: %w", err)
		}
	}
	b.Alert = b.Flags.Has(1)
	b.HasURL = b.Flags.Has(3)
	b.NativeUI = b.Flags.Has(4)
	if b.Flags.Has(0) {
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.botCallbackAnswer#36585ea4: field message: %w", err)
		}
		b.Message = value
	}
	if b.Flags.Has(2) {
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.botCallbackAnswer#36585ea4: field url: %w", err)
		}
		b.URL = value
	}
	{
		value, err := buf.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.botCallbackAnswer#36585ea4: field cache_time: %w", err)
		}
		b.CacheTime = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesBotCallbackAnswer.
var (
	_ bin.Encoder = &MessagesBotCallbackAnswer{}
	_ bin.Decoder = &MessagesBotCallbackAnswer{}
)
