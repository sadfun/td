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

// InputSecureValue represents TL type `inputSecureValue#db21d0a7`.
type InputSecureValue struct {
	// Flags field of InputSecureValue.
	Flags bin.Fields
	// Type field of InputSecureValue.
	Type SecureValueTypeClass
	// Data field of InputSecureValue.
	//
	// Use SetData and GetData helpers.
	Data SecureData
	// FrontSide field of InputSecureValue.
	//
	// Use SetFrontSide and GetFrontSide helpers.
	FrontSide InputSecureFileClass
	// ReverseSide field of InputSecureValue.
	//
	// Use SetReverseSide and GetReverseSide helpers.
	ReverseSide InputSecureFileClass
	// Selfie field of InputSecureValue.
	//
	// Use SetSelfie and GetSelfie helpers.
	Selfie InputSecureFileClass
	// Translation field of InputSecureValue.
	//
	// Use SetTranslation and GetTranslation helpers.
	Translation []InputSecureFileClass
	// Files field of InputSecureValue.
	//
	// Use SetFiles and GetFiles helpers.
	Files []InputSecureFileClass
	// PlainData field of InputSecureValue.
	//
	// Use SetPlainData and GetPlainData helpers.
	PlainData SecurePlainDataClass
}

// InputSecureValueTypeID is TL type id of InputSecureValue.
const InputSecureValueTypeID = 0xdb21d0a7

// Encode implements bin.Encoder.
func (i *InputSecureValue) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputSecureValue#db21d0a7 as nil")
	}
	b.PutID(InputSecureValueTypeID)
	if err := i.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputSecureValue#db21d0a7: field flags: %w", err)
	}
	if i.Type == nil {
		return fmt.Errorf("unable to encode inputSecureValue#db21d0a7: field type is nil")
	}
	if err := i.Type.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputSecureValue#db21d0a7: field type: %w", err)
	}
	if i.Flags.Has(0) {
		if err := i.Data.Encode(b); err != nil {
			return fmt.Errorf("unable to encode inputSecureValue#db21d0a7: field data: %w", err)
		}
	}
	if i.Flags.Has(1) {
		if i.FrontSide == nil {
			return fmt.Errorf("unable to encode inputSecureValue#db21d0a7: field front_side is nil")
		}
		if err := i.FrontSide.Encode(b); err != nil {
			return fmt.Errorf("unable to encode inputSecureValue#db21d0a7: field front_side: %w", err)
		}
	}
	if i.Flags.Has(2) {
		if i.ReverseSide == nil {
			return fmt.Errorf("unable to encode inputSecureValue#db21d0a7: field reverse_side is nil")
		}
		if err := i.ReverseSide.Encode(b); err != nil {
			return fmt.Errorf("unable to encode inputSecureValue#db21d0a7: field reverse_side: %w", err)
		}
	}
	if i.Flags.Has(3) {
		if i.Selfie == nil {
			return fmt.Errorf("unable to encode inputSecureValue#db21d0a7: field selfie is nil")
		}
		if err := i.Selfie.Encode(b); err != nil {
			return fmt.Errorf("unable to encode inputSecureValue#db21d0a7: field selfie: %w", err)
		}
	}
	if i.Flags.Has(6) {
		b.PutVectorHeader(len(i.Translation))
		for idx, v := range i.Translation {
			if v == nil {
				return fmt.Errorf("unable to encode inputSecureValue#db21d0a7: field translation element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode inputSecureValue#db21d0a7: field translation element with index %d: %w", idx, err)
			}
		}
	}
	if i.Flags.Has(4) {
		b.PutVectorHeader(len(i.Files))
		for idx, v := range i.Files {
			if v == nil {
				return fmt.Errorf("unable to encode inputSecureValue#db21d0a7: field files element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode inputSecureValue#db21d0a7: field files element with index %d: %w", idx, err)
			}
		}
	}
	if i.Flags.Has(5) {
		if i.PlainData == nil {
			return fmt.Errorf("unable to encode inputSecureValue#db21d0a7: field plain_data is nil")
		}
		if err := i.PlainData.Encode(b); err != nil {
			return fmt.Errorf("unable to encode inputSecureValue#db21d0a7: field plain_data: %w", err)
		}
	}
	return nil
}

// SetData sets value of Data conditional field.
func (i *InputSecureValue) SetData(value SecureData) {
	i.Flags.Set(0)
	i.Data = value
}

// GetData returns value of Data conditional field and
// boolean which is true if field was set.
func (i *InputSecureValue) GetData() (value SecureData, ok bool) {
	if !i.Flags.Has(0) {
		return value, false
	}
	return i.Data, true
}

// SetFrontSide sets value of FrontSide conditional field.
func (i *InputSecureValue) SetFrontSide(value InputSecureFileClass) {
	i.Flags.Set(1)
	i.FrontSide = value
}

