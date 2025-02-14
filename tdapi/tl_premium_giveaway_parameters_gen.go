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

// PremiumGiveawayParameters represents TL type `premiumGiveawayParameters#f9f5bf5e`.
type PremiumGiveawayParameters struct {
	// Identifier of the channel chat, which will be automatically boosted by the winners of
	// the giveaway for duration of the Premium subscription
	BoostedChatID int64
	// Identifiers of other channel chats that must be subscribed by the users to be eligible
	// for the giveaway. There can be up to getOption("giveaway_additional_chat_count_max")
	// additional chats
	AdditionalChatIDs []int64
	// Point in time (Unix timestamp) when the giveaway is expected to be performed; must be
	// 60-getOption("giveaway_duration_max") seconds in the future in scheduled giveaways
	WinnersSelectionDate int32
	// True, if only new members of the chats will be eligible for the giveaway
	OnlyNewMembers bool
	// True, if the list of winners of the giveaway will be available to everyone
	HasPublicWinners bool
	// The list of two-letter ISO 3166-1 alpha-2 codes of countries, users from which will be
	// eligible for the giveaway. If empty, then all users can participate in the giveaway.
	CountryCodes []string
	// Additional description of the giveaway prize; 0-128 characters
	PrizeDescription string
}

// PremiumGiveawayParametersTypeID is TL type id of PremiumGiveawayParameters.
const PremiumGiveawayParametersTypeID = 0xf9f5bf5e

// Ensuring interfaces in compile-time for PremiumGiveawayParameters.
var (
	_ bin.Encoder     = &PremiumGiveawayParameters{}
	_ bin.Decoder     = &PremiumGiveawayParameters{}
	_ bin.BareEncoder = &PremiumGiveawayParameters{}
	_ bin.BareDecoder = &PremiumGiveawayParameters{}
)

func (p *PremiumGiveawayParameters) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.BoostedChatID == 0) {
		return false
	}
	if !(p.AdditionalChatIDs == nil) {
		return false
	}
	if !(p.WinnersSelectionDate == 0) {
		return false
	}
	if !(p.OnlyNewMembers == false) {
		return false
	}
	if !(p.HasPublicWinners == false) {
		return false
	}
	if !(p.CountryCodes == nil) {
		return false
	}
	if !(p.PrizeDescription == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PremiumGiveawayParameters) String() string {
	if p == nil {
		return "PremiumGiveawayParameters(nil)"
	}
	type Alias PremiumGiveawayParameters
	return fmt.Sprintf("PremiumGiveawayParameters%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PremiumGiveawayParameters) TypeID() uint32 {
	return PremiumGiveawayParametersTypeID
}

// TypeName returns name of type in TL schema.
func (*PremiumGiveawayParameters) TypeName() string {
	return "premiumGiveawayParameters"
}

// TypeInfo returns info about TL type.
func (p *PremiumGiveawayParameters) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "premiumGiveawayParameters",
		ID:   PremiumGiveawayParametersTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "BoostedChatID",
			SchemaName: "boosted_chat_id",
		},
		{
			Name:       "AdditionalChatIDs",
			SchemaName: "additional_chat_ids",
		},
		{
			Name:       "WinnersSelectionDate",
			SchemaName: "winners_selection_date",
		},
		{
			Name:       "OnlyNewMembers",
			SchemaName: "only_new_members",
		},
		{
			Name:       "HasPublicWinners",
			SchemaName: "has_public_winners",
		},
		{
			Name:       "CountryCodes",
			SchemaName: "country_codes",
		},
		{
			Name:       "PrizeDescription",
			SchemaName: "prize_description",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PremiumGiveawayParameters) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode premiumGiveawayParameters#f9f5bf5e as nil")
	}
	b.PutID(PremiumGiveawayParametersTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PremiumGiveawayParameters) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode premiumGiveawayParameters#f9f5bf5e as nil")
	}
	b.PutInt53(p.BoostedChatID)
	b.PutInt(len(p.AdditionalChatIDs))
	for _, v := range p.AdditionalChatIDs {
		b.PutInt53(v)
	}
	b.PutInt32(p.WinnersSelectionDate)
	b.PutBool(p.OnlyNewMembers)
	b.PutBool(p.HasPublicWinners)
	b.PutInt(len(p.CountryCodes))
	for _, v := range p.CountryCodes {
		b.PutString(v)
	}
	b.PutString(p.PrizeDescription)
	return nil
}

