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

// MessageEntityUnknown represents TL type `messageEntityUnknown#bb92ba95`.
type MessageEntityUnknown struct {
	// Offset field of MessageEntityUnknown.
	Offset int
	// Length field of MessageEntityUnknown.
	Length int
}

// MessageEntityUnknownTypeID is TL type id of MessageEntityUnknown.
const MessageEntityUnknownTypeID = 0xbb92ba95

// Encode implements bin.Encoder.
func (m *MessageEntityUnknown) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageEntityUnknown#bb92ba95 as nil")
	}
	b.PutID(MessageEntityUnknownTypeID)
	b.PutInt(m.Offset)
	b.PutInt(m.Length)
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageEntityUnknown) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageEntityUnknown#bb92ba95 to nil")
	}
	if err := b.ConsumeID(MessageEntityUnknownTypeID); err != nil {
		return fmt.Errorf("unable to decode messageEntityUnknown#bb92ba95: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageEntityUnknown#bb92ba95: field offset: %w", err)
		}
		m.Offset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageEntityUnknown#bb92ba95: field length: %w", err)
		}
		m.Length = value
	}
	return nil
}

// construct implements constructor of MessageEntityClass.
func (m MessageEntityUnknown) construct() MessageEntityClass { return &m }

// Ensuring interfaces in compile-time for MessageEntityUnknown.
var (
	_ bin.Encoder = &MessageEntityUnknown{}
	_ bin.Decoder = &MessageEntityUnknown{}

	_ MessageEntityClass = &MessageEntityUnknown{}
)

// MessageEntityMention represents TL type `messageEntityMention#fa04579d`.
type MessageEntityMention struct {
	// Offset field of MessageEntityMention.
	Offset int
	// Length field of MessageEntityMention.
	Length int
}

// MessageEntityMentionTypeID is TL type id of MessageEntityMention.
const MessageEntityMentionTypeID = 0xfa04579d

// Encode implements bin.Encoder.
func (m *MessageEntityMention) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageEntityMention#fa04579d as nil")
	}
	b.PutID(MessageEntityMentionTypeID)
	b.PutInt(m.Offset)
	b.PutInt(m.Length)
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageEntityMention) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageEntityMention#fa04579d to nil")
	}
	if err := b.ConsumeID(MessageEntityMentionTypeID); err != nil {
		return fmt.Errorf("unable to decode messageEntityMention#fa04579d: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageEntityMention#fa04579d: field offset: %w", err)
		}
		m.Offset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageEntityMention#fa04579d: field length: %w", err)
		}
		m.Length = value
	}
	return nil
}

// construct implements constructor of MessageEntityClass.
func (m MessageEntityMention) construct() MessageEntityClass { return &m }

// Ensuring interfaces in compile-time for MessageEntityMention.
var (
	_ bin.Encoder = &MessageEntityMention{}
	_ bin.Decoder = &MessageEntityMention{}

	_ MessageEntityClass = &MessageEntityMention{}
)

// MessageEntityHashtag represents TL type `messageEntityHashtag#6f635b0d`.
type MessageEntityHashtag struct {
	// Offset field of MessageEntityHashtag.
	Offset int
	// Length field of MessageEntityHashtag.
	Length int
}

// MessageEntityHashtagTypeID is TL type id of MessageEntityHashtag.
const MessageEntityHashtagTypeID = 0x6f635b0d

// Encode implements bin.Encoder.
func (m *MessageEntityHashtag) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageEntityHashtag#6f635b0d as nil")
	}
	b.PutID(MessageEntityHashtagTypeID)
	b.PutInt(m.Offset)
	b.PutInt(m.Length)
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageEntityHashtag) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageEntityHashtag#6f635b0d to nil")
	}
	if err := b.ConsumeID(MessageEntityHashtagTypeID); err != nil {
		return fmt.Errorf("unable to decode messageEntityHashtag#6f635b0d: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageEntityHashtag#6f635b0d: field offset: %w", err)
		}
		m.Offset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageEntityHashtag#6f635b0d: field length: %w", err)
		}
		m.Length = value
	}
	return nil
}

// construct implements constructor of MessageEntityClass.
func (m MessageEntityHashtag) construct() MessageEntityClass { return &m }

// Ensuring interfaces in compile-time for MessageEntityHashtag.
var (
	_ bin.Encoder = &MessageEntityHashtag{}
	_ bin.Decoder = &MessageEntityHashtag{}

	_ MessageEntityClass = &MessageEntityHashtag{}
)

