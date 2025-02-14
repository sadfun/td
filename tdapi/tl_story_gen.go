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

// Story represents TL type `story#8e358ca2`.
type Story struct {
	// Unique story identifier among stories of the given sender
	ID int32
	// Identifier of the chat that posted the story
	SenderChatID int64
	// Point in time (Unix timestamp) when the story was published
	Date int32
	// True, if the story is being sent by the current user
	IsBeingSent bool
	// True, if the story is being edited by the current user
	IsBeingEdited bool
	// True, if the story was edited
	IsEdited bool
	// True, if the story is saved in the sender's profile and will be available there after
	// expiration
	IsPinned bool
	// True, if the story is visible only for the current user
	IsVisibleOnlyForSelf bool
	// True, if the story can be deleted
	CanBeDeleted bool
	// True, if the story can be edited
	CanBeEdited bool
	// True, if the story can be forwarded as a message. Otherwise, screenshots and saving of
	// the story content must be also forbidden
	CanBeForwarded bool
	// True, if the story can be replied in the chat with the story sender
	CanBeReplied bool
	// True, if the story's is_pinned value can be changed
	CanToggleIsPinned bool
	// True, if the story statistics are available through getStoryStatistics
	CanGetStatistics bool
	// True, if interactions with the story can be received through getStoryInteractions
	CanGetInteractions bool
	// True, if users viewed the story can't be received, because the story has expired more
	// than getOption("story_viewers_expiration_delay") seconds ago
	HasExpiredViewers bool
	// Information about the original story; may be null if the story wasn't reposted
	RepostInfo StoryRepostInfo
	// Information about interactions with the story; may be null if the story isn't owned or
	// there were no interactions
	InteractionInfo StoryInteractionInfo
	// Type of the chosen reaction; may be null if none
	ChosenReactionType ReactionTypeClass
	// Privacy rules affecting story visibility; may be approximate for non-owned stories
	PrivacySettings StoryPrivacySettingsClass
	// Content of the story
	Content StoryContentClass
	// Clickable areas to be shown on the story content
	Areas []StoryArea
	// Caption of the story
	Caption FormattedText
}

// StoryTypeID is TL type id of Story.
const StoryTypeID = 0x8e358ca2

// Ensuring interfaces in compile-time for Story.
var (
	_ bin.Encoder     = &Story{}
	_ bin.Decoder     = &Story{}
	_ bin.BareEncoder = &Story{}
	_ bin.BareDecoder = &Story{}
)

