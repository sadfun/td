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

// EncryptedChatEmpty represents TL type `encryptedChatEmpty#ab7ec0a0`.
type EncryptedChatEmpty struct {
	// ID field of EncryptedChatEmpty.
	ID int
}

// EncryptedChatEmptyTypeID is TL type id of EncryptedChatEmpty.
const EncryptedChatEmptyTypeID = 0xab7ec0a0

// Encode implements bin.Encoder.
func (e *EncryptedChatEmpty) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode encryptedChatEmpty#ab7ec0a0 as nil")
	}
	b.PutID(EncryptedChatEmptyTypeID)
	b.PutInt(e.ID)
	return nil
}

// Decode implements bin.Decoder.
func (e *EncryptedChatEmpty) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode encryptedChatEmpty#ab7ec0a0 to nil")
	}
	if err := b.ConsumeID(EncryptedChatEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode encryptedChatEmpty#ab7ec0a0: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChatEmpty#ab7ec0a0: field id: %w", err)
		}
		e.ID = value
	}
	return nil
}

// construct implements constructor of EncryptedChatClass.
func (e EncryptedChatEmpty) construct() EncryptedChatClass { return &e }

// Ensuring interfaces in compile-time for EncryptedChatEmpty.
var (
	_ bin.Encoder = &EncryptedChatEmpty{}
	_ bin.Decoder = &EncryptedChatEmpty{}

	_ EncryptedChatClass = &EncryptedChatEmpty{}
)

// EncryptedChatWaiting represents TL type `encryptedChatWaiting#3bf703dc`.
type EncryptedChatWaiting struct {
	// ID field of EncryptedChatWaiting.
	ID int
	// AccessHash field of EncryptedChatWaiting.
	AccessHash int64
	// Date field of EncryptedChatWaiting.
	Date int
	// AdminID field of EncryptedChatWaiting.
	AdminID int
	// ParticipantID field of EncryptedChatWaiting.
	ParticipantID int
}

// EncryptedChatWaitingTypeID is TL type id of EncryptedChatWaiting.
const EncryptedChatWaitingTypeID = 0x3bf703dc

// Encode implements bin.Encoder.
func (e *EncryptedChatWaiting) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode encryptedChatWaiting#3bf703dc as nil")
	}
	b.PutID(EncryptedChatWaitingTypeID)
	b.PutInt(e.ID)
	b.PutLong(e.AccessHash)
	b.PutInt(e.Date)
	b.PutInt(e.AdminID)
	b.PutInt(e.ParticipantID)
	return nil
}

// Decode implements bin.Decoder.
func (e *EncryptedChatWaiting) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode encryptedChatWaiting#3bf703dc to nil")
	}
	if err := b.ConsumeID(EncryptedChatWaitingTypeID); err != nil {
		return fmt.Errorf("unable to decode encryptedChatWaiting#3bf703dc: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChatWaiting#3bf703dc: field id: %w", err)
		}
		e.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChatWaiting#3bf703dc: field access_hash: %w", err)
		}
		e.AccessHash = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChatWaiting#3bf703dc: field date: %w", err)
		}
		e.Date = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChatWaiting#3bf703dc: field admin_id: %w", err)
		}
		e.AdminID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChatWaiting#3bf703dc: field participant_id: %w", err)
		}
		e.ParticipantID = value
	}
	return nil
}

// construct implements constructor of EncryptedChatClass.
func (e EncryptedChatWaiting) construct() EncryptedChatClass { return &e }

// Ensuring interfaces in compile-time for EncryptedChatWaiting.
var (
	_ bin.Encoder = &EncryptedChatWaiting{}
	_ bin.Decoder = &EncryptedChatWaiting{}

	_ EncryptedChatClass = &EncryptedChatWaiting{}
)

// EncryptedChatRequested represents TL type `encryptedChatRequested#62718a82`.
type EncryptedChatRequested struct {
	// Flags field of EncryptedChatRequested.
	Flags bin.Fields
	// FolderID field of EncryptedChatRequested.
	//
	// Use SetFolderID and GetFolderID helpers.
	FolderID int
	// ID field of EncryptedChatRequested.
	ID int
	// AccessHash field of EncryptedChatRequested.
	AccessHash int64
	// Date field of EncryptedChatRequested.
	Date int
	// AdminID field of EncryptedChatRequested.
	AdminID int
	// ParticipantID field of EncryptedChatRequested.
	ParticipantID int
	// GA field of EncryptedChatRequested.
	GA []byte
}

