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

// ChatBoostLevelFeatures represents TL type `chatBoostLevelFeatures#fe9f29de`.
type ChatBoostLevelFeatures struct {
	// Target chat boost level
	Level int32
	// Number of stories that the chat can publish daily
	StoryPerDayCount int32
	// Number of custom emoji reactions that can be added to the list of available reactions
	CustomEmojiReactionCount int32
	// Number of custom colors for chat title
	TitleColorCount int32
	// Number of custom colors for profile photo background
	ProfileAccentColorCount int32
	// True, if custom emoji for profile background can be set
	CanSetProfileBackgroundCustomEmoji bool
	// Number of custom colors for background of empty chat photo, replies to messages and
	// link previews
	AccentColorCount int32
	// True, if custom emoji for reply header and link preview background can be set
	CanSetBackgroundCustomEmoji bool
	// True, if emoji status can be set
	CanSetEmojiStatus bool
	// Number of chat theme backgrounds that can be set as chat background
	ChatThemeBackgroundCount int32
	// True, if custom background can be set in the chat for all users
	CanSetCustomBackground bool
}

// ChatBoostLevelFeaturesTypeID is TL type id of ChatBoostLevelFeatures.
const ChatBoostLevelFeaturesTypeID = 0xfe9f29de

// Ensuring interfaces in compile-time for ChatBoostLevelFeatures.
var (
	_ bin.Encoder     = &ChatBoostLevelFeatures{}
	_ bin.Decoder     = &ChatBoostLevelFeatures{}
	_ bin.BareEncoder = &ChatBoostLevelFeatures{}
	_ bin.BareDecoder = &ChatBoostLevelFeatures{}
)

