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

// EmojiKeywordsDifference represents TL type `emojiKeywordsDifference#5cc761bd`.
type EmojiKeywordsDifference struct {
	// LangCode field of EmojiKeywordsDifference.
	LangCode string
	// FromVersion field of EmojiKeywordsDifference.
	FromVersion int
	// Version field of EmojiKeywordsDifference.
	Version int
	// Keywords field of EmojiKeywordsDifference.
	Keywords []EmojiKeywordClass
}

// EmojiKeywordsDifferenceTypeID is TL type id of EmojiKeywordsDifference.
const EmojiKeywordsDifferenceTypeID = 0x5cc761bd

// Encode implements bin.Encoder.
func (e *EmojiKeywordsDifference) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode emojiKeywordsDifference#5cc761bd as nil")
	}
	b.PutID(EmojiKeywordsDifferenceTypeID)
	b.PutString(e.LangCode)
	b.PutInt(e.FromVersion)
	b.PutInt(e.Version)
	b.PutVectorHeader(len(e.Keywords))
	for idx, v := range e.Keywords {
		if v == nil {
			return fmt.Errorf("unable to encode emojiKeywordsDifference#5cc761bd: field keywords element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode emojiKeywordsDifference#5cc761bd: field keywords element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (e *EmojiKeywordsDifference) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode emojiKeywordsDifference#5cc761bd to nil")
	}
	if err := b.ConsumeID(EmojiKeywordsDifferenceTypeID); err != nil {
		return fmt.Errorf("unable to decode emojiKeywordsDifference#5cc761bd: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode emojiKeywordsDifference#5cc761bd: field lang_code: %w", err)
		}
		e.LangCode = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode emojiKeywordsDifference#5cc761bd: field from_version: %w", err)
		}
		e.FromVersion = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode emojiKeywordsDifference#5cc761bd: field version: %w", err)
		}
		e.Version = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode emojiKeywordsDifference#5cc761bd: field keywords: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeEmojiKeyword(b)
			if err != nil {
				return fmt.Errorf("unable to decode emojiKeywordsDifference#5cc761bd: field keywords: %w", err)
			}
			e.Keywords = append(e.Keywords, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for EmojiKeywordsDifference.
var (
	_ bin.Encoder = &EmojiKeywordsDifference{}
	_ bin.Decoder = &EmojiKeywordsDifference{}
)
