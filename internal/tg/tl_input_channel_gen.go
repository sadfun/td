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

// InputChannelEmpty represents TL type `inputChannelEmpty#ee8c1e86`.
type InputChannelEmpty struct {
}

// InputChannelEmptyTypeID is TL type id of InputChannelEmpty.
const InputChannelEmptyTypeID = 0xee8c1e86

// Encode implements bin.Encoder.
func (i *InputChannelEmpty) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputChannelEmpty#ee8c1e86 as nil")
	}
	b.PutID(InputChannelEmptyTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputChannelEmpty) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputChannelEmpty#ee8c1e86 to nil")
	}
	if err := b.ConsumeID(InputChannelEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode inputChannelEmpty#ee8c1e86: %w", err)
	}
	return nil
}

// construct implements constructor of InputChannelClass.
func (i InputChannelEmpty) construct() InputChannelClass { return &i }

// Ensuring interfaces in compile-time for InputChannelEmpty.
var (
	_ bin.Encoder = &InputChannelEmpty{}
	_ bin.Decoder = &InputChannelEmpty{}

	_ InputChannelClass = &InputChannelEmpty{}
)

// InputChannel represents TL type `inputChannel#afeb712e`.
type InputChannel struct {
	// ChannelID field of InputChannel.
	ChannelID int
	// AccessHash field of InputChannel.
	AccessHash int64
}

// InputChannelTypeID is TL type id of InputChannel.
const InputChannelTypeID = 0xafeb712e

// Encode implements bin.Encoder.
func (i *InputChannel) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputChannel#afeb712e as nil")
	}
	b.PutID(InputChannelTypeID)
	b.PutInt(i.ChannelID)
	b.PutLong(i.AccessHash)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputChannel) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputChannel#afeb712e to nil")
	}
	if err := b.ConsumeID(InputChannelTypeID); err != nil {
		return fmt.Errorf("unable to decode inputChannel#afeb712e: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputChannel#afeb712e: field channel_id: %w", err)
		}
		i.ChannelID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputChannel#afeb712e: field access_hash: %w", err)
		}
		i.AccessHash = value
	}
	return nil
}

// construct implements constructor of InputChannelClass.
func (i InputChannel) construct() InputChannelClass { return &i }

// Ensuring interfaces in compile-time for InputChannel.
var (
	_ bin.Encoder = &InputChannel{}
	_ bin.Decoder = &InputChannel{}

	_ InputChannelClass = &InputChannel{}
)

// InputChannelFromMessage represents TL type `inputChannelFromMessage#2a286531`.
type InputChannelFromMessage struct {
	// Peer field of InputChannelFromMessage.
	Peer InputPeerClass
	// MsgID field of InputChannelFromMessage.
	MsgID int
	// ChannelID field of InputChannelFromMessage.
	ChannelID int
}

// InputChannelFromMessageTypeID is TL type id of InputChannelFromMessage.
const InputChannelFromMessageTypeID = 0x2a286531

// Encode implements bin.Encoder.
func (i *InputChannelFromMessage) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputChannelFromMessage#2a286531 as nil")
	}
	b.PutID(InputChannelFromMessageTypeID)
	if i.Peer == nil {
		return fmt.Errorf("unable to encode inputChannelFromMessage#2a286531: field peer is nil")
	}
	if err := i.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputChannelFromMessage#2a286531: field peer: %w", err)
	}
	b.PutInt(i.MsgID)
	b.PutInt(i.ChannelID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputChannelFromMessage) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputChannelFromMessage#2a286531 to nil")
	}
	if err := b.ConsumeID(InputChannelFromMessageTypeID); err != nil {
		return fmt.Errorf("unable to decode inputChannelFromMessage#2a286531: %w", err)
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputChannelFromMessage#2a286531: field peer: %w", err)
		}
		i.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputChannelFromMessage#2a286531: field msg_id: %w", err)
		}
		i.MsgID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputChannelFromMessage#2a286531: field channel_id: %w", err)
		}
		i.ChannelID = value
	}
	return nil
}

// construct implements constructor of InputChannelClass.
func (i InputChannelFromMessage) construct() InputChannelClass { return &i }

// Ensuring interfaces in compile-time for InputChannelFromMessage.
var (
	_ bin.Encoder = &InputChannelFromMessage{}
	_ bin.Decoder = &InputChannelFromMessage{}

	_ InputChannelClass = &InputChannelFromMessage{}
)

// InputChannelClass represents InputChannel generic type.
//
// Example:
//  g, err := DecodeInputChannel(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *InputChannelEmpty: // inputChannelEmpty#ee8c1e86
//  case *InputChannel: // inputChannel#afeb712e
//  case *InputChannelFromMessage: // inputChannelFromMessage#2a286531
//  default: panic(v)
//  }
type InputChannelClass interface {
	bin.Encoder
	bin.Decoder
	construct() InputChannelClass
}

// DecodeInputChannel implements binary de-serialization for InputChannelClass.
func DecodeInputChannel(buf *bin.Buffer) (InputChannelClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputChannelEmptyTypeID:
		// Decoding inputChannelEmpty#ee8c1e86.
		v := InputChannelEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputChannelClass: %w", err)
		}
		return &v, nil
	case InputChannelTypeID:
		// Decoding inputChannel#afeb712e.
		v := InputChannel{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputChannelClass: %w", err)
		}
		return &v, nil
	case InputChannelFromMessageTypeID:
		// Decoding inputChannelFromMessage#2a286531.
		v := InputChannelFromMessage{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputChannelClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputChannelClass: %w", bin.NewUnexpectedID(id))
	}
}
