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

// SecureValue represents TL type `secureValue#187fa0ca`.
type SecureValue struct {
	// Flags field of SecureValue.
	Flags bin.Fields
	// Type field of SecureValue.
	Type SecureValueTypeClass
	// Data field of SecureValue.
	//
	// Use SetData and GetData helpers.
	Data SecureData
	// FrontSide field of SecureValue.
	//
	// Use SetFrontSide and GetFrontSide helpers.
	FrontSide SecureFileClass
	// ReverseSide field of SecureValue.
	//
	// Use SetReverseSide and GetReverseSide helpers.
	ReverseSide SecureFileClass
	// Selfie field of SecureValue.
	//
	// Use SetSelfie and GetSelfie helpers.
	Selfie SecureFileClass
	// Translation field of SecureValue.
	//
	// Use SetTranslation and GetTranslation helpers.
	Translation []SecureFileClass
	// Files field of SecureValue.
	//
	// Use SetFiles and GetFiles helpers.
	Files []SecureFileClass
	// PlainData field of SecureValue.
	//
	// Use SetPlainData and GetPlainData helpers.
	PlainData SecurePlainDataClass
	// Hash field of SecureValue.
	Hash []byte
}

// SecureValueTypeID is TL type id of SecureValue.
const SecureValueTypeID = 0x187fa0ca

// Encode implements bin.Encoder.
func (s *SecureValue) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode secureValue#187fa0ca as nil")
	}
	b.PutID(SecureValueTypeID)
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode secureValue#187fa0ca: field flags: %w", err)
	}
	if s.Type == nil {
		return fmt.Errorf("unable to encode secureValue#187fa0ca: field type is nil")
	}
	if err := s.Type.Encode(b); err != nil {
		return fmt.Errorf("unable to encode secureValue#187fa0ca: field type: %w", err)
	}
	if s.Flags.Has(0) {
		if err := s.Data.Encode(b); err != nil {
			return fmt.Errorf("unable to encode secureValue#187fa0ca: field data: %w", err)
		}
	}
	if s.Flags.Has(1) {
		if s.FrontSide == nil {
			return fmt.Errorf("unable to encode secureValue#187fa0ca: field front_side is nil")
		}
		if err := s.FrontSide.Encode(b); err != nil {
			return fmt.Errorf("unable to encode secureValue#187fa0ca: field front_side: %w", err)
		}
	}
	if s.Flags.Has(2) {
		if s.ReverseSide == nil {
			return fmt.Errorf("unable to encode secureValue#187fa0ca: field reverse_side is nil")
		}
		if err := s.ReverseSide.Encode(b); err != nil {
			return fmt.Errorf("unable to encode secureValue#187fa0ca: field reverse_side: %w", err)
		}
	}
	if s.Flags.Has(3) {
		if s.Selfie == nil {
			return fmt.Errorf("unable to encode secureValue#187fa0ca: field selfie is nil")
		}
		if err := s.Selfie.Encode(b); err != nil {
			return fmt.Errorf("unable to encode secureValue#187fa0ca: field selfie: %w", err)
		}
	}
	if s.Flags.Has(6) {
		b.PutVectorHeader(len(s.Translation))
		for idx, v := range s.Translation {
			if v == nil {
				return fmt.Errorf("unable to encode secureValue#187fa0ca: field translation element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode secureValue#187fa0ca: field translation element with index %d: %w", idx, err)
			}
		}
	}
	if s.Flags.Has(4) {
		b.PutVectorHeader(len(s.Files))
		for idx, v := range s.Files {
			if v == nil {
				return fmt.Errorf("unable to encode secureValue#187fa0ca: field files element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode secureValue#187fa0ca: field files element with index %d: %w", idx, err)
			}
		}
	}
	if s.Flags.Has(5) {
		if s.PlainData == nil {
			return fmt.Errorf("unable to encode secureValue#187fa0ca: field plain_data is nil")
		}
		if err := s.PlainData.Encode(b); err != nil {
			return fmt.Errorf("unable to encode secureValue#187fa0ca: field plain_data: %w", err)
		}
	}
	b.PutBytes(s.Hash)
	return nil
}

// SetData sets value of Data conditional field.
func (s *SecureValue) SetData(value SecureData) {
	s.Flags.Set(0)
	s.Data = value
}

// GetData returns value of Data conditional field and
// boolean which is true if field was set.
func (s *SecureValue) GetData() (value SecureData, ok bool) {
	if !s.Flags.Has(0) {
		return value, false
	}
	return s.Data, true
}

// SetFrontSide sets value of FrontSide conditional field.
func (s *SecureValue) SetFrontSide(value SecureFileClass) {
	s.Flags.Set(1)
	s.FrontSide = value
}

