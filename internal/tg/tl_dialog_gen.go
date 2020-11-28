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

// Dialog represents TL type `dialog#2c171f72`.
type Dialog struct {
	// Flags field of Dialog.
	Flags bin.Fields
	// Pinned field of Dialog.
	Pinned bool
	// UnreadMark field of Dialog.
	UnreadMark bool
	// Peer field of Dialog.
	Peer PeerClass
	// TopMessage field of Dialog.
	TopMessage int
	// ReadInboxMaxID field of Dialog.
	ReadInboxMaxID int
	// ReadOutboxMaxID field of Dialog.
	ReadOutboxMaxID int
	// UnreadCount field of Dialog.
	UnreadCount int
	// UnreadMentionsCount field of Dialog.
	UnreadMentionsCount int
	// NotifySettings field of Dialog.
	NotifySettings PeerNotifySettings
	// Pts field of Dialog.
	//
	// Use SetPts and GetPts helpers.
	Pts int
	// Draft field of Dialog.
	//
	// Use SetDraft and GetDraft helpers.
	Draft DraftMessageClass
	// FolderID field of Dialog.
	//
	// Use SetFolderID and GetFolderID helpers.
	FolderID int
}

// DialogTypeID is TL type id of Dialog.
const DialogTypeID = 0x2c171f72

// Encode implements bin.Encoder.
func (d *Dialog) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode dialog#2c171f72 as nil")
	}
	b.PutID(DialogTypeID)
	if err := d.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode dialog#2c171f72: field flags: %w", err)
	}
	if d.Peer == nil {
		return fmt.Errorf("unable to encode dialog#2c171f72: field peer is nil")
	}
	if err := d.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode dialog#2c171f72: field peer: %w", err)
	}
	b.PutInt(d.TopMessage)
	b.PutInt(d.ReadInboxMaxID)
	b.PutInt(d.ReadOutboxMaxID)
	b.PutInt(d.UnreadCount)
	b.PutInt(d.UnreadMentionsCount)
	if err := d.NotifySettings.Encode(b); err != nil {
		return fmt.Errorf("unable to encode dialog#2c171f72: field notify_settings: %w", err)
	}
	if d.Flags.Has(0) {
		b.PutInt(d.Pts)
	}
	if d.Flags.Has(1) {
		if d.Draft == nil {
			return fmt.Errorf("unable to encode dialog#2c171f72: field draft is nil")
		}
		if err := d.Draft.Encode(b); err != nil {
			return fmt.Errorf("unable to encode dialog#2c171f72: field draft: %w", err)
		}
	}
	if d.Flags.Has(4) {
		b.PutInt(d.FolderID)
	}
	return nil
}

// SetPinned sets value of Pinned conditional field.
func (d *Dialog) SetPinned(value bool) {
	if value {
		d.Flags.Set(2)
	} else {
		d.Flags.Unset(2)
	}
}

// SetUnreadMark sets value of UnreadMark conditional field.
func (d *Dialog) SetUnreadMark(value bool) {
	if value {
		d.Flags.Set(3)
	} else {
		d.Flags.Unset(3)
	}
}

// SetPts sets value of Pts conditional field.
func (d *Dialog) SetPts(value int) {
	d.Flags.Set(0)
	d.Pts = value
}

// GetPts returns value of Pts conditional field and
// boolean which is true if field was set.
func (d *Dialog) GetPts() (value int, ok bool) {
	if !d.Flags.Has(0) {
		return value, false
	}
	return d.Pts, true
}

// SetDraft sets value of Draft conditional field.
func (d *Dialog) SetDraft(value DraftMessageClass) {
	d.Flags.Set(1)
	d.Draft = value
}

// GetDraft returns value of Draft conditional field and
// boolean which is true if field was set.
func (d *Dialog) GetDraft() (value DraftMessageClass, ok bool) {
	if !d.Flags.Has(1) {
		return value, false
	}
	return d.Draft, true
}

// SetFolderID sets value of FolderID conditional field.
func (d *Dialog) SetFolderID(value int) {
	d.Flags.Set(4)
	d.FolderID = value
}

// GetFolderID returns value of FolderID conditional field and
// boolean which is true if field was set.
func (d *Dialog) GetFolderID() (value int, ok bool) {
	if !d.Flags.Has(4) {
		return value, false
	}
	return d.FolderID, true
}

// Decode implements bin.Decoder.
func (d *Dialog) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode dialog#2c171f72 to nil")
	}
	if err := b.ConsumeID(DialogTypeID); err != nil {
		return fmt.Errorf("unable to decode dialog#2c171f72: %w", err)
	}
	{
		if err := d.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode dialog#2c171f72: field flags: %w", err)
		}
	}
	d.Pinned = d.Flags.Has(2)
	d.UnreadMark = d.Flags.Has(3)
	{
		value, err := DecodePeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode dialog#2c171f72: field peer: %w", err)
		}
		d.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode dialog#2c171f72: field top_message: %w", err)
		}
		d.TopMessage = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode dialog#2c171f72: field read_inbox_max_id: %w", err)
		}
		d.ReadInboxMaxID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode dialog#2c171f72: field read_outbox_max_id: %w", err)
		}
		d.ReadOutboxMaxID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode dialog#2c171f72: field unread_count: %w", err)
		}
		d.UnreadCount = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode dialog#2c171f72: field unread_mentions_count: %w", err)
		}
		d.UnreadMentionsCount = value
	}
	{
		if err := d.NotifySettings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode dialog#2c171f72: field notify_settings: %w", err)
		}
	}
	if d.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode dialog#2c171f72: field pts: %w", err)
		}
		d.Pts = value
	}
	if d.Flags.Has(1) {
		value, err := DecodeDraftMessage(b)
		if err != nil {
			return fmt.Errorf("unable to decode dialog#2c171f72: field draft: %w", err)
		}
		d.Draft = value
	}
	if d.Flags.Has(4) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode dialog#2c171f72: field folder_id: %w", err)
		}
		d.FolderID = value
	}
	return nil
}