// Decode implements bin.Decoder.
func (p *PremiumGiveawayParameters) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode premiumGiveawayParameters#f9f5bf5e to nil")
	}
	if err := b.ConsumeID(PremiumGiveawayParametersTypeID); err != nil {
		return fmt.Errorf("unable to decode premiumGiveawayParameters#f9f5bf5e: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PremiumGiveawayParameters) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode premiumGiveawayParameters#f9f5bf5e to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode premiumGiveawayParameters#f9f5bf5e: field boosted_chat_id: %w", err)
		}
		p.BoostedChatID = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode premiumGiveawayParameters#f9f5bf5e: field additional_chat_ids: %w", err)
		}

		if headerLen > 0 {
			p.AdditionalChatIDs = make([]int64, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode premiumGiveawayParameters#f9f5bf5e: field additional_chat_ids: %w", err)
			}
			p.AdditionalChatIDs = append(p.AdditionalChatIDs, value)
		}
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode premiumGiveawayParameters#f9f5bf5e: field winners_selection_date: %w", err)
		}
		p.WinnersSelectionDate = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode premiumGiveawayParameters#f9f5bf5e: field only_new_members: %w", err)
		}
		p.OnlyNewMembers = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode premiumGiveawayParameters#f9f5bf5e: field has_public_winners: %w", err)
		}
		p.HasPublicWinners = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode premiumGiveawayParameters#f9f5bf5e: field country_codes: %w", err)
		}

		if headerLen > 0 {
			p.CountryCodes = make([]string, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode premiumGiveawayParameters#f9f5bf5e: field country_codes: %w", err)
			}
			p.CountryCodes = append(p.CountryCodes, value)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode premiumGiveawayParameters#f9f5bf5e: field prize_description: %w", err)
		}
		p.PrizeDescription = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (p *PremiumGiveawayParameters) EncodeTDLibJSON(b tdjson.Encoder) error {
	if p == nil {
		return fmt.Errorf("can't encode premiumGiveawayParameters#f9f5bf5e as nil")
	}
	b.ObjStart()
	b.PutID("premiumGiveawayParameters")
	b.Comma()
	b.FieldStart("boosted_chat_id")
	b.PutInt53(p.BoostedChatID)
	b.Comma()
	b.FieldStart("additional_chat_ids")
	b.ArrStart()
	for _, v := range p.AdditionalChatIDs {
		b.PutInt53(v)
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.FieldStart("winners_selection_date")
	b.PutInt32(p.WinnersSelectionDate)
	b.Comma()
	b.FieldStart("only_new_members")
	b.PutBool(p.OnlyNewMembers)
	b.Comma()
	b.FieldStart("has_public_winners")
	b.PutBool(p.HasPublicWinners)
	b.Comma()
	b.FieldStart("country_codes")
	b.ArrStart()
	for _, v := range p.CountryCodes {
		b.PutString(v)
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.FieldStart("prize_description")
	b.PutString(p.PrizeDescription)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (p *PremiumGiveawayParameters) DecodeTDLibJSON(b tdjson.Decoder) error {
	if p == nil {
		return fmt.Errorf("can't decode premiumGiveawayParameters#f9f5bf5e to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("premiumGiveawayParameters"); err != nil {
				return fmt.Errorf("unable to decode premiumGiveawayParameters#f9f5bf5e: %w", err)
			}
		case "boosted_chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode premiumGiveawayParameters#f9f5bf5e: field boosted_chat_id: %w", err)
			}
			p.BoostedChatID = value
		case "additional_chat_ids":
			if err := b.Arr(func(b tdjson.Decoder) error {
				value, err := b.Int53()
				if err != nil {
					return fmt.Errorf("unable to decode premiumGiveawayParameters#f9f5bf5e: field additional_chat_ids: %w", err)
				}
				p.AdditionalChatIDs = append(p.AdditionalChatIDs, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode premiumGiveawayParameters#f9f5bf5e: field additional_chat_ids: %w", err)
			}
		case "winners_selection_date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode premiumGiveawayParameters#f9f5bf5e: field winners_selection_date: %w", err)
			}
			p.WinnersSelectionDate = value
		case "only_new_members":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode premiumGiveawayParameters#f9f5bf5e: field only_new_members: %w", err)
			}
			p.OnlyNewMembers = value
		case "has_public_winners":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode premiumGiveawayParameters#f9f5bf5e: field has_public_winners: %w", err)
			}
			p.HasPublicWinners = value
		case "country_codes":
			if err := b.Arr(func(b tdjson.Decoder) error {
				value, err := b.String()
				if err != nil {
					return fmt.Errorf("unable to decode premiumGiveawayParameters#f9f5bf5e: field country_codes: %w", err)
				}
				p.CountryCodes = append(p.CountryCodes, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode premiumGiveawayParameters#f9f5bf5e: field country_codes: %w", err)
			}
		case "prize_description":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode premiumGiveawayParameters#f9f5bf5e: field prize_description: %w", err)
			}
			p.PrizeDescription = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetBoostedChatID returns value of BoostedChatID field.
func (p *PremiumGiveawayParameters) GetBoostedChatID() (value int64) {
	if p == nil {
		return
	}
	return p.BoostedChatID
}

// GetAdditionalChatIDs returns value of AdditionalChatIDs field.
func (p *PremiumGiveawayParameters) GetAdditionalChatIDs() (value []int64) {
	if p == nil {
		return
	}
	return p.AdditionalChatIDs
}

// GetWinnersSelectionDate returns value of WinnersSelectionDate field.
func (p *PremiumGiveawayParameters) GetWinnersSelectionDate() (value int32) {
	if p == nil {
		return
	}
	return p.WinnersSelectionDate
}

// GetOnlyNewMembers returns value of OnlyNewMembers field.
func (p *PremiumGiveawayParameters) GetOnlyNewMembers() (value bool) {
	if p == nil {
		return
	}
	return p.OnlyNewMembers
}

// GetHasPublicWinners returns value of HasPublicWinners field.
func (p *PremiumGiveawayParameters) GetHasPublicWinners() (value bool) {
	if p == nil {
		return
	}
	return p.HasPublicWinners
}

// GetCountryCodes returns value of CountryCodes field.
func (p *PremiumGiveawayParameters) GetCountryCodes() (value []string) {
	if p == nil {
		return
	}
	return p.CountryCodes
}

// GetPrizeDescription returns value of PrizeDescription field.
func (p *PremiumGiveawayParameters) GetPrizeDescription() (value string) {
	if p == nil {
		return
	}
	return p.PrizeDescription
}
