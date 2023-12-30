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

// GetDisallowedChatEmojiStatusesRequest represents TL type `getDisallowedChatEmojiStatuses#d2144da0`.
type GetDisallowedChatEmojiStatusesRequest struct {
}

// GetDisallowedChatEmojiStatusesRequestTypeID is TL type id of GetDisallowedChatEmojiStatusesRequest.
const GetDisallowedChatEmojiStatusesRequestTypeID = 0xd2144da0

// Ensuring interfaces in compile-time for GetDisallowedChatEmojiStatusesRequest.
var (
	_ bin.Encoder     = &GetDisallowedChatEmojiStatusesRequest{}
	_ bin.Decoder     = &GetDisallowedChatEmojiStatusesRequest{}
	_ bin.BareEncoder = &GetDisallowedChatEmojiStatusesRequest{}
	_ bin.BareDecoder = &GetDisallowedChatEmojiStatusesRequest{}
)

func (g *GetDisallowedChatEmojiStatusesRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetDisallowedChatEmojiStatusesRequest) String() string {
	if g == nil {
		return "GetDisallowedChatEmojiStatusesRequest(nil)"
	}
	type Alias GetDisallowedChatEmojiStatusesRequest
	return fmt.Sprintf("GetDisallowedChatEmojiStatusesRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetDisallowedChatEmojiStatusesRequest) TypeID() uint32 {
	return GetDisallowedChatEmojiStatusesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetDisallowedChatEmojiStatusesRequest) TypeName() string {
	return "getDisallowedChatEmojiStatuses"
}

// TypeInfo returns info about TL type.
func (g *GetDisallowedChatEmojiStatusesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getDisallowedChatEmojiStatuses",
		ID:   GetDisallowedChatEmojiStatusesRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetDisallowedChatEmojiStatusesRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getDisallowedChatEmojiStatuses#d2144da0 as nil")
	}
	b.PutID(GetDisallowedChatEmojiStatusesRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetDisallowedChatEmojiStatusesRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getDisallowedChatEmojiStatuses#d2144da0 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GetDisallowedChatEmojiStatusesRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getDisallowedChatEmojiStatuses#d2144da0 to nil")
	}
	if err := b.ConsumeID(GetDisallowedChatEmojiStatusesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getDisallowedChatEmojiStatuses#d2144da0: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetDisallowedChatEmojiStatusesRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getDisallowedChatEmojiStatuses#d2144da0 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetDisallowedChatEmojiStatusesRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getDisallowedChatEmojiStatuses#d2144da0 as nil")
	}
	b.ObjStart()
	b.PutID("getDisallowedChatEmojiStatuses")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetDisallowedChatEmojiStatusesRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getDisallowedChatEmojiStatuses#d2144da0 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getDisallowedChatEmojiStatuses"); err != nil {
				return fmt.Errorf("unable to decode getDisallowedChatEmojiStatuses#d2144da0: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetDisallowedChatEmojiStatuses invokes method getDisallowedChatEmojiStatuses#d2144da0 returning error if any.
func (c *Client) GetDisallowedChatEmojiStatuses(ctx context.Context) (*EmojiStatuses, error) {
	var result EmojiStatuses

	request := &GetDisallowedChatEmojiStatusesRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