// GetFrontSide returns value of FrontSide conditional field and
// boolean which is true if field was set.
func (s *SecureValue) GetFrontSide() (value SecureFileClass, ok bool) {
	if !s.Flags.Has(1) {
		return value, false
	}
	return s.FrontSide, true
}

// SetReverseSide sets value of ReverseSide conditional field.
func (s *SecureValue) SetReverseSide(value SecureFileClass) {
	s.Flags.Set(2)
	s.ReverseSide = value
}

// GetReverseSide returns value of ReverseSide conditional field and
// boolean which is true if field was set.
func (s *SecureValue) GetReverseSide() (value SecureFileClass, ok bool) {
	if !s.Flags.Has(2) {
		return value, false
	}
	return s.ReverseSide, true
}

// SetSelfie sets value of Selfie conditional field.
func (s *SecureValue) SetSelfie(value SecureFileClass) {
	s.Flags.Set(3)
	s.Selfie = value
}

// GetSelfie returns value of Selfie conditional field and
// boolean which is true if field was set.
func (s *SecureValue) GetSelfie() (value SecureFileClass, ok bool) {
	if !s.Flags.Has(3) {
		return value, false
	}
	return s.Selfie, true
}

// SetTranslation sets value of Translation conditional field.
func (s *SecureValue) SetTranslation(value []SecureFileClass) {
	s.Flags.Set(6)
	s.Translation = value
}

// GetTranslation returns value of Translation conditional field and
// boolean which is true if field was set.
func (s *SecureValue) GetTranslation() (value []SecureFileClass, ok bool) {
	if !s.Flags.Has(6) {
		return value, false
	}
	return s.Translation, true
}

// SetFiles sets value of Files conditional field.
func (s *SecureValue) SetFiles(value []SecureFileClass) {
	s.Flags.Set(4)
	s.Files = value
}

// GetFiles returns value of Files conditional field and
// boolean which is true if field was set.
func (s *SecureValue) GetFiles() (value []SecureFileClass, ok bool) {
	if !s.Flags.Has(4) {
		return value, false
	}
	return s.Files, true
}

// SetPlainData sets value of PlainData conditional field.
func (s *SecureValue) SetPlainData(value SecurePlainDataClass) {
	s.Flags.Set(5)
	s.PlainData = value
}

// GetPlainData returns value of PlainData conditional field and
// boolean which is true if field was set.
func (s *SecureValue) GetPlainData() (value SecurePlainDataClass, ok bool) {
	if !s.Flags.Has(5) {
		return value, false
	}
	return s.PlainData, true
}

// Decode implements bin.Decoder.
func (s *SecureValue) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode secureValue#187fa0ca to nil")
	}
	if err := b.ConsumeID(SecureValueTypeID); err != nil {
		return fmt.Errorf("unable to decode secureValue#187fa0ca: %w", err)
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode secureValue#187fa0ca: field flags: %w", err)
		}
	}
	{
		value, err := DecodeSecureValueType(b)
		if err != nil {
			return fmt.Errorf("unable to decode secureValue#187fa0ca: field type: %w", err)
		}
		s.Type = value
	}
	if s.Flags.Has(0) {
		if err := s.Data.Decode(b); err != nil {
			return fmt.Errorf("unable to decode secureValue#187fa0ca: field data: %w", err)
		}
	}
	if s.Flags.Has(1) {
		value, err := DecodeSecureFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode secureValue#187fa0ca: field front_side: %w", err)
		}
		s.FrontSide = value
	}
	if s.Flags.Has(2) {
		value, err := DecodeSecureFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode secureValue#187fa0ca: field reverse_side: %w", err)
		}
		s.ReverseSide = value
	}
	if s.Flags.Has(3) {
		value, err := DecodeSecureFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode secureValue#187fa0ca: field selfie: %w", err)
		}
		s.Selfie = value
	}
	if s.Flags.Has(6) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode secureValue#187fa0ca: field translation: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeSecureFile(b)
			if err != nil {
				return fmt.Errorf("unable to decode secureValue#187fa0ca: field translation: %w", err)
			}
			s.Translation = append(s.Translation, value)
		}
	}
	if s.Flags.Has(4) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode secureValue#187fa0ca: field files: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeSecureFile(b)
			if err != nil {
				return fmt.Errorf("unable to decode secureValue#187fa0ca: field files: %w", err)
			}
			s.Files = append(s.Files, value)
		}
	}
	if s.Flags.Has(5) {
		value, err := DecodeSecurePlainData(b)
		if err != nil {
			return fmt.Errorf("unable to decode secureValue#187fa0ca: field plain_data: %w", err)
		}
		s.PlainData = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode secureValue#187fa0ca: field hash: %w", err)
		}
		s.Hash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for SecureValue.
var (
	_ bin.Encoder = &SecureValue{}
	_ bin.Decoder = &SecureValue{}
)