// MessageEntityBotCommand represents TL type `messageEntityBotCommand#6cef8ac7`.
type MessageEntityBotCommand struct {
	// Offset field of MessageEntityBotCommand.
	Offset int
	// Length field of MessageEntityBotCommand.
	Length int
}

// MessageEntityBotCommandTypeID is TL type id of MessageEntityBotCommand.
const MessageEntityBotCommandTypeID = 0x6cef8ac7

// Encode implements bin.Encoder.
func (m *MessageEntityBotCommand) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageEntityBotCommand#6cef8ac7 as nil")
	}
	b.PutID(MessageEntityBotCommandTypeID)
	b.PutInt(m.Offset)
	b.PutInt(m.Length)
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageEntityBotCommand) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageEntityBotCommand#6cef8ac7 to nil")
	}
	if err := b.ConsumeID(MessageEntityBotCommandTypeID); err != nil {
		return fmt.Errorf("unable to decode messageEntityBotCommand#6cef8ac7: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageEntityBotCommand#6cef8ac7: field offset: %w", err)
		}
		m.Offset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageEntityBotCommand#6cef8ac7: field length: %w", err)
		}
		m.Length = value
	}
	return nil
}

// construct implements constructor of MessageEntityClass.
func (m MessageEntityBotCommand) construct() MessageEntityClass { return &m }

// Ensuring interfaces in compile-time for MessageEntityBotCommand.
var (
	_ bin.Encoder = &MessageEntityBotCommand{}
	_ bin.Decoder = &MessageEntityBotCommand{}

	_ MessageEntityClass = &MessageEntityBotCommand{}
)

// MessageEntityUrl represents TL type `messageEntityUrl#6ed02538`.
type MessageEntityUrl struct {
	// Offset field of MessageEntityUrl.
	Offset int
	// Length field of MessageEntityUrl.
	Length int
}

// MessageEntityUrlTypeID is TL type id of MessageEntityUrl.
const MessageEntityUrlTypeID = 0x6ed02538

// Encode implements bin.Encoder.
func (m *MessageEntityUrl) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageEntityUrl#6ed02538 as nil")
	}
	b.PutID(MessageEntityUrlTypeID)
	b.PutInt(m.Offset)
	b.PutInt(m.Length)
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageEntityUrl) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageEntityUrl#6ed02538 to nil")
	}
	if err := b.ConsumeID(MessageEntityUrlTypeID); err != nil {
		return fmt.Errorf("unable to decode messageEntityUrl#6ed02538: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageEntityUrl#6ed02538: field offset: %w", err)
		}
		m.Offset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageEntityUrl#6ed02538: field length: %w", err)
		}
		m.Length = value
	}
	return nil
}

// construct implements constructor of MessageEntityClass.
func (m MessageEntityUrl) construct() MessageEntityClass { return &m }

// Ensuring interfaces in compile-time for MessageEntityUrl.
var (
	_ bin.Encoder = &MessageEntityUrl{}
	_ bin.Decoder = &MessageEntityUrl{}

	_ MessageEntityClass = &MessageEntityUrl{}
)

// MessageEntityEmail represents TL type `messageEntityEmail#64e475c2`.
type MessageEntityEmail struct {
	// Offset field of MessageEntityEmail.
	Offset int
	// Length field of MessageEntityEmail.
	Length int
}

// MessageEntityEmailTypeID is TL type id of MessageEntityEmail.
const MessageEntityEmailTypeID = 0x64e475c2

// Encode implements bin.Encoder.
func (m *MessageEntityEmail) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageEntityEmail#64e475c2 as nil")
	}
	b.PutID(MessageEntityEmailTypeID)
	b.PutInt(m.Offset)
	b.PutInt(m.Length)
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageEntityEmail) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageEntityEmail#64e475c2 to nil")
	}
	if err := b.ConsumeID(MessageEntityEmailTypeID); err != nil {
		return fmt.Errorf("unable to decode messageEntityEmail#64e475c2: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageEntityEmail#64e475c2: field offset: %w", err)
		}
		m.Offset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageEntityEmail#64e475c2: field length: %w", err)
		}
		m.Length = value
	}
	return nil
}

// construct implements constructor of MessageEntityClass.
func (m MessageEntityEmail) construct() MessageEntityClass { return &m }

