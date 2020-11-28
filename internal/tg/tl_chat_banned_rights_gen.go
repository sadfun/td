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

// ChatBannedRights represents TL type `chatBannedRights#9f120418`.
type ChatBannedRights struct {
	// Flags field of ChatBannedRights.
	Flags bin.Fields
	// ViewMessages field of ChatBannedRights.
	ViewMessages bool
	// SendMessages field of ChatBannedRights.
	SendMessages bool
	// SendMedia field of ChatBannedRights.
	SendMedia bool
	// SendStickers field of ChatBannedRights.
	SendStickers bool
	// SendGifs field of ChatBannedRights.
	SendGifs bool
	// SendGames field of ChatBannedRights.
	SendGames bool
	// SendInline field of ChatBannedRights.
	SendInline bool
	// EmbedLinks field of ChatBannedRights.
	EmbedLinks bool
	// SendPolls field of ChatBannedRights.
	SendPolls bool
	// ChangeInfo field of ChatBannedRights.
	ChangeInfo bool
	// InviteUsers field of ChatBannedRights.
	InviteUsers bool
	// PinMessages field of ChatBannedRights.
	PinMessages bool
	// UntilDate field of ChatBannedRights.
	UntilDate int
}

// ChatBannedRightsTypeID is TL type id of ChatBannedRights.
const ChatBannedRightsTypeID = 0x9f120418

// Encode implements bin.Encoder.
func (c *ChatBannedRights) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatBannedRights#9f120418 as nil")
	}
	b.PutID(ChatBannedRightsTypeID)
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chatBannedRights#9f120418: field flags: %w", err)
	}
	b.PutInt(c.UntilDate)
	return nil
}

// SetViewMessages sets value of ViewMessages conditional field.
func (c *ChatBannedRights) SetViewMessages(value bool) {
	if value {
		c.Flags.Set(0)
	} else {
		c.Flags.Unset(0)
	}
}

// SetSendMessages sets value of SendMessages conditional field.
func (c *ChatBannedRights) SetSendMessages(value bool) {
	if value {
		c.Flags.Set(1)
	} else {
		c.Flags.Unset(1)
	}
}

// SetSendMedia sets value of SendMedia conditional field.
func (c *ChatBannedRights) SetSendMedia(value bool) {
	if value {
		c.Flags.Set(2)
	} else {
		c.Flags.Unset(2)
	}
}

// SetSendStickers sets value of SendStickers conditional field.
func (c *ChatBannedRights) SetSendStickers(value bool) {
	if value {
		c.Flags.Set(3)
	} else {
		c.Flags.Unset(3)
	}
}

// SetSendGifs sets value of SendGifs conditional field.
func (c *ChatBannedRights) SetSendGifs(value bool) {
	if value {
		c.Flags.Set(4)
	} else {
		c.Flags.Unset(4)
	}
}

// SetSendGames sets value of SendGames conditional field.
func (c *ChatBannedRights) SetSendGames(value bool) {
	if value {
		c.Flags.Set(5)
	} else {
		c.Flags.Unset(5)
	}
}

// SetSendInline sets value of SendInline conditional field.
func (c *ChatBannedRights) SetSendInline(value bool) {
	if value {
		c.Flags.Set(6)
	} else {
		c.Flags.Unset(6)
	}
}

// SetEmbedLinks sets value of EmbedLinks conditional field.
func (c *ChatBannedRights) SetEmbedLinks(value bool) {
	if value {
		c.Flags.Set(7)
	} else {
		c.Flags.Unset(7)
	}
}

// SetSendPolls sets value of SendPolls conditional field.
func (c *ChatBannedRights) SetSendPolls(value bool) {
	if value {
		c.Flags.Set(8)
	} else {
		c.Flags.Unset(8)
	}
}

// SetChangeInfo sets value of ChangeInfo conditional field.
func (c *ChatBannedRights) SetChangeInfo(value bool) {
	if value {
		c.Flags.Set(10)
	} else {
		c.Flags.Unset(10)
	}
}

// SetInviteUsers sets value of InviteUsers conditional field.
func (c *ChatBannedRights) SetInviteUsers(value bool) {
	if value {
		c.Flags.Set(15)
	} else {
		c.Flags.Unset(15)
	}
}

// SetPinMessages sets value of PinMessages conditional field.
func (c *ChatBannedRights) SetPinMessages(value bool) {
	if value {
		c.Flags.Set(17)
	} else {
		c.Flags.Unset(17)
	}
}

// Decode implements bin.Decoder.
func (c *ChatBannedRights) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatBannedRights#9f120418 to nil")
	}
	if err := b.ConsumeID(ChatBannedRightsTypeID); err != nil {
		return fmt.Errorf("unable to decode chatBannedRights#9f120418: %w", err)
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode chatBannedRights#9f120418: field flags: %w", err)
		}
	}
	c.ViewMessages = c.Flags.Has(0)
	c.SendMessages = c.Flags.Has(1)
	c.SendMedia = c.Flags.Has(2)
	c.SendStickers = c.Flags.Has(3)
	c.SendGifs = c.Flags.Has(4)
	c.SendGames = c.Flags.Has(5)
	c.SendInline = c.Flags.Has(6)
	c.EmbedLinks = c.Flags.Has(7)
	c.SendPolls = c.Flags.Has(8)
	c.ChangeInfo = c.Flags.Has(10)
	c.InviteUsers = c.Flags.Has(15)
	c.PinMessages = c.Flags.Has(17)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatBannedRights#9f120418: field until_date: %w", err)
		}
		c.UntilDate = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ChatBannedRights.
var (
	_ bin.Encoder = &ChatBannedRights{}
	_ bin.Decoder = &ChatBannedRights{}
)
