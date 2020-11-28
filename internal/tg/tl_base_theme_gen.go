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

// BaseThemeClassic represents TL type `baseThemeClassic#c3a12462`.
type BaseThemeClassic struct {
}

// BaseThemeClassicTypeID is TL type id of BaseThemeClassic.
const BaseThemeClassicTypeID = 0xc3a12462

// Encode implements bin.Encoder.
func (b *BaseThemeClassic) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode baseThemeClassic#c3a12462 as nil")
	}
	buf.PutID(BaseThemeClassicTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (b *BaseThemeClassic) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode baseThemeClassic#c3a12462 to nil")
	}
	if err := buf.ConsumeID(BaseThemeClassicTypeID); err != nil {
		return fmt.Errorf("unable to decode baseThemeClassic#c3a12462: %w", err)
	}
	return nil
}

// construct implements constructor of BaseThemeClass.
func (b BaseThemeClassic) construct() BaseThemeClass { return &b }

// Ensuring interfaces in compile-time for BaseThemeClassic.
var (
	_ bin.Encoder = &BaseThemeClassic{}
	_ bin.Decoder = &BaseThemeClassic{}

	_ BaseThemeClass = &BaseThemeClassic{}
)

// BaseThemeDay represents TL type `baseThemeDay#fbd81688`.
type BaseThemeDay struct {
}

// BaseThemeDayTypeID is TL type id of BaseThemeDay.
const BaseThemeDayTypeID = 0xfbd81688

// Encode implements bin.Encoder.
func (b *BaseThemeDay) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode baseThemeDay#fbd81688 as nil")
	}
	buf.PutID(BaseThemeDayTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (b *BaseThemeDay) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode baseThemeDay#fbd81688 to nil")
	}
	if err := buf.ConsumeID(BaseThemeDayTypeID); err != nil {
		return fmt.Errorf("unable to decode baseThemeDay#fbd81688: %w", err)
	}
	return nil
}

// construct implements constructor of BaseThemeClass.
func (b BaseThemeDay) construct() BaseThemeClass { return &b }

// Ensuring interfaces in compile-time for BaseThemeDay.
var (
	_ bin.Encoder = &BaseThemeDay{}
	_ bin.Decoder = &BaseThemeDay{}

	_ BaseThemeClass = &BaseThemeDay{}
)

// BaseThemeNight represents TL type `baseThemeNight#b7b31ea8`.
type BaseThemeNight struct {
}

// BaseThemeNightTypeID is TL type id of BaseThemeNight.
const BaseThemeNightTypeID = 0xb7b31ea8

// Encode implements bin.Encoder.
func (b *BaseThemeNight) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode baseThemeNight#b7b31ea8 as nil")
	}
	buf.PutID(BaseThemeNightTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (b *BaseThemeNight) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode baseThemeNight#b7b31ea8 to nil")
	}
	if err := buf.ConsumeID(BaseThemeNightTypeID); err != nil {
		return fmt.Errorf("unable to decode baseThemeNight#b7b31ea8: %w", err)
	}
	return nil
}

// construct implements constructor of BaseThemeClass.
func (b BaseThemeNight) construct() BaseThemeClass { return &b }

// Ensuring interfaces in compile-time for BaseThemeNight.
var (
	_ bin.Encoder = &BaseThemeNight{}
	_ bin.Decoder = &BaseThemeNight{}

	_ BaseThemeClass = &BaseThemeNight{}
)

// BaseThemeTinted represents TL type `baseThemeTinted#6d5f77ee`.
type BaseThemeTinted struct {
}

// BaseThemeTintedTypeID is TL type id of BaseThemeTinted.
const BaseThemeTintedTypeID = 0x6d5f77ee

// Encode implements bin.Encoder.
func (b *BaseThemeTinted) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode baseThemeTinted#6d5f77ee as nil")
	}
	buf.PutID(BaseThemeTintedTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (b *BaseThemeTinted) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode baseThemeTinted#6d5f77ee to nil")
	}
	if err := buf.ConsumeID(BaseThemeTintedTypeID); err != nil {
		return fmt.Errorf("unable to decode baseThemeTinted#6d5f77ee: %w", err)
	}
	return nil
}

// construct implements constructor of BaseThemeClass.
func (b BaseThemeTinted) construct() BaseThemeClass { return &b }

// Ensuring interfaces in compile-time for BaseThemeTinted.
var (
	_ bin.Encoder = &BaseThemeTinted{}
	_ bin.Decoder = &BaseThemeTinted{}

	_ BaseThemeClass = &BaseThemeTinted{}
)

// BaseThemeArctic represents TL type `baseThemeArctic#5b11125a`.
type BaseThemeArctic struct {
}

// BaseThemeArcticTypeID is TL type id of BaseThemeArctic.
const BaseThemeArcticTypeID = 0x5b11125a

// Encode implements bin.Encoder.
func (b *BaseThemeArctic) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode baseThemeArctic#5b11125a as nil")
	}
	buf.PutID(BaseThemeArcticTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (b *BaseThemeArctic) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode baseThemeArctic#5b11125a to nil")
	}
	if err := buf.ConsumeID(BaseThemeArcticTypeID); err != nil {
		return fmt.Errorf("unable to decode baseThemeArctic#5b11125a: %w", err)
	}
	return nil
}

// construct implements constructor of BaseThemeClass.
func (b BaseThemeArctic) construct() BaseThemeClass { return &b }

// Ensuring interfaces in compile-time for BaseThemeArctic.
var (
	_ bin.Encoder = &BaseThemeArctic{}
	_ bin.Decoder = &BaseThemeArctic{}

	_ BaseThemeClass = &BaseThemeArctic{}
)

// BaseThemeClass represents BaseTheme generic type.
//
// Example:
//  g, err := DecodeBaseTheme(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *BaseThemeClassic: // baseThemeClassic#c3a12462
//  case *BaseThemeDay: // baseThemeDay#fbd81688
//  case *BaseThemeNight: // baseThemeNight#b7b31ea8
//  case *BaseThemeTinted: // baseThemeTinted#6d5f77ee
//  case *BaseThemeArctic: // baseThemeArctic#5b11125a
//  default: panic(v)
//  }
type BaseThemeClass interface {
	bin.Encoder
	bin.Decoder
	construct() BaseThemeClass
}

// DecodeBaseTheme implements binary de-serialization for BaseThemeClass.
func DecodeBaseTheme(buf *bin.Buffer) (BaseThemeClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case BaseThemeClassicTypeID:
		// Decoding baseThemeClassic#c3a12462.
		v := BaseThemeClassic{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BaseThemeClass: %w", err)
		}
		return &v, nil
	case BaseThemeDayTypeID:
		// Decoding baseThemeDay#fbd81688.
		v := BaseThemeDay{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BaseThemeClass: %w", err)
		}
		return &v, nil
	case BaseThemeNightTypeID:
		// Decoding baseThemeNight#b7b31ea8.
		v := BaseThemeNight{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BaseThemeClass: %w", err)
		}
		return &v, nil
	case BaseThemeTintedTypeID:
		// Decoding baseThemeTinted#6d5f77ee.
		v := BaseThemeTinted{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BaseThemeClass: %w", err)
		}
		return &v, nil
	case BaseThemeArcticTypeID:
		// Decoding baseThemeArctic#5b11125a.
		v := BaseThemeArctic{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BaseThemeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode BaseThemeClass: %w", bin.NewUnexpectedID(id))
	}
}