func (s *Story) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.ID == 0) {
		return false
	}
	if !(s.SenderChatID == 0) {
		return false
	}
	if !(s.Date == 0) {
		return false
	}
	if !(s.IsBeingSent == false) {
		return false
	}
	if !(s.IsBeingEdited == false) {
		return false
	}
	if !(s.IsEdited == false) {
		return false
	}
	if !(s.IsPinned == false) {
		return false
	}
	if !(s.IsVisibleOnlyForSelf == false) {
		return false
	}
	if !(s.CanBeDeleted == false) {
		return false
	}
	if !(s.CanBeEdited == false) {
		return false
	}
	if !(s.CanBeForwarded == false) {
		return false
	}
	if !(s.CanBeReplied == false) {
		return false
	}
	if !(s.CanToggleIsPinned == false) {
		return false
	}
	if !(s.CanGetStatistics == false) {
		return false
	}
	if !(s.CanGetInteractions == false) {
		return false
	}
	if !(s.HasExpiredViewers == false) {
		return false
	}
	if !(s.RepostInfo.Zero()) {
		return false
	}
	if !(s.InteractionInfo.Zero()) {
		return false
	}
	if !(s.ChosenReactionType == nil) {
		return false
	}
	if !(s.PrivacySettings == nil) {
		return false
	}
	if !(s.Content == nil) {
		return false
	}
	if !(s.Areas == nil) {
		return false
	}
	if !(s.Caption.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *Story) String() string {
	if s == nil {
		return "Story(nil)"
	}
	type Alias Story
	return fmt.Sprintf("Story%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Story) TypeID() uint32 {
	return StoryTypeID
}

// TypeName returns name of type in TL schema.
func (*Story) TypeName() string {
	return "story"
}

// TypeInfo returns info about TL type.
func (s *Story) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "story",
		ID:   StoryTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "SenderChatID",
			SchemaName: "sender_chat_id",
		},
		{
			Name:       "Date",
			SchemaName: "date",
		},
		{
			Name:       "IsBeingSent",
			SchemaName: "is_being_sent",
		},
		{
			Name:       "IsBeingEdited",
			SchemaName: "is_being_edited",
		},
		{
			Name:       "IsEdited",
			SchemaName: "is_edited",
		},
		{
			Name:       "IsPinned",
			SchemaName: "is_pinned",
		},
		{
			Name:       "IsVisibleOnlyForSelf",
			SchemaName: "is_visible_only_for_self",
		},
		{
			Name:       "CanBeDeleted",
			SchemaName: "can_be_deleted",
		},
		{
			Name:       "CanBeEdited",
			SchemaName: "can_be_edited",
		},
		{
			Name:       "CanBeForwarded",
			SchemaName: "can_be_forwarded",
		},
		{
			Name:       "CanBeReplied",
			SchemaName: "can_be_replied",
		},
		{
			Name:       "CanToggleIsPinned",
			SchemaName: "can_toggle_is_pinned",
		},
		{
			Name:       "CanGetStatistics",
			SchemaName: "can_get_statistics",
		},
		{
			Name:       "CanGetInteractions",
			SchemaName: "can_get_interactions",
		},
		{
			Name:       "HasExpiredViewers",
			SchemaName: "has_expired_viewers",
		},
		{
			Name:       "RepostInfo",
			SchemaName: "repost_info",
		},
		{
			Name:       "InteractionInfo",
			SchemaName: "interaction_info",
		},
		{
			Name:       "ChosenReactionType",
			SchemaName: "chosen_reaction_type",
		},
		{
			Name:       "PrivacySettings",
			SchemaName: "privacy_settings",
		},
		{
			Name:       "Content",
			SchemaName: "content",
		},
		{
			Name:       "Areas",
			SchemaName: "areas",
		},
		{
			Name:       "Caption",
			SchemaName: "caption",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *Story) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode story#8e358ca2 as nil")
	}
	b.PutID(StoryTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *Story) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode story#8e358ca2 as nil")
	}
	b.PutInt32(s.ID)
	b.PutInt53(s.SenderChatID)
	b.PutInt32(s.Date)
	b.PutBool(s.IsBeingSent)
	b.PutBool(s.IsBeingEdited)
	b.PutBool(s.IsEdited)
	b.PutBool(s.IsPinned)
	b.PutBool(s.IsVisibleOnlyForSelf)
	b.PutBool(s.CanBeDeleted)
	b.PutBool(s.CanBeEdited)
	b.PutBool(s.CanBeForwarded)
	b.PutBool(s.CanBeReplied)
	b.PutBool(s.CanToggleIsPinned)
	b.PutBool(s.CanGetStatistics)
	b.PutBool(s.CanGetInteractions)
	b.PutBool(s.HasExpiredViewers)
	if err := s.RepostInfo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode story#8e358ca2: field repost_info: %w", err)
	}
	if err := s.InteractionInfo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode story#8e358ca2: field interaction_info: %w", err)
	}
	if s.ChosenReactionType == nil {
		return fmt.Errorf("unable to encode story#8e358ca2: field chosen_reaction_type is nil")
	}
	if err := s.ChosenReactionType.Encode(b); err != nil {
		return fmt.Errorf("unable to encode story#8e358ca2: field chosen_reaction_type: %w", err)
	}
	if s.PrivacySettings == nil {
		return fmt.Errorf("unable to encode story#8e358ca2: field privacy_settings is nil")
	}
	if err := s.PrivacySettings.Encode(b); err != nil {
		return fmt.Errorf("unable to encode story#8e358ca2: field privacy_settings: %w", err)
	}
	if s.Content == nil {
		return fmt.Errorf("unable to encode story#8e358ca2: field content is nil")
	}
	if err := s.Content.Encode(b); err != nil {
		return fmt.Errorf("unable to encode story#8e358ca2: field content: %w", err)
	}
	b.PutInt(len(s.Areas))
	for idx, v := range s.Areas {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare story#8e358ca2: field areas element with index %d: %w", idx, err)
		}
	}
	if err := s.Caption.Encode(b); err != nil {
		return fmt.Errorf("unable to encode story#8e358ca2: field caption: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *Story) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode story#8e358ca2 to nil")
	}
	if err := b.ConsumeID(StoryTypeID); err != nil {
		return fmt.Errorf("unable to decode story#8e358ca2: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *Story) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode story#8e358ca2 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode story#8e358ca2: field id: %w", err)
		}
		s.ID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode story#8e358ca2: field sender_chat_id: %w", err)
		}
		s.SenderChatID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode story#8e358ca2: field date: %w", err)
		}
		s.Date = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode story#8e358ca2: field is_being_sent: %w", err)
		}
		s.IsBeingSent = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode story#8e358ca2: field is_being_edited: %w", err)
		}
		s.IsBeingEdited = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode story#8e358ca2: field is_edited: %w", err)
		}
		s.IsEdited = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode story#8e358ca2: field is_pinned: %w", err)
		}
		s.IsPinned = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode story#8e358ca2: field is_visible_only_for_self: %w", err)
		}
		s.IsVisibleOnlyForSelf = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode story#8e358ca2: field can_be_deleted: %w", err)
		}
		s.CanBeDeleted = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode story#8e358ca2: field can_be_edited: %w", err)
		}
		s.CanBeEdited = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode story#8e358ca2: field can_be_forwarded: %w", err)
		}
		s.CanBeForwarded = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode story#8e358ca2: field can_be_replied: %w", err)
		}
		s.CanBeReplied = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode story#8e358ca2: field can_toggle_is_pinned: %w", err)
		}
		s.CanToggleIsPinned = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode story#8e358ca2: field can_get_statistics: %w", err)
		}
		s.CanGetStatistics = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode story#8e358ca2: field can_get_interactions: %w", err)
		}
		s.CanGetInteractions = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode story#8e358ca2: field has_expired_viewers: %w", err)
		}
		s.HasExpiredViewers = value
	}
	{
		if err := s.RepostInfo.Decode(b); err != nil {
			return fmt.Errorf("unable to decode story#8e358ca2: field repost_info: %w", err)
		}
	}
	{
		if err := s.InteractionInfo.Decode(b); err != nil {
			return fmt.Errorf("unable to decode story#8e358ca2: field interaction_info: %w", err)
		}
	}
	{
		value, err := DecodeReactionType(b)
		if err != nil {
			return fmt.Errorf("unable to decode story#8e358ca2: field chosen_reaction_type: %w", err)
		}
		s.ChosenReactionType = value
	}
	{
		value, err := DecodeStoryPrivacySettings(b)
		if err != nil {
			return fmt.Errorf("unable to decode story#8e358ca2: field privacy_settings: %w", err)
		}
		s.PrivacySettings = value
	}
	{
		value, err := DecodeStoryContent(b)
		if err != nil {
			return fmt.Errorf("unable to decode story#8e358ca2: field content: %w", err)
		}
		s.Content = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode story#8e358ca2: field areas: %w", err)
		}

		if headerLen > 0 {
			s.Areas = make([]StoryArea, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value StoryArea
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare story#8e358ca2: field areas: %w", err)
			}
			s.Areas = append(s.Areas, value)
		}
	}
	{
		if err := s.Caption.Decode(b); err != nil {
			return fmt.Errorf("unable to decode story#8e358ca2: field caption: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *Story) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode story#8e358ca2 as nil")
	}
	b.ObjStart()
	b.PutID("story")
	b.Comma()
	b.FieldStart("id")
	b.PutInt32(s.ID)
	b.Comma()
	b.FieldStart("sender_chat_id")
	b.PutInt53(s.SenderChatID)
	b.Comma()
	b.FieldStart("date")
	b.PutInt32(s.Date)
	b.Comma()
	b.FieldStart("is_being_sent")
	b.PutBool(s.IsBeingSent)
	b.Comma()
	b.FieldStart("is_being_edited")
	b.PutBool(s.IsBeingEdited)
	b.Comma()
	b.FieldStart("is_edited")
	b.PutBool(s.IsEdited)
	b.Comma()
	b.FieldStart("is_pinned")
	b.PutBool(s.IsPinned)
	b.Comma()
	b.FieldStart("is_visible_only_for_self")
	b.PutBool(s.IsVisibleOnlyForSelf)
	b.Comma()
	b.FieldStart("can_be_deleted")
	b.PutBool(s.CanBeDeleted)
	b.Comma()
	b.FieldStart("can_be_edited")
	b.PutBool(s.CanBeEdited)
	b.Comma()
	b.FieldStart("can_be_forwarded")
	b.PutBool(s.CanBeForwarded)
	b.Comma()
	b.FieldStart("can_be_replied")
	b.PutBool(s.CanBeReplied)
	b.Comma()
	b.FieldStart("can_toggle_is_pinned")
	b.PutBool(s.CanToggleIsPinned)
	b.Comma()
	b.FieldStart("can_get_statistics")
	b.PutBool(s.CanGetStatistics)
	b.Comma()
	b.FieldStart("can_get_interactions")
	b.PutBool(s.CanGetInteractions)
	b.Comma()
	b.FieldStart("has_expired_viewers")
	b.PutBool(s.HasExpiredViewers)
	b.Comma()
	b.FieldStart("repost_info")
	if err := s.RepostInfo.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode story#8e358ca2: field repost_info: %w", err)
	}
	b.Comma()
	b.FieldStart("interaction_info")
	if err := s.InteractionInfo.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode story#8e358ca2: field interaction_info: %w", err)
	}
	b.Comma()
	b.FieldStart("chosen_reaction_type")
	if s.ChosenReactionType == nil {
		return fmt.Errorf("unable to encode story#8e358ca2: field chosen_reaction_type is nil")
	}
	if err := s.ChosenReactionType.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode story#8e358ca2: field chosen_reaction_type: %w", err)
	}
	b.Comma()
	b.FieldStart("privacy_settings")
	if s.PrivacySettings == nil {
		return fmt.Errorf("unable to encode story#8e358ca2: field privacy_settings is nil")
	}
	if err := s.PrivacySettings.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode story#8e358ca2: field privacy_settings: %w", err)
	}
	b.Comma()
	b.FieldStart("content")
	if s.Content == nil {
		return fmt.Errorf("unable to encode story#8e358ca2: field content is nil")
	}
	if err := s.Content.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode story#8e358ca2: field content: %w", err)
	}
	b.Comma()
	b.FieldStart("areas")
	b.ArrStart()
	for idx, v := range s.Areas {
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode story#8e358ca2: field areas element with index %d: %w", idx, err)
		}
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.FieldStart("caption")
	if err := s.Caption.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode story#8e358ca2: field caption: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *Story) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode story#8e358ca2 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("story"); err != nil {
				return fmt.Errorf("unable to decode story#8e358ca2: %w", err)
			}
		case "id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode story#8e358ca2: field id: %w", err)
			}
			s.ID = value
		case "sender_chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode story#8e358ca2: field sender_chat_id: %w", err)
			}
			s.SenderChatID = value
		case "date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode story#8e358ca2: field date: %w", err)
			}
			s.Date = value
		case "is_being_sent":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode story#8e358ca2: field is_being_sent: %w", err)
			}
			s.IsBeingSent = value
		case "is_being_edited":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode story#8e358ca2: field is_being_edited: %w", err)
			}
			s.IsBeingEdited = value
		case "is_edited":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode story#8e358ca2: field is_edited: %w", err)
			}
			s.IsEdited = value
		case "is_pinned":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode story#8e358ca2: field is_pinned: %w", err)
			}
			s.IsPinned = value
		case "is_visible_only_for_self":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode story#8e358ca2: field is_visible_only_for_self: %w", err)
			}
			s.IsVisibleOnlyForSelf = value
		case "can_be_deleted":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode story#8e358ca2: field can_be_deleted: %w", err)
			}
			s.CanBeDeleted = value
		case "can_be_edited":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode story#8e358ca2: field can_be_edited: %w", err)
			}
			s.CanBeEdited = value
		case "can_be_forwarded":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode story#8e358ca2: field can_be_forwarded: %w", err)
			}
			s.CanBeForwarded = value
		case "can_be_replied":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode story#8e358ca2: field can_be_replied: %w", err)
			}
			s.CanBeReplied = value
		case "can_toggle_is_pinned":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode story#8e358ca2: field can_toggle_is_pinned: %w", err)
			}
			s.CanToggleIsPinned = value
		case "can_get_statistics":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode story#8e358ca2: field can_get_statistics: %w", err)
			}
			s.CanGetStatistics = value
		case "can_get_interactions":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode story#8e358ca2: field can_get_interactions: %w", err)
			}
			s.CanGetInteractions = value
		case "has_expired_viewers":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode story#8e358ca2: field has_expired_viewers: %w", err)
			}
			s.HasExpiredViewers = value
		case "repost_info":
			if err := s.RepostInfo.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode story#8e358ca2: field repost_info: %w", err)
			}
		case "interaction_info":
			if err := s.InteractionInfo.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode story#8e358ca2: field interaction_info: %w", err)
			}
		case "chosen_reaction_type":
			value, err := DecodeTDLibJSONReactionType(b)
			if err != nil {
				return fmt.Errorf("unable to decode story#8e358ca2: field chosen_reaction_type: %w", err)
			}
			s.ChosenReactionType = value
		case "privacy_settings":
			value, err := DecodeTDLibJSONStoryPrivacySettings(b)
			if err != nil {
				return fmt.Errorf("unable to decode story#8e358ca2: field privacy_settings: %w", err)
			}
			s.PrivacySettings = value
		case "content":
			value, err := DecodeTDLibJSONStoryContent(b)
			if err != nil {
				return fmt.Errorf("unable to decode story#8e358ca2: field content: %w", err)
			}
			s.Content = value
		case "areas":
			if err := b.Arr(func(b tdjson.Decoder) error {
				var value StoryArea
				if err := value.DecodeTDLibJSON(b); err != nil {
					return fmt.Errorf("unable to decode story#8e358ca2: field areas: %w", err)
				}
				s.Areas = append(s.Areas, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode story#8e358ca2: field areas: %w", err)
			}
		case "caption":
			if err := s.Caption.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode story#8e358ca2: field caption: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetID returns value of ID field.
func (s *Story) GetID() (value int32) {
	if s == nil {
		return
	}
	return s.ID
}

// GetSenderChatID returns value of SenderChatID field.
func (s *Story) GetSenderChatID() (value int64) {
	if s == nil {
		return
	}
	return s.SenderChatID
}

// GetDate returns value of Date field.
func (s *Story) GetDate() (value int32) {
	if s == nil {
		return
	}
	return s.Date
}

// GetIsBeingSent returns value of IsBeingSent field.
func (s *Story) GetIsBeingSent() (value bool) {
	if s == nil {
		return
	}
	return s.IsBeingSent
}

// GetIsBeingEdited returns value of IsBeingEdited field.
func (s *Story) GetIsBeingEdited() (value bool) {
	if s == nil {
		return
	}
	return s.IsBeingEdited
}

// GetIsEdited returns value of IsEdited field.
func (s *Story) GetIsEdited() (value bool) {
	if s == nil {
		return
	}
	return s.IsEdited
}

// GetIsPinned returns value of IsPinned field.
func (s *Story) GetIsPinned() (value bool) {
	if s == nil {
		return
	}
	return s.IsPinned
}

// GetIsVisibleOnlyForSelf returns value of IsVisibleOnlyForSelf field.
func (s *Story) GetIsVisibleOnlyForSelf() (value bool) {
	if s == nil {
		return
	}
	return s.IsVisibleOnlyForSelf
}

// GetCanBeDeleted returns value of CanBeDeleted field.
func (s *Story) GetCanBeDeleted() (value bool) {
	if s == nil {
		return
	}
	return s.CanBeDeleted
}

// GetCanBeEdited returns value of CanBeEdited field.
func (s *Story) GetCanBeEdited() (value bool) {
	if s == nil {
		return
	}
	return s.CanBeEdited
}

// GetCanBeForwarded returns value of CanBeForwarded field.
func (s *Story) GetCanBeForwarded() (value bool) {
	if s == nil {
		return
	}
	return s.CanBeForwarded
}

// GetCanBeReplied returns value of CanBeReplied field.
func (s *Story) GetCanBeReplied() (value bool) {
	if s == nil {
		return
	}
	return s.CanBeReplied
}

// GetCanToggleIsPinned returns value of CanToggleIsPinned field.
func (s *Story) GetCanToggleIsPinned() (value bool) {
	if s == nil {
		return
	}
	return s.CanToggleIsPinned
}

// GetCanGetStatistics returns value of CanGetStatistics field.
func (s *Story) GetCanGetStatistics() (value bool) {
	if s == nil {
		return
	}
	return s.CanGetStatistics
}

// GetCanGetInteractions returns value of CanGetInteractions field.
func (s *Story) GetCanGetInteractions() (value bool) {
	if s == nil {
		return
	}
	return s.CanGetInteractions
}

// GetHasExpiredViewers returns value of HasExpiredViewers field.
func (s *Story) GetHasExpiredViewers() (value bool) {
	if s == nil {
		return
	}
	return s.HasExpiredViewers
}

// GetRepostInfo returns value of RepostInfo field.
func (s *Story) GetRepostInfo() (value StoryRepostInfo) {
	if s == nil {
		return
	}
	return s.RepostInfo
}

// GetInteractionInfo returns value of InteractionInfo field.
func (s *Story) GetInteractionInfo() (value StoryInteractionInfo) {
	if s == nil {
		return
	}
	return s.InteractionInfo
}

// GetChosenReactionType returns value of ChosenReactionType field.
func (s *Story) GetChosenReactionType() (value ReactionTypeClass) {
	if s == nil {
		return
	}
	return s.ChosenReactionType
}

// GetPrivacySettings returns value of PrivacySettings field.
func (s *Story) GetPrivacySettings() (value StoryPrivacySettingsClass) {
	if s == nil {
		return
	}
	return s.PrivacySettings
}

// GetContent returns value of Content field.
func (s *Story) GetContent() (value StoryContentClass) {
	if s == nil {
		return
	}
	return s.Content
}

// GetAreas returns value of Areas field.
func (s *Story) GetAreas() (value []StoryArea) {
	if s == nil {
		return
	}
	return s.Areas
}

// GetCaption returns value of Caption field.
func (s *Story) GetCaption() (value FormattedText) {
	if s == nil {
		return
	}
	return s.Caption
}
