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

// GetChatMessagePositionRequest represents TL type `getChatMessagePosition#e9d7d767`.
type GetChatMessagePositionRequest struct {
	// Identifier of the chat in which to find message position
	ChatID int64
	// Message identifier
	MessageID int64
	// Filter for message content; searchMessagesFilterEmpty,
	// searchMessagesFilterUnreadMention, searchMessagesFilterUnreadReaction, and
	// searchMessagesFilterFailedToSend are unsupported in this function
	Filter SearchMessagesFilterClass
	// If not 0, only messages in the specified thread will be considered; supergroups only
	MessageThreadID int64
	// If not null, only messages in the specified Saved Messages topic will be considered;
	// pass null to consider all relevant messages, or for chats other than Saved Messages
	SavedMessagesTopic SavedMessagesTopicClass
}

// GetChatMessagePositionRequestTypeID is TL type id of GetChatMessagePositionRequest.
const GetChatMessagePositionRequestTypeID = 0xe9d7d767

// Ensuring interfaces in compile-time for GetChatMessagePositionRequest.
var (
	_ bin.Encoder     = &GetChatMessagePositionRequest{}
	_ bin.Decoder     = &GetChatMessagePositionRequest{}
	_ bin.BareEncoder = &GetChatMessagePositionRequest{}
	_ bin.BareDecoder = &GetChatMessagePositionRequest{}
)