func (c *ChatBoostLevelFeatures) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Level == 0) {
		return false
	}
	if !(c.StoryPerDayCount == 0) {
		return false
	}
	if !(c.CustomEmojiReactionCount == 0) {
		return false
	}
	if !(c.TitleColorCount == 0) {
		return false
	}
	if !(c.ProfileAccentColorCount == 0) {
		return false
	}
	if !(c.CanSetProfileBackgroundCustomEmoji == false) {
		return false
	}
	if !(c.AccentColorCount == 0) {
		return false
	}
	if !(c.CanSetBackgroundCustomEmoji == false) {
		return false
	}
	if !(c.CanSetEmojiStatus == false) {
		return false
	}
	if !(c.ChatThemeBackgroundCount == 0) {
		return false
	}
	if !(c.CanSetCustomBackground == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatBoostLevelFeatures) String() string {
	if c == nil {
		return "ChatBoostLevelFeatures(nil)"
	}
	type Alias ChatBoostLevelFeatures
	return fmt.Sprintf("ChatBoostLevelFeatures%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatBoostLevelFeatures) TypeID() uint32 {
	return ChatBoostLevelFeaturesTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatBoostLevelFeatures) TypeName() string {
	return "chatBoostLevelFeatures"
}

// TypeInfo returns info about TL type.
func (c *ChatBoostLevelFeatures) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatBoostLevelFeatures",
		ID:   ChatBoostLevelFeaturesTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Level",
			SchemaName: "level",
		},
		{
			Name:       "StoryPerDayCount",
			SchemaName: "story_per_day_count",
		},
		{
			Name:       "CustomEmojiReactionCount",
			SchemaName: "custom_emoji_reaction_count",
		},
		{
			Name:       "TitleColorCount",
			SchemaName: "title_color_count",
		},
		{
			Name:       "ProfileAccentColorCount",
			SchemaName: "profile_accent_color_count",
		},
		{
			Name:       "CanSetProfileBackgroundCustomEmoji",
			SchemaName: "can_set_profile_background_custom_emoji",
		},
		{
			Name:       "AccentColorCount",
			SchemaName: "accent_color_count",
		},
		{
			Name:       "CanSetBackgroundCustomEmoji",
			SchemaName: "can_set_background_custom_emoji",
		},
		{
			Name:       "CanSetEmojiStatus",
			SchemaName: "can_set_emoji_status",
		},
		{
			Name:       "ChatThemeBackgroundCount",
			SchemaName: "chat_theme_background_count",
		},
		{
			Name:       "CanSetCustomBackground",
			SchemaName: "can_set_custom_background",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatBoostLevelFeatures) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatBoostLevelFeatures#fe9f29de as nil")
	}
	b.PutID(ChatBoostLevelFeaturesTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatBoostLevelFeatures) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatBoostLevelFeatures#fe9f29de as nil")
	}
	b.PutInt32(c.Level)
	b.PutInt32(c.StoryPerDayCount)
	b.PutInt32(c.CustomEmojiReactionCount)
	b.PutInt32(c.TitleColorCount)
	b.PutInt32(c.ProfileAccentColorCount)
	b.PutBool(c.CanSetProfileBackgroundCustomEmoji)
	b.PutInt32(c.AccentColorCount)
	b.PutBool(c.CanSetBackgroundCustomEmoji)
	b.PutBool(c.CanSetEmojiStatus)
	b.PutInt32(c.ChatThemeBackgroundCount)
	b.PutBool(c.CanSetCustomBackground)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatBoostLevelFeatures) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatBoostLevelFeatures#fe9f29de to nil")
	}
	if err := b.ConsumeID(ChatBoostLevelFeaturesTypeID); err != nil {
		return fmt.Errorf("unable to decode chatBoostLevelFeatures#fe9f29de: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatBoostLevelFeatures) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatBoostLevelFeatures#fe9f29de to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chatBoostLevelFeatures#fe9f29de: field level: %w", err)
		}
		c.Level = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chatBoostLevelFeatures#fe9f29de: field story_per_day_count: %w", err)
		}
		c.StoryPerDayCount = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chatBoostLevelFeatures#fe9f29de: field custom_emoji_reaction_count: %w", err)
		}
		c.CustomEmojiReactionCount = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chatBoostLevelFeatures#fe9f29de: field title_color_count: %w", err)
		}
		c.TitleColorCount = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chatBoostLevelFeatures#fe9f29de: field profile_accent_color_count: %w", err)
		}
		c.ProfileAccentColorCount = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatBoostLevelFeatures#fe9f29de: field can_set_profile_background_custom_emoji: %w", err)
		}
		c.CanSetProfileBackgroundCustomEmoji = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chatBoostLevelFeatures#fe9f29de: field accent_color_count: %w", err)
		}
		c.AccentColorCount = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatBoostLevelFeatures#fe9f29de: field can_set_background_custom_emoji: %w", err)
		}
		c.CanSetBackgroundCustomEmoji = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatBoostLevelFeatures#fe9f29de: field can_set_emoji_status: %w", err)
		}
		c.CanSetEmojiStatus = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chatBoostLevelFeatures#fe9f29de: field chat_theme_background_count: %w", err)
		}
		c.ChatThemeBackgroundCount = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatBoostLevelFeatures#fe9f29de: field can_set_custom_background: %w", err)
		}
		c.CanSetCustomBackground = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ChatBoostLevelFeatures) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chatBoostLevelFeatures#fe9f29de as nil")
	}
	b.ObjStart()
	b.PutID("chatBoostLevelFeatures")
	b.Comma()
	b.FieldStart("level")
	b.PutInt32(c.Level)
	b.Comma()
	b.FieldStart("story_per_day_count")
	b.PutInt32(c.StoryPerDayCount)
	b.Comma()
	b.FieldStart("custom_emoji_reaction_count")
	b.PutInt32(c.CustomEmojiReactionCount)
	b.Comma()
	b.FieldStart("title_color_count")
	b.PutInt32(c.TitleColorCount)
	b.Comma()
	b.FieldStart("profile_accent_color_count")
	b.PutInt32(c.ProfileAccentColorCount)
	b.Comma()
	b.FieldStart("can_set_profile_background_custom_emoji")
	b.PutBool(c.CanSetProfileBackgroundCustomEmoji)
	b.Comma()
	b.FieldStart("accent_color_count")
	b.PutInt32(c.AccentColorCount)
	b.Comma()
	b.FieldStart("can_set_background_custom_emoji")
	b.PutBool(c.CanSetBackgroundCustomEmoji)
	b.Comma()
	b.FieldStart("can_set_emoji_status")
	b.PutBool(c.CanSetEmojiStatus)
	b.Comma()
	b.FieldStart("chat_theme_background_count")
	b.PutInt32(c.ChatThemeBackgroundCount)
	b.Comma()
	b.FieldStart("can_set_custom_background")
	b.PutBool(c.CanSetCustomBackground)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ChatBoostLevelFeatures) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode chatBoostLevelFeatures#fe9f29de to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("chatBoostLevelFeatures"); err != nil {
				return fmt.Errorf("unable to decode chatBoostLevelFeatures#fe9f29de: %w", err)
			}
		case "level":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode chatBoostLevelFeatures#fe9f29de: field level: %w", err)
			}
			c.Level = value
		case "story_per_day_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode chatBoostLevelFeatures#fe9f29de: field story_per_day_count: %w", err)
			}
			c.StoryPerDayCount = value
		case "custom_emoji_reaction_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode chatBoostLevelFeatures#fe9f29de: field custom_emoji_reaction_count: %w", err)
			}
			c.CustomEmojiReactionCount = value
		case "title_color_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode chatBoostLevelFeatures#fe9f29de: field title_color_count: %w", err)
			}
			c.TitleColorCount = value
		case "profile_accent_color_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode chatBoostLevelFeatures#fe9f29de: field profile_accent_color_count: %w", err)
			}
			c.ProfileAccentColorCount = value
		case "can_set_profile_background_custom_emoji":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatBoostLevelFeatures#fe9f29de: field can_set_profile_background_custom_emoji: %w", err)
			}
			c.CanSetProfileBackgroundCustomEmoji = value
		case "accent_color_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode chatBoostLevelFeatures#fe9f29de: field accent_color_count: %w", err)
			}
			c.AccentColorCount = value
		case "can_set_background_custom_emoji":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatBoostLevelFeatures#fe9f29de: field can_set_background_custom_emoji: %w", err)
			}
			c.CanSetBackgroundCustomEmoji = value
		case "can_set_emoji_status":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatBoostLevelFeatures#fe9f29de: field can_set_emoji_status: %w", err)
			}
			c.CanSetEmojiStatus = value
		case "chat_theme_background_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode chatBoostLevelFeatures#fe9f29de: field chat_theme_background_count: %w", err)
			}
			c.ChatThemeBackgroundCount = value
		case "can_set_custom_background":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatBoostLevelFeatures#fe9f29de: field can_set_custom_background: %w", err)
			}
			c.CanSetCustomBackground = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetLevel returns value of Level field.