// GetFrontSide returns value of FrontSide conditional field and
// boolean which is true if field was set.
func (i *InputSecureValue) GetFrontSide() (value InputSecureFileClass, ok bool) {
	if !i.Flags.Has(1) {
		return value, false
	}
	return i.FrontSide, true
}

// SetReverseSide sets value of ReverseSide conditional field.
func (i *InputSecureValue) SetReverseSide(value InputSecureFileClass) {
	i.Flags.Set(2)
	i.ReverseSide = value
}

// GetReverseSide returns value of ReverseSide conditional field and
// boolean which is true if field was set.
func (i *InputSecureValue) GetReverseSide() (value InputSecureFileClass, ok bool) {
	if !i.Flags.Has(2) {
		return value, false
	}
	return i.ReverseSide, true
}

// SetSelfie sets value of Selfie conditional field.
func (i *InputSecureValue) SetSelfie(value InputSecureFileClass) {
	i.Flags.Set(3)
	i.Selfie = value
}

// GetSelfie returns value of Selfie conditional field and
// boolean which is true if field was set.
func (i *InputSecureValue) GetSelfie() (value InputSecureFileClass, ok bool) {
	if !i.Flags.Has(3) {
		return value, false
	}
	return i.Selfie, true
}

// SetTranslation sets value of Translation conditional field.
func (i *InputSecureValue) SetTranslation(value []InputSecureFileClass) {
	i.Flags.Set(6)
	i.Translation = value
}

// GetTranslation returns value of Translation conditional field and
// boolean which is true if field was set.
func (i *InputSecureValue) GetTranslation() (value []InputSecureFileClass, ok bool) {
	if !i.Flags.Has(6) {
		return value, false
	}
	return i.Translation, true
}

// SetFiles sets value of Files conditional field.
func (i *InputSecureValue) SetFiles(value []InputSecureFileClass) {
	i.Flags.Set(4)
	i.Files = value
}

// GetFiles returns value of Files conditional field and
// boolean which is true if field was set.
func (i *InputSecureValue) GetFiles() (value []InputSecureFileClass, ok bool) {
	if !i.Flags.Has(4) {
		return value, false
	}
	return i.Files, true
}

// SetPlainData sets value of PlainData conditional field.
func (i *InputSecureValue) SetPlainData(value SecurePlainDataClass) {
	i.Flags.Set(5)
	i.PlainData = value
}

// GetPlainData returns value of PlainData conditional field and
// boolean which is true if field was set.
func (i *InputSecureValue) GetPlainData() (value SecurePlainDataClass, ok bool) {
	if !i.Flags.Has(5) {
		return value, false
	}
	return i.PlainData, true
}

// Decode implements bin.Decoder.
func (i *InputSecureValue) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputSecureValue#db21d0a7 to nil")
	}
	if err := b.ConsumeID(InputSecureValueTypeID); err != nil {
		return fmt.Errorf("unable to decode inputSecureValue#db21d0a7: %w", err)
	}
	{
		if err := i.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputSecureValue#db21d0a7: field flags: %w", err)
		}
	}
	{
		value, err := DecodeSecureValueType(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputSecureValue#db21d0a7: field type: %w", err)
		}
		i.Type = value
	}
	if i.Flags.Has(0) {
		if err := i.Data.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputSecureValue#db21d0a7: field data: %w", err)
		}
	}
	if i.Flags.Has(1) {
		value, err := DecodeInputSecureFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputSecureValue#db21d0a7: field front_side: %w", err)
		}
		i.FrontSide = value
	}
	if i.Flags.Has(2) {
		value, err := DecodeInputSecureFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputSecureValue#db21d0a7: field reverse_side: %w", err)
		}
		i.ReverseSide = value
	}
	if i.Flags.Has(3) {
		value, err := DecodeInputSecureFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputSecureValue#db21d0a7: field selfie: %w", err)
		}
		i.Selfie = value
	}
	if i.Flags.Has(6) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode inputSecureValue#db21d0a7: field translation: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputSecureFile(b)
			if err != nil {
				return fmt.Errorf("unable to decode inputSecureValue#db21d0a7: field translation: %w", err)
			}
			i.Translation = append(i.Translation, value)
		}
	}
	if i.Flags.Has(4) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode inputSecureValue#db21d0a7: field files: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputSecureFile(b)
			if err != nil {
				return fmt.Errorf("unable to decode inputSecureValue#db21d0a7: field files: %w", err)
			}
			i.Files = append(i.Files, value)
		}
	}
	if i.Flags.Has(5) {
		value, err := DecodeSecurePlainData(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputSecureValue#db21d0a7: field plain_data: %w", err)
		}
		i.PlainData = value
	}
	return nil
}

// Ensuring interfaces in compile-time for InputSecureValue.
var (
	_ bin.Encoder = &InputSecureValue{}
	_ bin.Decoder = &InputSecureValue{}
)
