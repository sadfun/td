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

// LangPackLanguage represents TL type `langPackLanguage#eeca5ce3`.
type LangPackLanguage struct {
	// Flags field of LangPackLanguage.
	Flags bin.Fields
	// Official field of LangPackLanguage.
	Official bool
	// Rtl field of LangPackLanguage.
	Rtl bool
	// Beta field of LangPackLanguage.
	Beta bool
	// Name field of LangPackLanguage.
	Name string
	// NativeName field of LangPackLanguage.
	NativeName string
	// LangCode field of LangPackLanguage.
	LangCode string
	// BaseLangCode field of LangPackLanguage.
	//
	// Use SetBaseLangCode and GetBaseLangCode helpers.
	BaseLangCode string
	// PluralCode field of LangPackLanguage.
	PluralCode string
	// StringsCount field of LangPackLanguage.
	StringsCount int
	// TranslatedCount field of LangPackLanguage.
	TranslatedCount int
	// TranslationsURL field of LangPackLanguage.
	TranslationsURL string
}

// LangPackLanguageTypeID is TL type id of LangPackLanguage.
const LangPackLanguageTypeID = 0xeeca5ce3

// Encode implements bin.Encoder.
func (l *LangPackLanguage) Encode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode langPackLanguage#eeca5ce3 as nil")
	}
	b.PutID(LangPackLanguageTypeID)
	if err := l.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode langPackLanguage#eeca5ce3: field flags: %w", err)
	}
	b.PutString(l.Name)
	b.PutString(l.NativeName)
	b.PutString(l.LangCode)
	if l.Flags.Has(1) {
		b.PutString(l.BaseLangCode)
	}
	b.PutString(l.PluralCode)
	b.PutInt(l.StringsCount)
	b.PutInt(l.TranslatedCount)
	b.PutString(l.TranslationsURL)
	return nil
}

// SetOfficial sets value of Official conditional field.
func (l *LangPackLanguage) SetOfficial(value bool) {
	if value {
		l.Flags.Set(0)
	} else {
		l.Flags.Unset(0)
	}
}

// SetRtl sets value of Rtl conditional field.
func (l *LangPackLanguage) SetRtl(value bool) {
	if value {
		l.Flags.Set(2)
	} else {
		l.Flags.Unset(2)
	}
}

// SetBeta sets value of Beta conditional field.
func (l *LangPackLanguage) SetBeta(value bool) {
	if value {
		l.Flags.Set(3)
	} else {
		l.Flags.Unset(3)
	}
}

// SetBaseLangCode sets value of BaseLangCode conditional field.
func (l *LangPackLanguage) SetBaseLangCode(value string) {
	l.Flags.Set(1)
	l.BaseLangCode = value
}

// GetBaseLangCode returns value of BaseLangCode conditional field and
// boolean which is true if field was set.
func (l *LangPackLanguage) GetBaseLangCode() (value string, ok bool) {
	if !l.Flags.Has(1) {
		return value, false
	}
	return l.BaseLangCode, true
}

// Decode implements bin.Decoder.
func (l *LangPackLanguage) Decode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode langPackLanguage#eeca5ce3 to nil")
	}
	if err := b.ConsumeID(LangPackLanguageTypeID); err != nil {
		return fmt.Errorf("unable to decode langPackLanguage#eeca5ce3: %w", err)
	}
	{
		if err := l.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode langPackLanguage#eeca5ce3: field flags: %w", err)
		}
	}
	l.Official = l.Flags.Has(0)
	l.Rtl = l.Flags.Has(2)
	l.Beta = l.Flags.Has(3)
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode langPackLanguage#eeca5ce3: field name: %w", err)
		}
		l.Name = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode langPackLanguage#eeca5ce3: field native_name: %w", err)
		}
		l.NativeName = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode langPackLanguage#eeca5ce3: field lang_code: %w", err)
		}
		l.LangCode = value
	}
	if l.Flags.Has(1) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode langPackLanguage#eeca5ce3: field base_lang_code: %w", err)
		}
		l.BaseLangCode = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode langPackLanguage#eeca5ce3: field plural_code: %w", err)
		}
		l.PluralCode = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode langPackLanguage#eeca5ce3: field strings_count: %w", err)
		}
		l.StringsCount = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode langPackLanguage#eeca5ce3: field translated_count: %w", err)
		}
		l.TranslatedCount = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode langPackLanguage#eeca5ce3: field translations_url: %w", err)
		}
		l.TranslationsURL = value
	}
	return nil
}

// Ensuring interfaces in compile-time for LangPackLanguage.
var (
	_ bin.Encoder = &LangPackLanguage{}
	_ bin.Decoder = &LangPackLanguage{}
)