// Ensuring interfaces in compile-time for MessageEntityEmail.
var (
	_ bin.Encoder = &MessageEntityEmail{}
	_ bin.Decoder = &MessageEntityEmail{}

	_ MessageEntityClass = &MessageEntityEmail{}
)

// MessageEntityBold represents TL type `messageEntityBold#bd610bc9`.
type MessageEntityBold struct {
	// Offset field of MessageEntityBold.
	Offset int
	// Length field of MessageEntityBold.
	Length int
}

// MessageEntityBoldTypeID is TL type id of MessageEntityBold.
const MessageEntityBoldTypeID = 0xbd610bc9

// Encode implements bin.Encoder.
func (m *MessageEntityBold) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageEntityBold#bd610bc9 as nil")
	}
	b.PutID(MessageEntityBoldTypeID)
	b.PutInt(m.Offset)
	b.PutInt(m.Length)
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageEntityBold) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageEntityBold#bd610bc9 to nil")
	}
	if err := b.ConsumeID(MessageEntityBoldTypeID); err != nil {
		return fmt.Errorf("unable to decode messageEntityBold#bd610bc9: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageEntityBold#bd610bc9: field offset: %w", err)
		}
		m.Offset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageEntityBold#bd610bc9: field length: %w", err)
		}
		m.Length = value
	}
	return nil
}

// construct implements constructor of MessageEntityClass.
func (m MessageEntityBold) construct() MessageEntityClass { return &m }

// Ensuring interfaces in compile-time for MessageEntityBold.
var (
	_ bin.Encoder = &MessageEntityBold{}
	_ bin.Decoder = &MessageEntityBold{}

	_ MessageEntityClass = &MessageEntityBold{}
)

// MessageEntityItalic represents TL type `messageEntityItalic#826f8b60`.
type MessageEntityItalic struct {
	// Offset field of MessageEntityItalic.
	Offset int
	// Length field of MessageEntityItalic.
	Length int
}

// MessageEntityItalicTypeID is TL type id of MessageEntityItalic.
const MessageEntityItalicTypeID = 0x826f8b60

// Encode implements bin.Encoder.
func (m *MessageEntityItalic) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageEntityItalic#826f8b60 as nil")
	}
	b.PutID(MessageEntityItalicTypeID)
	b.PutInt(m.Offset)
	b.PutInt(m.Length)
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageEntityItalic) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageEntityItalic#826f8b60 to nil")
	}
	if err := b.ConsumeID(MessageEntityItalicTypeID); err != nil {
		return fmt.Errorf("unable to decode messageEntityItalic#826f8b60: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageEntityItalic#826f8b60: field offset: %w", err)
		}
		m.Offset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageEntityItalic#826f8b60: field length: %w", err)
		}
		m.Length = value
	}
	return nil
}

// construct implements constructor of MessageEntityClass.
func (m MessageEntityItalic) construct() MessageEntityClass { return &m }

// Ensuring interfaces in compile-time for MessageEntityItalic.
var (
	_ bin.Encoder = &MessageEntityItalic{}
	_ bin.Decoder = &MessageEntityItalic{}

	_ MessageEntityClass = &MessageEntityItalic{}
)

// MessageEntityCode represents TL type `messageEntityCode#28a20571`.
type MessageEntityCode struct {
	// Offset field of MessageEntityCode.
	Offset int
	// Length field of MessageEntityCode.
	Length int
}

// MessageEntityCodeTypeID is TL type id of MessageEntityCode.
const MessageEntityCodeTypeID = 0x28a20571

// Encode implements bin.Encoder.
func (m *MessageEntityCode) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageEntityCode#28a20571 as nil")
	}
	b.PutID(MessageEntityCodeTypeID)
	b.PutInt(m.Offset)
	b.PutInt(m.Length)
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageEntityCode) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageEntityCode#28a20571 to nil")
	}
	if err := b.ConsumeID(MessageEntityCodeTypeID); err != nil {
		return fmt.Errorf("unable to decode messageEntityCode#28a20571: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageEntityCode#28a20571: field offset: %w", err)
		}
		m.Offset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageEntityCode#28a20571: field length: %w", err)
		}
		m.Length = value
	}
	return nil
}

// construct implements constructor of MessageEntityClass.
func (m MessageEntityCode) construct() MessageEntityClass { return &m }

// Ensuring interfaces in compile-time for MessageEntityCode.
var (
	_ bin.Encoder = &MessageEntityCode{}
	_ bin.Decoder = &MessageEntityCode{}

	_ MessageEntityClass = &MessageEntityCode{}
)