func (g *GetChatMessagePositionRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ChatID == 0) {
		return false
	}
	if !(g.MessageID == 0) {
		return false
	}
	if !(g.Filter == nil) {
		return false
	}
	if !(g.MessageThreadID == 0) {
		return false
	}
	if !(g.SavedMessagesTopic == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetChatMessagePositionRequest) String() string {
	if g == nil {
		return "GetChatMessagePositionRequest(nil)"
	}
	type Alias GetChatMessagePositionRequest
	return fmt.Sprintf("GetChatMessagePositionRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetChatMessagePositionRequest) TypeID() uint32 {
	return GetChatMessagePositionRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetChatMessagePositionRequest) TypeName() string {
	return "getChatMessagePosition"
}

// TypeInfo returns info about TL type.
func (g *GetChatMessagePositionRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getChatMessagePosition",
		ID:   GetChatMessagePositionRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "MessageID",
			SchemaName: "message_id",
		},
		{
			Name:       "Filter",
			SchemaName: "filter",
		},
		{
			Name:       "MessageThreadID",
			SchemaName: "message_thread_id",
		},
		{
			Name:       "SavedMessagesTopic",
			SchemaName: "saved_messages_topic",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetChatMessagePositionRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatMessagePosition#e9d7d767 as nil")
	}
	b.PutID(GetChatMessagePositionRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetChatMessagePositionRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatMessagePosition#e9d7d767 as nil")
	}
	b.PutInt53(g.ChatID)
	b.PutInt53(g.MessageID)
	if g.Filter == nil {
		return fmt.Errorf("unable to encode getChatMessagePosition#e9d7d767: field filter is nil")
	}
	if err := g.Filter.Encode(b); err != nil {
		return fmt.Errorf("unable to encode getChatMessagePosition#e9d7d767: field filter: %w", err)
	}
	b.PutInt53(g.MessageThreadID)
	if g.SavedMessagesTopic == nil {
		return fmt.Errorf("unable to encode getChatMessagePosition#e9d7d767: field saved_messages_topic is nil")
	}
	if err := g.SavedMessagesTopic.Encode(b); err != nil {
		return fmt.Errorf("unable to encode getChatMessagePosition#e9d7d767: field saved_messages_topic: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GetChatMessagePositionRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatMessagePosition#e9d7d767 to nil")
	}
	if err := b.ConsumeID(GetChatMessagePositionRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getChatMessagePosition#e9d7d767: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetChatMessagePositionRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatMessagePosition#e9d7d767 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getChatMessagePosition#e9d7d767: field chat_id: %w", err)
		}
		g.ChatID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getChatMessagePosition#e9d7d767: field message_id: %w", err)
		}
		g.MessageID = value
	}
	{
		value, err := DecodeSearchMessagesFilter(b)
		if err != nil {
			return fmt.Errorf("unable to decode getChatMessagePosition#e9d7d767: field filter: %w", err)
		}
		g.Filter = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getChatMessagePosition#e9d7d767: field message_thread_id: %w", err)
		}
		g.MessageThreadID = value
	}
	{
		value, err := DecodeSavedMessagesTopic(b)
		if err != nil {
			return fmt.Errorf("unable to decode getChatMessagePosition#e9d7d767: field saved_messages_topic: %w", err)
		}
		g.SavedMessagesTopic = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetChatMessagePositionRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatMessagePosition#e9d7d767 as nil")
	}
	b.ObjStart()
	b.PutID("getChatMessagePosition")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(g.ChatID)
	b.Comma()
	b.FieldStart("message_id")
	b.PutInt53(g.MessageID)
	b.Comma()
	b.FieldStart("filter")
	if g.Filter == nil {
		return fmt.Errorf("unable to encode getChatMessagePosition#e9d7d767: field filter is nil")
	}
	if err := g.Filter.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode getChatMessagePosition#e9d7d767: field filter: %w", err)
	}
	b.Comma()
	b.FieldStart("message_thread_id")
	b.PutInt53(g.MessageThreadID)
	b.Comma()
	b.FieldStart("saved_messages_topic")
	if g.SavedMessagesTopic == nil {
		return fmt.Errorf("unable to encode getChatMessagePosition#e9d7d767: field saved_messages_topic is nil")
	}
	if err := g.SavedMessagesTopic.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode getChatMessagePosition#e9d7d767: field saved_messages_topic: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetChatMessagePositionRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatMessagePosition#e9d7d767 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getChatMessagePosition"); err != nil {
				return fmt.Errorf("unable to decode getChatMessagePosition#e9d7d767: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getChatMessagePosition#e9d7d767: field chat_id: %w", err)
			}
			g.ChatID = value
		case "message_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getChatMessagePosition#e9d7d767: field message_id: %w", err)
			}
			g.MessageID = value
		case "filter":
			value, err := DecodeTDLibJSONSearchMessagesFilter(b)
			if err != nil {
				return fmt.Errorf("unable to decode getChatMessagePosition#e9d7d767: field filter: %w", err)
			}
			g.Filter = value
		case "message_thread_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getChatMessagePosition#e9d7d767: field message_thread_id: %w", err)
			}
			g.MessageThreadID = value
		case "saved_messages_topic":
			value, err := DecodeTDLibJSONSavedMessagesTopic(b)
			if err != nil {
				return fmt.Errorf("unable to decode getChatMessagePosition#e9d7d767: field saved_messages_topic: %w", err)
			}
			g.SavedMessagesTopic = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (g *GetChatMessagePositionRequest) GetChatID() (value int64) {
	if g == nil {
		return
	}
	return g.ChatID
}

// GetMessageID returns value of MessageID field.
func (g *GetChatMessagePositionRequest) GetMessageID() (value int64) {
	if g == nil {
		return
	}
	return g.MessageID
}

// GetFilter returns value of Filter field.
func (g *GetChatMessagePositionRequest) GetFilter() (value SearchMessagesFilterClass) {
	if g == nil {
		return
	}
	return g.Filter
}

// GetMessageThreadID returns value of MessageThreadID field.
func (g *GetChatMessagePositionRequest) GetMessageThreadID() (value int64) {
	if g == nil {
		return
	}
	return g.MessageThreadID
}

// GetSavedMessagesTopic returns value of SavedMessagesTopic field.
func (g *GetChatMessagePositionRequest) GetSavedMessagesTopic() (value SavedMessagesTopicClass) {
	if g == nil {
		return
	}
	return g.SavedMessagesTopic
}

// GetChatMessagePosition invokes method getChatMessagePosition#e9d7d767 returning error if any.
func (c *Client) GetChatMessagePosition(ctx context.Context, request *GetChatMessagePositionRequest) (*Count, error) {
	var result Count

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
