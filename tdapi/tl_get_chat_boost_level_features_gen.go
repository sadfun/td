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

// GetChatBoostLevelFeaturesRequest represents TL type `getChatBoostLevelFeatures#71f15dfb`.
type GetChatBoostLevelFeaturesRequest struct {
	// Chat boost level
	Level int32
}

// GetChatBoostLevelFeaturesRequestTypeID is TL type id of GetChatBoostLevelFeaturesRequest.
const GetChatBoostLevelFeaturesRequestTypeID = 0x71f15dfb

// Ensuring interfaces in compile-time for GetChatBoostLevelFeaturesRequest.
var (
	_ bin.Encoder     = &GetChatBoostLevelFeaturesRequest{}
	_ bin.Decoder     = &GetChatBoostLevelFeaturesRequest{}
	_ bin.BareEncoder = &GetChatBoostLevelFeaturesRequest{}
	_ bin.BareDecoder = &GetChatBoostLevelFeaturesRequest{}
)

func (g *GetChatBoostLevelFeaturesRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Level == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetChatBoostLevelFeaturesRequest) String() string {
	if g == nil {
		return "GetChatBoostLevelFeaturesRequest(nil)"
	}
	type Alias GetChatBoostLevelFeaturesRequest
	return fmt.Sprintf("GetChatBoostLevelFeaturesRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetChatBoostLevelFeaturesRequest) TypeID() uint32 {
	return GetChatBoostLevelFeaturesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetChatBoostLevelFeaturesRequest) TypeName() string {
	return "getChatBoostLevelFeatures"
}

// TypeInfo returns info about TL type.
func (g *GetChatBoostLevelFeaturesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getChatBoostLevelFeatures",
		ID:   GetChatBoostLevelFeaturesRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Level",
			SchemaName: "level",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetChatBoostLevelFeaturesRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatBoostLevelFeatures#71f15dfb as nil")
	}
	b.PutID(GetChatBoostLevelFeaturesRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetChatBoostLevelFeaturesRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatBoostLevelFeatures#71f15dfb as nil")
	}
	b.PutInt32(g.Level)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetChatBoostLevelFeaturesRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatBoostLevelFeatures#71f15dfb to nil")
	}
	if err := b.ConsumeID(GetChatBoostLevelFeaturesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getChatBoostLevelFeatures#71f15dfb: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetChatBoostLevelFeaturesRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatBoostLevelFeatures#71f15dfb to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getChatBoostLevelFeatures#71f15dfb: field level: %w", err)
		}
		g.Level = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetChatBoostLevelFeaturesRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatBoostLevelFeatures#71f15dfb as nil")
	}
	b.ObjStart()
	b.PutID("getChatBoostLevelFeatures")
	b.Comma()
	b.FieldStart("level")
	b.PutInt32(g.Level)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetChatBoostLevelFeaturesRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatBoostLevelFeatures#71f15dfb to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getChatBoostLevelFeatures"); err != nil {
				return fmt.Errorf("unable to decode getChatBoostLevelFeatures#71f15dfb: %w", err)
			}
		case "level":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode getChatBoostLevelFeatures#71f15dfb: field level: %w", err)
			}
			g.Level = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetLevel returns value of Level field.
func (g *GetChatBoostLevelFeaturesRequest) GetLevel() (value int32) {
	if g == nil {
		return
	}
	return g.Level
}

// GetChatBoostLevelFeatures invokes method getChatBoostLevelFeatures#71f15dfb returning error if any.
func (c *Client) GetChatBoostLevelFeatures(ctx context.Context, level int32) (*ChatBoostLevelFeatures, error) {
	var result ChatBoostLevelFeatures

	request := &GetChatBoostLevelFeaturesRequest{
		Level: level,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