// EncryptedChatRequestedTypeID is TL type id of EncryptedChatRequested.
const EncryptedChatRequestedTypeID = 0x62718a82

// Encode implements bin.Encoder.
func (e *EncryptedChatRequested) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode encryptedChatRequested#62718a82 as nil")
	}
	b.PutID(EncryptedChatRequestedTypeID)
	if err := e.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode encryptedChatRequested#62718a82: field flags: %w", err)
	}
	if e.Flags.Has(0) {
		b.PutInt(e.FolderID)
	}
	b.PutInt(e.ID)
	b.PutLong(e.AccessHash)
	b.PutInt(e.Date)
	b.PutInt(e.AdminID)
	b.PutInt(e.ParticipantID)
	b.PutBytes(e.GA)
	return nil
}

// SetFolderID sets value of FolderID conditional field.
func (e *EncryptedChatRequested) SetFolderID(value int) {
	e.Flags.Set(0)
	e.FolderID = value
}

// GetFolderID returns value of FolderID conditional field and
// boolean which is true if field was set.
func (e *EncryptedChatRequested) GetFolderID() (value int, ok bool) {
	if !e.Flags.Has(0) {
		return value, false
	}
	return e.FolderID, true
}

// Decode implements bin.Decoder.
func (e *EncryptedChatRequested) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode encryptedChatRequested#62718a82 to nil")
	}
	if err := b.ConsumeID(EncryptedChatRequestedTypeID); err != nil {
		return fmt.Errorf("unable to decode encryptedChatRequested#62718a82: %w", err)
	}
	{
		if err := e.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode encryptedChatRequested#62718a82: field flags: %w", err)
		}
	}
	if e.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChatRequested#62718a82: field folder_id: %w", err)
		}
		e.FolderID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChatRequested#62718a82: field id: %w", err)
		}
		e.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChatRequested#62718a82: field access_hash: %w", err)
		}
		e.AccessHash = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChatRequested#62718a82: field date: %w", err)
		}
		e.Date = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChatRequested#62718a82: field admin_id: %w", err)
		}
		e.AdminID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChatRequested#62718a82: field participant_id: %w", err)
		}
		e.ParticipantID = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChatRequested#62718a82: field g_a: %w", err)
		}
		e.GA = value
	}
	return nil
}

// construct implements constructor of EncryptedChatClass.
func (e EncryptedChatRequested) construct() EncryptedChatClass { return &e }

// Ensuring interfaces in compile-time for EncryptedChatRequested.
var (
	_ bin.Encoder = &EncryptedChatRequested{}
	_ bin.Decoder = &EncryptedChatRequested{}

	_ EncryptedChatClass = &EncryptedChatRequested{}
)

// EncryptedChat represents TL type `encryptedChat#fa56ce36`.
type EncryptedChat struct {
	// ID field of EncryptedChat.
	ID int
	// AccessHash field of EncryptedChat.
	AccessHash int64
	// Date field of EncryptedChat.
	Date int
	// AdminID field of EncryptedChat.
	AdminID int
	// ParticipantID field of EncryptedChat.
	ParticipantID int
	// GAOrB field of EncryptedChat.
	GAOrB []byte
	// KeyFingerprint field of EncryptedChat.
	KeyFingerprint int64
}

// EncryptedChatTypeID is TL type id of EncryptedChat.
const EncryptedChatTypeID = 0xfa56ce36

// Encode implements bin.Encoder.
func (e *EncryptedChat) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode encryptedChat#fa56ce36 as nil")
	}
	b.PutID(EncryptedChatTypeID)
	b.PutInt(e.ID)
	b.PutLong(e.AccessHash)
	b.PutInt(e.Date)
	b.PutInt(e.AdminID)
	b.PutInt(e.ParticipantID)
	b.PutBytes(e.GAOrB)
	b.PutLong(e.KeyFingerprint)
	return nil
}