// MessageEntityPre represents TL type `messageEntityPre#73924be0`.
type MessageEntityPre struct {
	// Offset field of MessageEntityPre.
	Offset int
	// Length field of MessageEntityPre.
	Length int
	// Language field of MessageEntityPre.
	Language string
}

// MessageEntityPreTypeID is TL type id of MessageEntityPre.
const MessageEntityPreTypeID = 0x73924be0

// Encode implements bin.Encoder.
func (m *MessageEntityPre) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageEntityPre#73924be0 as nil")
	}
	b.PutID(MessageEntityPreTypeID)
	b.PutInt(m.Offset)
	b.PutInt(m.Length)
	b.PutString(m.Language)
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageEntityPre) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageEntityPre#73924be0 to nil")
	}
	if err := b.ConsumeID(MessageEntityPreTypeID); err != nil {
		return fmt.Errorf("unable to decode messageEntityPre#73924be0: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageEntityPre#73924be0: field offset: %w", err)
		}
		m.Offset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageEntityPre#73924be0: field length: %w", err)
		}
		m.Length = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messageEntityPre#73924be0: field language: %w", err)
		}
		m.Language = value
	}
	return nil
}

// construct implements constructor of MessageEntityClass.
func (m MessageEntityPre) construct() MessageEntityClass { return &m }

// Ensuring interfaces in compile-time for MessageEntityPre.
var (
	_ bin.Encoder = &MessageEntityPre{}
	_ bin.Decoder = &MessageEntityPre{}

	_ MessageEntityClass = &MessageEntityPre{}
)

// MessageEntityTextUrl represents TL type `messageEntityTextUrl#76a6d327`.
type MessageEntityTextUrl struct {
	// Offset field of MessageEntityTextUrl.
	Offset int
	// Length field of MessageEntityTextUrl.
	Length int
	// URL field of MessageEntityTextUrl.
	URL string
}

// MessageEntityTextUrlTypeID is TL type id of MessageEntityTextUrl.
const MessageEntityTextUrlTypeID = 0x76a6d327

// Encode implements bin.Encoder.
func (m *MessageEntityTextUrl) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageEntityTextUrl#76a6d327 as nil")
	}
	b.PutID(MessageEntityTextUrlTypeID)
	b.PutInt(m.Offset)
	b.PutInt(m.Length)
	b.PutString(m.URL)
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageEntityTextUrl) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageEntityTextUrl#76a6d327 to nil")
	}
	if err := b.ConsumeID(MessageEntityTextUrlTypeID); err != nil {
		return fmt.Errorf("unable to decode messageEntityTextUrl#76a6d327: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageEntityTextUrl#76a6d327: field offset: %w", err)
		}
		m.Offset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageEntityTextUrl#76a6d327: field length: %w", err)
		}
		m.Length = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messageEntityTextUrl#76a6d327: field url: %w", err)
		}
		m.URL = value
	}
	return nil
}

// construct implements constructor of MessageEntityClass.
func (m MessageEntityTextUrl) construct() MessageEntityClass { return &m }

// Ensuring interfaces in compile-time for MessageEntityTextUrl.
var (
	_ bin.Encoder = &MessageEntityTextUrl{}
	_ bin.Decoder = &MessageEntityTextUrl{}

	_ MessageEntityClass = &MessageEntityTextUrl{}
)

// MessageEntityMentionName represents TL type `messageEntityMentionName#352dca58`.
type MessageEntityMentionName struct {
	// Offset field of MessageEntityMentionName.
	Offset int
	// Length field of MessageEntityMentionName.
	Length int
	// UserID field of MessageEntityMentionName.
	UserID int
}

// MessageEntityMentionNameTypeID is TL type id of MessageEntityMentionName.
const MessageEntityMentionNameTypeID = 0x352dca58

// Encode implements bin.Encoder.
func (m *MessageEntityMentionName) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageEntityMentionName#352dca58 as nil")
	}
	b.PutID(MessageEntityMentionNameTypeID)
	b.PutInt(m.Offset)
	b.PutInt(m.Length)
	b.PutInt(m.UserID)
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageEntityMentionName) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageEntityMentionName#352dca58 to nil")
	}
	if err := b.ConsumeID(MessageEntityMentionNameTypeID); err != nil {
		return fmt.Errorf("unable to decode messageEntityMentionName#352dca58: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageEntityMentionName#352dca58: field offset: %w", err)
		}
		m.Offset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageEntityMentionName#352dca58: field length: %w", err)
		}
		m.Length = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageEntityMentionName#352dca58: field user_id: %w", err)
		}
		m.UserID = value
	}
	return nil
}