func (c *ChatBoostLevelFeatures) GetLevel() (value int32) {
	if c == nil {
		return
	}
	return c.Level
}

// GetStoryPerDayCount returns value of StoryPerDayCount field.
func (c *ChatBoostLevelFeatures) GetStoryPerDayCount() (value int32) {
	if c == nil {
		return
	}
	return c.StoryPerDayCount
}

// GetCustomEmojiReactionCount returns value of CustomEmojiReactionCount field.
func (c *ChatBoostLevelFeatures) GetCustomEmojiReactionCount() (value int32) {
	if c == nil {
		return
	}
	return c.CustomEmojiReactionCount
}

// GetTitleColorCount returns value of TitleColorCount field.
func (c *ChatBoostLevelFeatures) GetTitleColorCount() (value int32) {
	if c == nil {
		return
	}
	return c.TitleColorCount
}

// GetProfileAccentColorCount returns value of ProfileAccentColorCount field.
func (c *ChatBoostLevelFeatures) GetProfileAccentColorCount() (value int32) {
	if c == nil {
		return
	}
	return c.ProfileAccentColorCount
}

// GetCanSetProfileBackgroundCustomEmoji returns value of CanSetProfileBackgroundCustomEmoji field.
func (c *ChatBoostLevelFeatures) GetCanSetProfileBackgroundCustomEmoji() (value bool) {
	if c == nil {
		return
	}
	return c.CanSetProfileBackgroundCustomEmoji
}

// GetAccentColorCount returns value of AccentColorCount field.
func (c *ChatBoostLevelFeatures) GetAccentColorCount() (value int32) {
	if c == nil {
		return
	}
	return c.AccentColorCount
}

// GetCanSetBackgroundCustomEmoji returns value of CanSetBackgroundCustomEmoji field.
func (c *ChatBoostLevelFeatures) GetCanSetBackgroundCustomEmoji() (value bool) {
	if c == nil {
		return
	}
	return c.CanSetBackgroundCustomEmoji
}

// GetCanSetEmojiStatus returns value of CanSetEmojiStatus field.
func (c *ChatBoostLevelFeatures) GetCanSetEmojiStatus() (value bool) {
	if c == nil {
		return
	}
	return c.CanSetEmojiStatus
}

// GetChatThemeBackgroundCount returns value of ChatThemeBackgroundCount field.
func (c *ChatBoostLevelFeatures) GetChatThemeBackgroundCount() (value int32) {
	if c == nil {
		return
	}
	return c.ChatThemeBackgroundCount
}

// GetCanSetCustomBackground returns value of CanSetCustomBackground field.
func (c *ChatBoostLevelFeatures) GetCanSetCustomBackground() (value bool) {
	if c == nil {
		return
	}
	return c.CanSetCustomBackground
}
