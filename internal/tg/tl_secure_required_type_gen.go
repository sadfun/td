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

// SecureRequiredType represents TL type `secureRequiredType#829d99da`.
type SecureRequiredType struct {
	// Flags field of SecureRequiredType.
	Flags bin.Fields
	// NativeNames field of SecureRequiredType.
	NativeNames bool
	// SelfieRequired field of SecureRequiredType.
	SelfieRequired bool
	// TranslationRequired field of SecureRequiredType.
	TranslationRequired bool
	// Type field of SecureRequiredType.
	Type SecureValueTypeClass
}

// SecureRequiredTypeTypeID is TL type id of SecureRequiredType.
const SecureRequiredTypeTypeID = 0x829d99da

// Encode implements bin.Encoder.
func (s *SecureRequiredType) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode secureRequiredType#829d99da as nil")
	}
	b.PutID(SecureRequiredTypeTypeID)
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode secureRequiredType#829d99da: field flags: %w", err)
	}
	if s.Type == nil {
		return fmt.Errorf("unable to encode secureRequiredType#829d99da: field type is nil")
	}
	if err := s.Type.Encode(b); err != nil {
		return fmt.Errorf("unable to encode secureRequiredType#829d99da: field type: %w", err)
	}
	return nil
}

// SetNativeNames sets value of NativeNames conditional field.
func (s *SecureRequiredType) SetNativeNames(value bool) {
	if value {
		s.Flags.Set(0)
	} else {
		s.Flags.Unset(0)
	}
}

// SetSelfieRequired sets value of SelfieRequired conditional field.
func (s *SecureRequiredType) SetSelfieRequired(value bool) {
	if value {
		s.Flags.Set(1)
	} else {
		s.Flags.Unset(1)
	}
}

// SetTranslationRequired sets value of TranslationRequired conditional field.
func (s *SecureRequiredType) SetTranslationRequired(value bool) {
	if value {
		s.Flags.Set(2)
	} else {
		s.Flags.Unset(2)
	}
}

// Decode implements bin.Decoder.
func (s *SecureRequiredType) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode secureRequiredType#829d99da to nil")
	}
	if err := b.ConsumeID(SecureRequiredTypeTypeID); err != nil {
		return fmt.Errorf("unable to decode secureRequiredType#829d99da: %w", err)
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode secureRequiredType#829d99da: field flags: %w", err)
		}
	}
	s.NativeNames = s.Flags.Has(0)
	s.SelfieRequired = s.Flags.Has(1)
	s.TranslationRequired = s.Flags.Has(2)
	{
		value, err := DecodeSecureValueType(b)
		if err != nil {
			return fmt.Errorf("unable to decode secureRequiredType#829d99da: field type: %w", err)
		}
		s.Type = value
	}
	return nil
}

// construct implements constructor of SecureRequiredTypeClass.
func (s SecureRequiredType) construct() SecureRequiredTypeClass { return &s }

// Ensuring interfaces in compile-time for SecureRequiredType.
var (
	_ bin.Encoder = &SecureRequiredType{}
	_ bin.Decoder = &SecureRequiredType{}

	_ SecureRequiredTypeClass = &SecureRequiredType{}
)

// SecureRequiredTypeOneOf represents TL type `secureRequiredTypeOneOf#27477b4`.
type SecureRequiredTypeOneOf struct {
	// Types field of SecureRequiredTypeOneOf.
	Types []SecureRequiredTypeClass
}

// SecureRequiredTypeOneOfTypeID is TL type id of SecureRequiredTypeOneOf.
const SecureRequiredTypeOneOfTypeID = 0x27477b4

// Encode implements bin.Encoder.
func (s *SecureRequiredTypeOneOf) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode secureRequiredTypeOneOf#27477b4 as nil")
	}
	b.PutID(SecureRequiredTypeOneOfTypeID)
	b.PutVectorHeader(len(s.Types))
	for idx, v := range s.Types {
		if v == nil {
			return fmt.Errorf("unable to encode secureRequiredTypeOneOf#27477b4: field types element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode secureRequiredTypeOneOf#27477b4: field types element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SecureRequiredTypeOneOf) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode secureRequiredTypeOneOf#27477b4 to nil")
	}
	if err := b.ConsumeID(SecureRequiredTypeOneOfTypeID); err != nil {
		return fmt.Errorf("unable to decode secureRequiredTypeOneOf#27477b4: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode secureRequiredTypeOneOf#27477b4: field types: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeSecureRequiredType(b)
			if err != nil {
				return fmt.Errorf("unable to decode secureRequiredTypeOneOf#27477b4: field types: %w", err)
			}
			s.Types = append(s.Types, value)
		}
	}
	return nil
}

// construct implements constructor of SecureRequiredTypeClass.
func (s SecureRequiredTypeOneOf) construct() SecureRequiredTypeClass { return &s }

// Ensuring interfaces in compile-time for SecureRequiredTypeOneOf.
var (
	_ bin.Encoder = &SecureRequiredTypeOneOf{}
	_ bin.Decoder = &SecureRequiredTypeOneOf{}

	_ SecureRequiredTypeClass = &SecureRequiredTypeOneOf{}
)

// SecureRequiredTypeClass represents SecureRequiredType generic type.
//
// Example:
//  g, err := DecodeSecureRequiredType(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *SecureRequiredType: // secureRequiredType#829d99da
//  case *SecureRequiredTypeOneOf: // secureRequiredTypeOneOf#27477b4
//  default: panic(v)
//  }
type SecureRequiredTypeClass interface {
	bin.Encoder
	bin.Decoder
	construct() SecureRequiredTypeClass
}

// DecodeSecureRequiredType implements binary de-serialization for SecureRequiredTypeClass.
func DecodeSecureRequiredType(buf *bin.Buffer) (SecureRequiredTypeClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case SecureRequiredTypeTypeID:
		// Decoding secureRequiredType#829d99da.
		v := SecureRequiredType{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SecureRequiredTypeClass: %w", err)
		}
		return &v, nil
	case SecureRequiredTypeOneOfTypeID:
		// Decoding secureRequiredTypeOneOf#27477b4.
		v := SecureRequiredTypeOneOf{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SecureRequiredTypeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode SecureRequiredTypeClass: %w", bin.NewUnexpectedID(id))
	}
}