// construct implements constructor of MessageEntityClass.
func (m MessageEntityMentionName) construct() MessageEntityClass { return &m }

// Ensuring interfaces in compile-time for MessageEntityMentionName.
var (
	_ bin.Encoder = &MessageEntityMentionName{}
	_ bin.Decoder = &MessageEntityMentionName{}

	_ MessageEntityClass = &MessageEntityMentionName{}
)

// InputMessageEntityMentionName represents TL type `inputMessageEntityMentionName#208e68c9`.
type InputMessageEntityMentionName struct {
	// Offset field of InputMessageEntityMentionName.
	Offset int
	// Length field of InputMessageEntityMentionName.
	Length int
	// UserID field of InputMessageEntityMentionName.
	UserID InputUserClass
}

// InputMessageEntityMentionNameTypeID is TL type id of InputMessageEntityMentionName.
const InputMessageEntityMentionNameTypeID = 0x208e68c9

// Encode implements bin.Encoder.
func (i *InputMessageEntityMentionName) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputMessageEntityMentionName#208e68c9 as nil")
	}
	b.PutID(InputMessageEntityMentionNameTypeID)
	b.PutInt(i.Offset)
	b.PutInt(i.Length)
	if i.UserID == nil {
		return fmt.Errorf("unable to encode inputMessageEntityMentionName#208e68c9: field user_id is nil")
	}
	if err := i.UserID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputMessageEntityMentionName#208e68c9: field user_id: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputMessageEntityMentionName) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputMessageEntityMentionName#208e68c9 to nil")
	}
	if err := b.ConsumeID(InputMessageEntityMentionNameTypeID); err != nil {
		return fmt.Errorf("unable to decode inputMessageEntityMentionName#208e68c9: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputMessageEntityMentionName#208e68c9: field offset: %w", err)
		}
		i.Offset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputMessageEntityMentionName#208e68c9: field length: %w", err)
		}
		i.Length = value
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputMessageEntityMentionName#208e68c9: field user_id: %w", err)
		}
		i.UserID = value
	}
	return nil
}

// construct implements constructor of MessageEntityClass.
func (i InputMessageEntityMentionName) construct() MessageEntityClass { return &i }

// Ensuring interfaces in compile-time for InputMessageEntityMentionName.
var (
	_ bin.Encoder = &InputMessageEntityMentionName{}
	_ bin.Decoder = &InputMessageEntityMentionName{}

	_ MessageEntityClass = &InputMessageEntityMentionName{}
)

// MessageEntityPhone represents TL type `messageEntityPhone#9b69e34b`.
type MessageEntityPhone struct {
	// Offset field of MessageEntityPhone.
	Offset int
	// Length field of MessageEntityPhone.
	Length int
}

// MessageEntityPhoneTypeID is TL type id of MessageEntityPhone.
const MessageEntityPhoneTypeID = 0x9b69e34b

// Encode implements bin.Encoder.
func (m *MessageEntityPhone) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageEntityPhone#9b69e34b as nil")
	}
	b.PutID(MessageEntityPhoneTypeID)
	b.PutInt(m.Offset)
	b.PutInt(m.Length)
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageEntityPhone) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageEntityPhone#9b69e34b to nil")
	}
	if err := b.ConsumeID(MessageEntityPhoneTypeID); err != nil {
		return fmt.Errorf("unable to decode messageEntityPhone#9b69e34b: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageEntityPhone#9b69e34b: field offset: %w", err)
		}
		m.Offset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageEntityPhone#9b69e34b: field length: %w", err)
		}
		m.Length = value
	}
	return nil
}

// construct implements constructor of MessageEntityClass.
func (m MessageEntityPhone) construct() MessageEntityClass { return &m }

// Ensuring interfaces in compile-time for MessageEntityPhone.
var (
	_ bin.Encoder = &MessageEntityPhone{}
	_ bin.Decoder = &MessageEntityPhone{}

	_ MessageEntityClass = &MessageEntityPhone{}
)

// MessageEntityCashtag represents TL type `messageEntityCashtag#4c4e743f`.
type MessageEntityCashtag struct {
	// Offset field of MessageEntityCashtag.
	Offset int
	// Length field of MessageEntityCashtag.
	Length int
}