// Decode implements bin.Decoder.
func (e *EncryptedChat) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode encryptedChat#fa56ce36 to nil")
	}
	if err := b.ConsumeID(EncryptedChatTypeID); err != nil {
		return fmt.Errorf("unable to decode encryptedChat#fa56ce36: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChat#fa56ce36: field id: %w", err)
		}
		e.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChat#fa56ce36: field access_hash: %w", err)
		}
		e.AccessHash = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChat#fa56ce36: field date: %w", err)
		}
		e.Date = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChat#fa56ce36: field admin_id: %w", err)
		}
		e.AdminID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChat#fa56ce36: field participant_id: %w", err)
		}
		e.ParticipantID = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChat#fa56ce36: field g_a_or_b: %w", err)
		}
		e.GAOrB = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChat#fa56ce36: field key_fingerprint: %w", err)
		}
		e.KeyFingerprint = value
	}
	return nil
}

// construct implements constructor of EncryptedChatClass.
func (e EncryptedChat) construct() EncryptedChatClass { return &e }

// Ensuring interfaces in compile-time for EncryptedChat.
var (
	_ bin.Encoder = &EncryptedChat{}
	_ bin.Decoder = &EncryptedChat{}

	_ EncryptedChatClass = &EncryptedChat{}
)

// EncryptedChatDiscarded represents TL type `encryptedChatDiscarded#13d6dd27`.
type EncryptedChatDiscarded struct {
	// ID field of EncryptedChatDiscarded.
	ID int
}

// EncryptedChatDiscardedTypeID is TL type id of EncryptedChatDiscarded.
const EncryptedChatDiscardedTypeID = 0x13d6dd27

// Encode implements bin.Encoder.
func (e *EncryptedChatDiscarded) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode encryptedChatDiscarded#13d6dd27 as nil")
	}
	b.PutID(EncryptedChatDiscardedTypeID)
	b.PutInt(e.ID)
	return nil
}

// Decode implements bin.Decoder.
func (e *EncryptedChatDiscarded) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode encryptedChatDiscarded#13d6dd27 to nil")
	}
	if err := b.ConsumeID(EncryptedChatDiscardedTypeID); err != nil {
		return fmt.Errorf("unable to decode encryptedChatDiscarded#13d6dd27: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChatDiscarded#13d6dd27: field id: %w", err)
		}
		e.ID = value
	}
	return nil
}

// construct implements constructor of EncryptedChatClass.
func (e EncryptedChatDiscarded) construct() EncryptedChatClass { return &e }

// Ensuring interfaces in compile-time for EncryptedChatDiscarded.
var (
	_ bin.Encoder = &EncryptedChatDiscarded{}
	_ bin.Decoder = &EncryptedChatDiscarded{}

	_ EncryptedChatClass = &EncryptedChatDiscarded{}
)

// EncryptedChatClass represents EncryptedChat generic type.
//
// Example:
//  g, err := DecodeEncryptedChat(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *EncryptedChatEmpty: // encryptedChatEmpty#ab7ec0a0
//  case *EncryptedChatWaiting: // encryptedChatWaiting#3bf703dc
//  case *EncryptedChatRequested: // encryptedChatRequested#62718a82
//  case *EncryptedChat: // encryptedChat#fa56ce36
//  case *EncryptedChatDiscarded: // encryptedChatDiscarded#13d6dd27
//  default: panic(v)
//  }
type EncryptedChatClass interface {
	bin.Encoder
	bin.Decoder
	construct() EncryptedChatClass
}

// DecodeEncryptedChat implements binary de-serialization for EncryptedChatClass.
func DecodeEncryptedChat(buf *bin.Buffer) (EncryptedChatClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case EncryptedChatEmptyTypeID:
		// Decoding encryptedChatEmpty#ab7ec0a0.
		v := EncryptedChatEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode EncryptedChatClass: %w", err)
		}
		return &v, nil
	case EncryptedChatWaitingTypeID:
		// Decoding encryptedChatWaiting#3bf703dc.
		v := EncryptedChatWaiting{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode EncryptedChatClass: %w", err)
		}
		return &v, nil
	case EncryptedChatRequestedTypeID:
		// Decoding encryptedChatRequested#62718a82.
		v := EncryptedChatRequested{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode EncryptedChatClass: %w", err)
		}
		return &v, nil
	case EncryptedChatTypeID:
		// Decoding encryptedChat#fa56ce36.
		v := EncryptedChat{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode EncryptedChatClass: %w", err)
		}
		return &v, nil
	case EncryptedChatDiscardedTypeID:
		// Decoding encryptedChatDiscarded#13d6dd27.
		v := EncryptedChatDiscarded{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode EncryptedChatClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode EncryptedChatClass: %w", bin.NewUnexpectedID(id))
	}
}