// construct implements constructor of DialogClass.
func (d Dialog) construct() DialogClass { return &d }

// Ensuring interfaces in compile-time for Dialog.
var (
	_ bin.Encoder = &Dialog{}
	_ bin.Decoder = &Dialog{}

	_ DialogClass = &Dialog{}
)

// DialogFolder represents TL type `dialogFolder#71bd134c`.
type DialogFolder struct {
	// Flags field of DialogFolder.
	Flags bin.Fields
	// Pinned field of DialogFolder.
	Pinned bool
	// Folder field of DialogFolder.
	Folder Folder
	// Peer field of DialogFolder.
	Peer PeerClass
	// TopMessage field of DialogFolder.
	TopMessage int
	// UnreadMutedPeersCount field of DialogFolder.
	UnreadMutedPeersCount int
	// UnreadUnmutedPeersCount field of DialogFolder.
	UnreadUnmutedPeersCount int
	// UnreadMutedMessagesCount field of DialogFolder.
	UnreadMutedMessagesCount int
	// UnreadUnmutedMessagesCount field of DialogFolder.
	UnreadUnmutedMessagesCount int
}

// DialogFolderTypeID is TL type id of DialogFolder.
const DialogFolderTypeID = 0x71bd134c

// Encode implements bin.Encoder.
func (d *DialogFolder) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode dialogFolder#71bd134c as nil")
	}
	b.PutID(DialogFolderTypeID)
	if err := d.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode dialogFolder#71bd134c: field flags: %w", err)
	}
	if err := d.Folder.Encode(b); err != nil {
		return fmt.Errorf("unable to encode dialogFolder#71bd134c: field folder: %w", err)
	}
	if d.Peer == nil {
		return fmt.Errorf("unable to encode dialogFolder#71bd134c: field peer is nil")
	}
	if err := d.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode dialogFolder#71bd134c: field peer: %w", err)
	}
	b.PutInt(d.TopMessage)
	b.PutInt(d.UnreadMutedPeersCount)
	b.PutInt(d.UnreadUnmutedPeersCount)
	b.PutInt(d.UnreadMutedMessagesCount)
	b.PutInt(d.UnreadUnmutedMessagesCount)
	return nil
}

// SetPinned sets value of Pinned conditional field.
func (d *DialogFolder) SetPinned(value bool) {
	if value {
		d.Flags.Set(2)
	} else {
		d.Flags.Unset(2)
	}
}

// Decode implements bin.Decoder.
func (d *DialogFolder) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode dialogFolder#71bd134c to nil")
	}
	if err := b.ConsumeID(DialogFolderTypeID); err != nil {
		return fmt.Errorf("unable to decode dialogFolder#71bd134c: %w", err)
	}
	{
		if err := d.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode dialogFolder#71bd134c: field flags: %w", err)
		}
	}
	d.Pinned = d.Flags.Has(2)
	{
		if err := d.Folder.Decode(b); err != nil {
			return fmt.Errorf("unable to decode dialogFolder#71bd134c: field folder: %w", err)
		}
	}
	{
		value, err := DecodePeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode dialogFolder#71bd134c: field peer: %w", err)
		}
		d.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode dialogFolder#71bd134c: field top_message: %w", err)
		}
		d.TopMessage = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode dialogFolder#71bd134c: field unread_muted_peers_count: %w", err)
		}
		d.UnreadMutedPeersCount = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode dialogFolder#71bd134c: field unread_unmuted_peers_count: %w", err)
		}
		d.UnreadUnmutedPeersCount = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode dialogFolder#71bd134c: field unread_muted_messages_count: %w", err)
		}
		d.UnreadMutedMessagesCount = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode dialogFolder#71bd134c: field unread_unmuted_messages_count: %w", err)
		}
		d.UnreadUnmutedMessagesCount = value
	}
	return nil
}

// construct implements constructor of DialogClass.
func (d DialogFolder) construct() DialogClass { return &d }

// Ensuring interfaces in compile-time for DialogFolder.
var (
	_ bin.Encoder = &DialogFolder{}
	_ bin.Decoder = &DialogFolder{}

	_ DialogClass = &DialogFolder{}
)

// DialogClass represents Dialog generic type.
//
// Example:
//  g, err := DecodeDialog(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *Dialog: // dialog#2c171f72
//  case *DialogFolder: // dialogFolder#71bd134c
//  default: panic(v)
//  }
type DialogClass interface {
	bin.Encoder
	bin.Decoder
	construct() DialogClass
}

// DecodeDialog implements binary de-serialization for DialogClass.
func DecodeDialog(buf *bin.Buffer) (DialogClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case DialogTypeID:
		// Decoding dialog#2c171f72.
		v := Dialog{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode DialogClass: %w", err)
		}
		return &v, nil
	case DialogFolderTypeID:
		// Decoding dialogFolder#71bd134c.
		v := DialogFolder{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode DialogClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode DialogClass: %w", bin.NewUnexpectedID(id))
	}
}