// MessageEntityCashtagTypeID is TL type id of MessageEntityCashtag.
const MessageEntityCashtagTypeID = 0x4c4e743f

// Encode implements bin.Encoder.
func (m *MessageEntityCashtag) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageEntityCashtag#4c4e743f as nil")
	}
	b.PutID(MessageEntityCashtagTypeID)
	b.PutInt(m.Offset)
	b.PutInt(m.Length)
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageEntityCashtag) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageEntityCashtag#4c4e743f to nil")
	}
	if err := b.ConsumeID(MessageEntityCashtagTypeID); err != nil {
		return fmt.Errorf("unable to decode messageEntityCashtag#4c4e743f: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageEntityCashtag#4c4e743f: field offset: %w", err)
		}
		m.Offset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageEntityCashtag#4c4e743f: field length: %w", err)
		}
		m.Length = value
	}
	return nil
}

// construct implements constructor of MessageEntityClass.
func (m MessageEntityCashtag) construct() MessageEntityClass { return &m }

// Ensuring interfaces in compile-time for MessageEntityCashtag.
var (
	_ bin.Encoder = &MessageEntityCashtag{}
	_ bin.Decoder = &MessageEntityCashtag{}

	_ MessageEntityClass = &MessageEntityCashtag{}
)

// MessageEntityUnderline represents TL type `messageEntityUnderline#9c4e7e8b`.
type MessageEntityUnderline struct {
	// Offset field of MessageEntityUnderline.
	Offset int
	// Length field of MessageEntityUnderline.
	Length int
}

// MessageEntityUnderlineTypeID is TL type id of MessageEntityUnderline.
const MessageEntityUnderlineTypeID = 0x9c4e7e8b

// Encode implements bin.Encoder.
func (m *MessageEntityUnderline) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageEntityUnderline#9c4e7e8b as nil")
	}
	b.PutID(MessageEntityUnderlineTypeID)
	b.PutInt(m.Offset)
	b.PutInt(m.Length)
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageEntityUnderline) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageEntityUnderline#9c4e7e8b to nil")
	}
	if err := b.ConsumeID(MessageEntityUnderlineTypeID); err != nil {
		return fmt.Errorf("unable to decode messageEntityUnderline#9c4e7e8b: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageEntityUnderline#9c4e7e8b: field offset: %w", err)
		}
		m.Offset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageEntityUnderline#9c4e7e8b: field length: %w", err)
		}
		m.Length = value
	}
	return nil
}

// construct implements constructor of MessageEntityClass.
func (m MessageEntityUnderline) construct() MessageEntityClass { return &m }

// Ensuring interfaces in compile-time for MessageEntityUnderline.
var (
	_ bin.Encoder = &MessageEntityUnderline{}
	_ bin.Decoder = &MessageEntityUnderline{}

	_ MessageEntityClass = &MessageEntityUnderline{}
)

// MessageEntityStrike represents TL type `messageEntityStrike#bf0693d4`.
type MessageEntityStrike struct {
	// Offset field of MessageEntityStrike.
	Offset int
	// Length field of MessageEntityStrike.
	Length int
}

// MessageEntityStrikeTypeID is TL type id of MessageEntityStrike.
const MessageEntityStrikeTypeID = 0xbf0693d4

// Encode implements bin.Encoder.
func (m *MessageEntityStrike) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageEntityStrike#bf0693d4 as nil")
	}
	b.PutID(MessageEntityStrikeTypeID)
	b.PutInt(m.Offset)
	b.PutInt(m.Length)
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageEntityStrike) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageEntityStrike#bf0693d4 to nil")
	}
	if err := b.ConsumeID(MessageEntityStrikeTypeID); err != nil {
		return fmt.Errorf("unable to decode messageEntityStrike#bf0693d4: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageEntityStrike#bf0693d4: field offset: %w", err)
		}
		m.Offset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageEntityStrike#bf0693d4: field length: %w", err)
		}
		m.Length = value
	}
	return nil
}

// construct implements constructor of MessageEntityClass.
func (m MessageEntityStrike) construct() MessageEntityClass { return &m }

// Ensuring interfaces in compile-time for MessageEntityStrike.
var (
	_ bin.Encoder = &MessageEntityStrike{}
	_ bin.Decoder = &MessageEntityStrike{}

	_ MessageEntityClass = &MessageEntityStrike{}
)

