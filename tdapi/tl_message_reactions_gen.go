// Code generated by gotdgen, DO NOT EDIT.

package tdapi

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// MessageReactions represents TL type `messageReactions#91986cd8`.
type MessageReactions struct {
	// List of added reactions
	Reactions []MessageReaction
	// True, if the reactions are tags and Telegram Premium users can filter messages by
	// them; currently, always false
	AreTags bool
}

// MessageReactionsTypeID is TL type id of MessageReactions.
const MessageReactionsTypeID = 0x91986cd8

// Ensuring interfaces in compile-time for MessageReactions.
var (
	_ bin.Encoder     = &MessageReactions{}
	_ bin.Decoder     = &MessageReactions{}
	_ bin.BareEncoder = &MessageReactions{}
	_ bin.BareDecoder = &MessageReactions{}
)

func (m *MessageReactions) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.Reactions == nil) {
		return false
	}
	if !(m.AreTags == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessageReactions) String() string {
	if m == nil {
		return "MessageReactions(nil)"
	}
	type Alias MessageReactions
	return fmt.Sprintf("MessageReactions%+v", Alias(*m))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessageReactions) TypeID() uint32 {
	return MessageReactionsTypeID
}

// TypeName returns name of type in TL schema.
func (*MessageReactions) TypeName() string {
	return "messageReactions"
}

// TypeInfo returns info about TL type.
func (m *MessageReactions) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messageReactions",
		ID:   MessageReactionsTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Reactions",
			SchemaName: "reactions",
		},
		{
			Name:       "AreTags",
			SchemaName: "are_tags",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *MessageReactions) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageReactions#91986cd8 as nil")
	}
	b.PutID(MessageReactionsTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MessageReactions) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageReactions#91986cd8 as nil")
	}
	b.PutInt(len(m.Reactions))
	for idx, v := range m.Reactions {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare messageReactions#91986cd8: field reactions element with index %d: %w", idx, err)
		}
	}
	b.PutBool(m.AreTags)
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageReactions) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageReactions#91986cd8 to nil")
	}
	if err := b.ConsumeID(MessageReactionsTypeID); err != nil {
		return fmt.Errorf("unable to decode messageReactions#91986cd8: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MessageReactions) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageReactions#91986cd8 to nil")
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageReactions#91986cd8: field reactions: %w", err)
		}

		if headerLen > 0 {
			m.Reactions = make([]MessageReaction, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value MessageReaction
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare messageReactions#91986cd8: field reactions: %w", err)
			}
			m.Reactions = append(m.Reactions, value)
		}
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode messageReactions#91986cd8: field are_tags: %w", err)
		}
		m.AreTags = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (m *MessageReactions) EncodeTDLibJSON(b tdjson.Encoder) error {
	if m == nil {
		return fmt.Errorf("can't encode messageReactions#91986cd8 as nil")
	}
	b.ObjStart()
	b.PutID("messageReactions")
	b.Comma()
	b.FieldStart("reactions")
	b.ArrStart()
	for idx, v := range m.Reactions {
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode messageReactions#91986cd8: field reactions element with index %d: %w", idx, err)
		}
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.FieldStart("are_tags")
	b.PutBool(m.AreTags)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (m *MessageReactions) DecodeTDLibJSON(b tdjson.Decoder) error {
	if m == nil {
		return fmt.Errorf("can't decode messageReactions#91986cd8 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("messageReactions"); err != nil {
				return fmt.Errorf("unable to decode messageReactions#91986cd8: %w", err)
			}
		case "reactions":
			if err := b.Arr(func(b tdjson.Decoder) error {
				var value MessageReaction
				if err := value.DecodeTDLibJSON(b); err != nil {
					return fmt.Errorf("unable to decode messageReactions#91986cd8: field reactions: %w", err)
				}
				m.Reactions = append(m.Reactions, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode messageReactions#91986cd8: field reactions: %w", err)
			}
		case "are_tags":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode messageReactions#91986cd8: field are_tags: %w", err)
			}
			m.AreTags = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetReactions returns value of Reactions field.
func (m *MessageReactions) GetReactions() (value []MessageReaction) {
	if m == nil {
		return
	}
	return m.Reactions
}

// GetAreTags returns value of AreTags field.
func (m *MessageReactions) GetAreTags() (value bool) {
	if m == nil {
		return
	}
	return m.AreTags
}
