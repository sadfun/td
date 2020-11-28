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

// StatsGroupTopAdmin represents TL type `statsGroupTopAdmin#6014f412`.
type StatsGroupTopAdmin struct {
	// UserID field of StatsGroupTopAdmin.
	UserID int
	// Deleted field of StatsGroupTopAdmin.
	Deleted int
	// Kicked field of StatsGroupTopAdmin.
	Kicked int
	// Banned field of StatsGroupTopAdmin.
	Banned int
}

// StatsGroupTopAdminTypeID is TL type id of StatsGroupTopAdmin.
const StatsGroupTopAdminTypeID = 0x6014f412

// Encode implements bin.Encoder.
func (s *StatsGroupTopAdmin) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode statsGroupTopAdmin#6014f412 as nil")
	}
	b.PutID(StatsGroupTopAdminTypeID)
	b.PutInt(s.UserID)
	b.PutInt(s.Deleted)
	b.PutInt(s.Kicked)
	b.PutInt(s.Banned)
	return nil
}

// Decode implements bin.Decoder.
func (s *StatsGroupTopAdmin) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode statsGroupTopAdmin#6014f412 to nil")
	}
	if err := b.ConsumeID(StatsGroupTopAdminTypeID); err != nil {
		return fmt.Errorf("unable to decode statsGroupTopAdmin#6014f412: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode statsGroupTopAdmin#6014f412: field user_id: %w", err)
		}
		s.UserID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode statsGroupTopAdmin#6014f412: field deleted: %w", err)
		}
		s.Deleted = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode statsGroupTopAdmin#6014f412: field kicked: %w", err)
		}
		s.Kicked = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode statsGroupTopAdmin#6014f412: field banned: %w", err)
		}
		s.Banned = value
	}
	return nil
}

// Ensuring interfaces in compile-time for StatsGroupTopAdmin.
var (
	_ bin.Encoder = &StatsGroupTopAdmin{}
	_ bin.Decoder = &StatsGroupTopAdmin{}
)