// MessageEntityBlockquote represents TL type `messageEntityBlockquote#20df5d0`.
type MessageEntityBlockquote struct {
	// Offset field of MessageEntityBlockquote.
	Offset int
	// Length field of MessageEntityBlockquote.
	Length int
}

// MessageEntityBlockquoteTypeID is TL type id of MessageEntityBlockquote.
const MessageEntityBlockquoteTypeID = 0x20df5d0

// Encode implements bin.Encoder.
func (m *MessageEntityBlockquote) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageEntityBlockquote#20df5d0 as nil")
	}
	b.PutID(MessageEntityBlockquoteTypeID)
	b.PutInt(m.Offset)
	b.PutInt(m.Length)
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageEntityBlockquote) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageEntityBlockquote#20df5d0 to nil")
	}
	if err := b.ConsumeID(MessageEntityBlockquoteTypeID); err != nil {
		return fmt.Errorf("unable to decode messageEntityBlockquote#20df5d0: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageEntityBlockquote#20df5d0: field offset: %w", err)
		}
		m.Offset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageEntityBlockquote#20df5d0: field length: %w", err)
		}
		m.Length = value
	}
	return nil
}

// construct implements constructor of MessageEntityClass.
func (m MessageEntityBlockquote) construct() MessageEntityClass { return &m }

// Ensuring interfaces in compile-time for MessageEntityBlockquote.
var (
	_ bin.Encoder = &MessageEntityBlockquote{}
	_ bin.Decoder = &MessageEntityBlockquote{}

	_ MessageEntityClass = &MessageEntityBlockquote{}
)

// MessageEntityBankCard represents TL type `messageEntityBankCard#761e6af4`.
type MessageEntityBankCard struct {
	// Offset field of MessageEntityBankCard.
	Offset int
	// Length field of MessageEntityBankCard.
	Length int
}

// MessageEntityBankCardTypeID is TL type id of MessageEntityBankCard.
const MessageEntityBankCardTypeID = 0x761e6af4

// Encode implements bin.Encoder.
func (m *MessageEntityBankCard) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageEntityBankCard#761e6af4 as nil")
	}
	b.PutID(MessageEntityBankCardTypeID)
	b.PutInt(m.Offset)
	b.PutInt(m.Length)
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageEntityBankCard) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageEntityBankCard#761e6af4 to nil")
	}
	if err := b.ConsumeID(MessageEntityBankCardTypeID); err != nil {
		return fmt.Errorf("unable to decode messageEntityBankCard#761e6af4: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageEntityBankCard#761e6af4: field offset: %w", err)
		}
		m.Offset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageEntityBankCard#761e6af4: field length: %w", err)
		}
		m.Length = value
	}
	return nil
}

// construct implements constructor of MessageEntityClass.
func (m MessageEntityBankCard) construct() MessageEntityClass { return &m }

// Ensuring interfaces in compile-time for MessageEntityBankCard.
var (
	_ bin.Encoder = &MessageEntityBankCard{}
	_ bin.Decoder = &MessageEntityBankCard{}

	_ MessageEntityClass = &MessageEntityBankCard{}
)

// MessageEntityClass represents MessageEntity generic type.
//
// Example:
//  g, err := DecodeMessageEntity(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *MessageEntityUnknown: // messageEntityUnknown#bb92ba95
//  case *MessageEntityMention: // messageEntityMention#fa04579d
//  case *MessageEntityHashtag: // messageEntityHashtag#6f635b0d
//  case *MessageEntityBotCommand: // messageEntityBotCommand#6cef8ac7
//  case *MessageEntityUrl: // messageEntityUrl#6ed02538
//  case *MessageEntityEmail: // messageEntityEmail#64e475c2
//  case *MessageEntityBold: // messageEntityBold#bd610bc9
//  case *MessageEntityItalic: // messageEntityItalic#826f8b60
//  case *MessageEntityCode: // messageEntityCode#28a20571
//  case *MessageEntityPre: // messageEntityPre#73924be0
//  case *MessageEntityTextUrl: // messageEntityTextUrl#76a6d327
//  case *MessageEntityMentionName: // messageEntityMentionName#352dca58
//  case *InputMessageEntityMentionName: // inputMessageEntityMentionName#208e68c9
//  case *MessageEntityPhone: // messageEntityPhone#9b69e34b
//  case *MessageEntityCashtag: // messageEntityCashtag#4c4e743f
//  case *MessageEntityUnderline: // messageEntityUnderline#9c4e7e8b
//  case *MessageEntityStrike: // messageEntityStrike#bf0693d4
//  case *MessageEntityBlockquote: // messageEntityBlockquote#20df5d0
//  case *MessageEntityBankCard: // messageEntityBankCard#761e6af4
//  default: panic(v)
//  }
type MessageEntityClass interface {
	bin.Encoder
	bin.Decoder
	construct() MessageEntityClass
}

