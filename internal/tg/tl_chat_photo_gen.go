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

// ChatPhotoEmpty represents TL type `chatPhotoEmpty#37c1011c`.
type ChatPhotoEmpty struct {
}

// ChatPhotoEmptyTypeID is TL type id of ChatPhotoEmpty.
const ChatPhotoEmptyTypeID = 0x37c1011c

// Encode implements bin.Encoder.
func (c *ChatPhotoEmpty) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatPhotoEmpty#37c1011c as nil")
	}
	b.PutID(ChatPhotoEmptyTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatPhotoEmpty) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatPhotoEmpty#37c1011c to nil")
	}
	if err := b.ConsumeID(ChatPhotoEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode chatPhotoEmpty#37c1011c: %w", err)
	}
	return nil
}

// construct implements constructor of ChatPhotoClass.
func (c ChatPhotoEmpty) construct() ChatPhotoClass { return &c }

// Ensuring interfaces in compile-time for ChatPhotoEmpty.
var (
	_ bin.Encoder = &ChatPhotoEmpty{}
	_ bin.Decoder = &ChatPhotoEmpty{}

	_ ChatPhotoClass = &ChatPhotoEmpty{}
)

// ChatPhoto represents TL type `chatPhoto#d20b9f3c`.
type ChatPhoto struct {
	// Flags field of ChatPhoto.
	Flags bin.Fields
	// HasVideo field of ChatPhoto.
	HasVideo bool
	// PhotoSmall field of ChatPhoto.
	PhotoSmall FileLocationToBeDeprecated
	// PhotoBig field of ChatPhoto.
	PhotoBig FileLocationToBeDeprecated
	// DCID field of ChatPhoto.
	DCID int
}

// ChatPhotoTypeID is TL type id of ChatPhoto.
const ChatPhotoTypeID = 0xd20b9f3c

// Encode implements bin.Encoder.
func (c *ChatPhoto) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatPhoto#d20b9f3c as nil")
	}
	b.PutID(ChatPhotoTypeID)
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chatPhoto#d20b9f3c: field flags: %w", err)
	}
	if err := c.PhotoSmall.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chatPhoto#d20b9f3c: field photo_small: %w", err)
	}
	if err := c.PhotoBig.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chatPhoto#d20b9f3c: field photo_big: %w", err)
	}
	b.PutInt(c.DCID)
	return nil
}

// SetHasVideo sets value of HasVideo conditional field.
func (c *ChatPhoto) SetHasVideo(value bool) {
	if value {
		c.Flags.Set(0)
	} else {
		c.Flags.Unset(0)
	}
}

// Decode implements bin.Decoder.
func (c *ChatPhoto) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatPhoto#d20b9f3c to nil")
	}
	if err := b.ConsumeID(ChatPhotoTypeID); err != nil {
		return fmt.Errorf("unable to decode chatPhoto#d20b9f3c: %w", err)
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode chatPhoto#d20b9f3c: field flags: %w", err)
		}
	}
	c.HasVideo = c.Flags.Has(0)
	{
		if err := c.PhotoSmall.Decode(b); err != nil {
			return fmt.Errorf("unable to decode chatPhoto#d20b9f3c: field photo_small: %w", err)
		}
	}
	{
		if err := c.PhotoBig.Decode(b); err != nil {
			return fmt.Errorf("unable to decode chatPhoto#d20b9f3c: field photo_big: %w", err)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatPhoto#d20b9f3c: field dc_id: %w", err)
		}
		c.DCID = value
	}
	return nil
}

// construct implements constructor of ChatPhotoClass.
func (c ChatPhoto) construct() ChatPhotoClass { return &c }

// Ensuring interfaces in compile-time for ChatPhoto.
var (
	_ bin.Encoder = &ChatPhoto{}
	_ bin.Decoder = &ChatPhoto{}

	_ ChatPhotoClass = &ChatPhoto{}
)

// ChatPhotoClass represents ChatPhoto generic type.
//
// Example:
//  g, err := DecodeChatPhoto(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *ChatPhotoEmpty: // chatPhotoEmpty#37c1011c
//  case *ChatPhoto: // chatPhoto#d20b9f3c
//  default: panic(v)
//  }
type ChatPhotoClass interface {
	bin.Encoder
	bin.Decoder
	construct() ChatPhotoClass
}

// DecodeChatPhoto implements binary de-serialization for ChatPhotoClass.
func DecodeChatPhoto(buf *bin.Buffer) (ChatPhotoClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case ChatPhotoEmptyTypeID:
		// Decoding chatPhotoEmpty#37c1011c.
		v := ChatPhotoEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatPhotoClass: %w", err)
		}
		return &v, nil
	case ChatPhotoTypeID:
		// Decoding chatPhoto#d20b9f3c.
		v := ChatPhoto{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatPhotoClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ChatPhotoClass: %w", bin.NewUnexpectedID(id))
	}
}
