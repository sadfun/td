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

// GetStoryPublicForwardsRequest represents TL type `getStoryPublicForwards#68f7dcbb`.
type GetStoryPublicForwardsRequest struct {
	// The identifier of the sender of the story
	StorySenderChatID int64
	// The identifier of the story
	StoryID int32
	// Offset of the first entry to return as received from the previous request; use empty
	// string to get the first chunk of results
	Offset string
	// The maximum number of messages and stories to be returned; must be positive and can't
	// be greater than 100. For optimal performance, the number of returned objects is chosen
	// by TDLib and can be smaller than the specified limit
	Limit int32
}

// GetStoryPublicForwardsRequestTypeID is TL type id of GetStoryPublicForwardsRequest.
const GetStoryPublicForwardsRequestTypeID = 0x68f7dcbb

// Ensuring interfaces in compile-time for GetStoryPublicForwardsRequest.
var (
	_ bin.Encoder     = &GetStoryPublicForwardsRequest{}
	_ bin.Decoder     = &GetStoryPublicForwardsRequest{}
	_ bin.BareEncoder = &GetStoryPublicForwardsRequest{}
	_ bin.BareDecoder = &GetStoryPublicForwardsRequest{}
)

func (g *GetStoryPublicForwardsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.StorySenderChatID == 0) {
		return false
	}
	if !(g.StoryID == 0) {
		return false
	}
	if !(g.Offset == "") {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetStoryPublicForwardsRequest) String() string {
	if g == nil {
		return "GetStoryPublicForwardsRequest(nil)"
	}
	type Alias GetStoryPublicForwardsRequest
	return fmt.Sprintf("GetStoryPublicForwardsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetStoryPublicForwardsRequest) TypeID() uint32 {
	return GetStoryPublicForwardsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetStoryPublicForwardsRequest) TypeName() string {
	return "getStoryPublicForwards"
}

// TypeInfo returns info about TL type.
func (g *GetStoryPublicForwardsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getStoryPublicForwards",
		ID:   GetStoryPublicForwardsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "StorySenderChatID",
			SchemaName: "story_sender_chat_id",
		},
		{
			Name:       "StoryID",
			SchemaName: "story_id",
		},
		{
			Name:       "Offset",
			SchemaName: "offset",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetStoryPublicForwardsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getStoryPublicForwards#68f7dcbb as nil")
	}
	b.PutID(GetStoryPublicForwardsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetStoryPublicForwardsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getStoryPublicForwards#68f7dcbb as nil")
	}
	b.PutInt53(g.StorySenderChatID)
	b.PutInt32(g.StoryID)
	b.PutString(g.Offset)
	b.PutInt32(g.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetStoryPublicForwardsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getStoryPublicForwards#68f7dcbb to nil")
	}
	if err := b.ConsumeID(GetStoryPublicForwardsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getStoryPublicForwards#68f7dcbb: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetStoryPublicForwardsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getStoryPublicForwards#68f7dcbb to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getStoryPublicForwards#68f7dcbb: field story_sender_chat_id: %w", err)
		}
		g.StorySenderChatID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getStoryPublicForwards#68f7dcbb: field story_id: %w", err)
		}
		g.StoryID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode getStoryPublicForwards#68f7dcbb: field offset: %w", err)
		}
		g.Offset = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getStoryPublicForwards#68f7dcbb: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetStoryPublicForwardsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getStoryPublicForwards#68f7dcbb as nil")
	}
	b.ObjStart()
	b.PutID("getStoryPublicForwards")
	b.Comma()
	b.FieldStart("story_sender_chat_id")
	b.PutInt53(g.StorySenderChatID)
	b.Comma()
	b.FieldStart("story_id")
	b.PutInt32(g.StoryID)
	b.Comma()
	b.FieldStart("offset")
	b.PutString(g.Offset)
	b.Comma()
	b.FieldStart("limit")
	b.PutInt32(g.Limit)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetStoryPublicForwardsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getStoryPublicForwards#68f7dcbb to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getStoryPublicForwards"); err != nil {
				return fmt.Errorf("unable to decode getStoryPublicForwards#68f7dcbb: %w", err)
			}
		case "story_sender_chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getStoryPublicForwards#68f7dcbb: field story_sender_chat_id: %w", err)
			}
			g.StorySenderChatID = value
		case "story_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode getStoryPublicForwards#68f7dcbb: field story_id: %w", err)
			}
			g.StoryID = value
		case "offset":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode getStoryPublicForwards#68f7dcbb: field offset: %w", err)
			}
			g.Offset = value
		case "limit":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode getStoryPublicForwards#68f7dcbb: field limit: %w", err)
			}
			g.Limit = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetStorySenderChatID returns value of StorySenderChatID field.
func (g *GetStoryPublicForwardsRequest) GetStorySenderChatID() (value int64) {
	if g == nil {
		return
	}
	return g.StorySenderChatID
}

// GetStoryID returns value of StoryID field.
func (g *GetStoryPublicForwardsRequest) GetStoryID() (value int32) {
	if g == nil {
		return
	}
	return g.StoryID
}

// GetOffset returns value of Offset field.
func (g *GetStoryPublicForwardsRequest) GetOffset() (value string) {
	if g == nil {
		return
	}
	return g.Offset
}

// GetLimit returns value of Limit field.
func (g *GetStoryPublicForwardsRequest) GetLimit() (value int32) {
	if g == nil {
		return
	}
	return g.Limit
}

// GetStoryPublicForwards invokes method getStoryPublicForwards#68f7dcbb returning error if any.
func (c *Client) GetStoryPublicForwards(ctx context.Context, request *GetStoryPublicForwardsRequest) (*PublicForwards, error) {
	var result PublicForwards

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