// DecodeMessageEntity implements binary de-serialization for MessageEntityClass.
func DecodeMessageEntity(buf *bin.Buffer) (MessageEntityClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case MessageEntityUnknownTypeID:
		// Decoding messageEntityUnknown#bb92ba95.
		v := MessageEntityUnknown{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageEntityClass: %w", err)
		}
		return &v, nil
	case MessageEntityMentionTypeID:
		// Decoding messageEntityMention#fa04579d.
		v := MessageEntityMention{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageEntityClass: %w", err)
		}
		return &v, nil
	case MessageEntityHashtagTypeID:
		// Decoding messageEntityHashtag#6f635b0d.
		v := MessageEntityHashtag{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageEntityClass: %w", err)
		}
		return &v, nil
	case MessageEntityBotCommandTypeID:
		// Decoding messageEntityBotCommand#6cef8ac7.
		v := MessageEntityBotCommand{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageEntityClass: %w", err)
		}
		return &v, nil
	case MessageEntityUrlTypeID:
		// Decoding messageEntityUrl#6ed02538.
		v := MessageEntityUrl{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageEntityClass: %w", err)
		}
		return &v, nil
	case MessageEntityEmailTypeID:
		// Decoding messageEntityEmail#64e475c2.
		v := MessageEntityEmail{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageEntityClass: %w", err)
		}
		return &v, nil
	case MessageEntityBoldTypeID:
		// Decoding messageEntityBold#bd610bc9.
		v := MessageEntityBold{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageEntityClass: %w", err)
		}
		return &v, nil
	case MessageEntityItalicTypeID:
		// Decoding messageEntityItalic#826f8b60.
		v := MessageEntityItalic{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageEntityClass: %w", err)
		}
		return &v, nil
	case MessageEntityCodeTypeID:
		// Decoding messageEntityCode#28a20571.
		v := MessageEntityCode{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageEntityClass: %w", err)
		}
		return &v, nil
	case MessageEntityPreTypeID:
		// Decoding messageEntityPre#73924be0.
		v := MessageEntityPre{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageEntityClass: %w", err)
		}
		return &v, nil
	case MessageEntityTextUrlTypeID:
		// Decoding messageEntityTextUrl#76a6d327.
		v := MessageEntityTextUrl{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageEntityClass: %w", err)
		}
		return &v, nil
	case MessageEntityMentionNameTypeID:
		// Decoding messageEntityMentionName#352dca58.
		v := MessageEntityMentionName{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageEntityClass: %w", err)
		}
		return &v, nil
	case InputMessageEntityMentionNameTypeID:
		// Decoding inputMessageEntityMentionName#208e68c9.
		v := InputMessageEntityMentionName{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageEntityClass: %w", err)
		}
		return &v, nil
	case MessageEntityPhoneTypeID:
		// Decoding messageEntityPhone#9b69e34b.
		v := MessageEntityPhone{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageEntityClass: %w", err)
		}
		return &v, nil
	case MessageEntityCashtagTypeID:
		// Decoding messageEntityCashtag#4c4e743f.
		v := MessageEntityCashtag{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageEntityClass: %w", err)
		}
		return &v, nil
	case MessageEntityUnderlineTypeID:
		// Decoding messageEntityUnderline#9c4e7e8b.
		v := MessageEntityUnderline{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageEntityClass: %w", err)
		}
		return &v, nil
	case MessageEntityStrikeTypeID:
		// Decoding messageEntityStrike#bf0693d4.
		v := MessageEntityStrike{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageEntityClass: %w", err)
		}
		return &v, nil
	case MessageEntityBlockquoteTypeID:
		// Decoding messageEntityBlockquote#20df5d0.
		v := MessageEntityBlockquote{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageEntityClass: %w", err)
		}
		return &v, nil
	case MessageEntityBankCardTypeID:
		// Decoding messageEntityBankCard#761e6af4.
		v := MessageEntityBankCard{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageEntityClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode MessageEntityClass: %w", bin.NewUnexpectedID(id))
	}
}
