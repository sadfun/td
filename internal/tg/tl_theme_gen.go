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

// Theme represents TL type `theme#28f1114`.
type Theme struct {
	// Flags field of Theme.
	Flags bin.Fields
	// Creator field of Theme.
	Creator bool
	// Default field of Theme.
	Default bool
	// ID field of Theme.
	ID int64
	// AccessHash field of Theme.
	AccessHash int64
	// Slug field of Theme.
	Slug string
	// Title field of Theme.
	Title string
	// Document field of Theme.
	//
	// Use SetDocument and GetDocument helpers.
	Document DocumentClass
	// Settings field of Theme.
	//
	// Use SetSettings and GetSettings helpers.
	Settings ThemeSettings
	// InstallsCount field of Theme.
	InstallsCount int
}

// ThemeTypeID is TL type id of Theme.
const ThemeTypeID = 0x28f1114

// Encode implements bin.Encoder.
func (t *Theme) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode theme#28f1114 as nil")
	}
	b.PutID(ThemeTypeID)
	if err := t.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode theme#28f1114: field flags: %w", err)
	}
	b.PutLong(t.ID)
	b.PutLong(t.AccessHash)
	b.PutString(t.Slug)
	b.PutString(t.Title)
	if t.Flags.Has(2) {
		if t.Document == nil {
			return fmt.Errorf("unable to encode theme#28f1114: field document is nil")
		}
		if err := t.Document.Encode(b); err != nil {
			return fmt.Errorf("unable to encode theme#28f1114: field document: %w", err)
		}
	}
	if t.Flags.Has(3) {
		if err := t.Settings.Encode(b); err != nil {
			return fmt.Errorf("unable to encode theme#28f1114: field settings: %w", err)
		}
	}
	b.PutInt(t.InstallsCount)
	return nil
}

// SetCreator sets value of Creator conditional field.
func (t *Theme) SetCreator(value bool) {
	if value {
		t.Flags.Set(0)
	} else {
		t.Flags.Unset(0)
	}
}

// SetDefault sets value of Default conditional field.
func (t *Theme) SetDefault(value bool) {
	if value {
		t.Flags.Set(1)
	} else {
		t.Flags.Unset(1)
	}
}

// SetDocument sets value of Document conditional field.
func (t *Theme) SetDocument(value DocumentClass) {
	t.Flags.Set(2)
	t.Document = value
}

// GetDocument returns value of Document conditional field and
// boolean which is true if field was set.
func (t *Theme) GetDocument() (value DocumentClass, ok bool) {
	if !t.Flags.Has(2) {
		return value, false
	}
	return t.Document, true
}

// SetSettings sets value of Settings conditional field.
func (t *Theme) SetSettings(value ThemeSettings) {
	t.Flags.Set(3)
	t.Settings = value
}

// GetSettings returns value of Settings conditional field and
// boolean which is true if field was set.
func (t *Theme) GetSettings() (value ThemeSettings, ok bool) {
	if !t.Flags.Has(3) {
		return value, false
	}
	return t.Settings, true
}

// Decode implements bin.Decoder.
func (t *Theme) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode theme#28f1114 to nil")
	}
	if err := b.ConsumeID(ThemeTypeID); err != nil {
		return fmt.Errorf("unable to decode theme#28f1114: %w", err)
	}
	{
		if err := t.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode theme#28f1114: field flags: %w", err)
		}
	}
	t.Creator = t.Flags.Has(0)
	t.Default = t.Flags.Has(1)
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode theme#28f1114: field id: %w", err)
		}
		t.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode theme#28f1114: field access_hash: %w", err)
		}
		t.AccessHash = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode theme#28f1114: field slug: %w", err)
		}
		t.Slug = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode theme#28f1114: field title: %w", err)
		}
		t.Title = value
	}
	if t.Flags.Has(2) {
		value, err := DecodeDocument(b)
		if err != nil {
			return fmt.Errorf("unable to decode theme#28f1114: field document: %w", err)
		}
		t.Document = value
	}
	if t.Flags.Has(3) {
		if err := t.Settings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode theme#28f1114: field settings: %w", err)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode theme#28f1114: field installs_count: %w", err)
		}
		t.InstallsCount = value
	}
	return nil
}

// Ensuring interfaces in compile-time for Theme.
var (
	_ bin.Encoder = &Theme{}
	_ bin.Decoder = &Theme{}
)
