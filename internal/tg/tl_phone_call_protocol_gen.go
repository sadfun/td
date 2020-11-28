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

// PhoneCallProtocol represents TL type `phoneCallProtocol#fc878fc8`.
type PhoneCallProtocol struct {
	// Flags field of PhoneCallProtocol.
	Flags bin.Fields
	// UDPP2P field of PhoneCallProtocol.
	UDPP2P bool
	// UDPReflector field of PhoneCallProtocol.
	UDPReflector bool
	// MinLayer field of PhoneCallProtocol.
	MinLayer int
	// MaxLayer field of PhoneCallProtocol.
	MaxLayer int
	// LibraryVersions field of PhoneCallProtocol.
	LibraryVersions []string
}

// PhoneCallProtocolTypeID is TL type id of PhoneCallProtocol.
const PhoneCallProtocolTypeID = 0xfc878fc8

// Encode implements bin.Encoder.
func (p *PhoneCallProtocol) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode phoneCallProtocol#fc878fc8 as nil")
	}
	b.PutID(PhoneCallProtocolTypeID)
	if err := p.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phoneCallProtocol#fc878fc8: field flags: %w", err)
	}
	b.PutInt(p.MinLayer)
	b.PutInt(p.MaxLayer)
	b.PutVectorHeader(len(p.LibraryVersions))
	for _, v := range p.LibraryVersions {
		b.PutString(v)
	}
	return nil
}

// SetUDPP2P sets value of UDPP2P conditional field.
func (p *PhoneCallProtocol) SetUDPP2P(value bool) {
	if value {
		p.Flags.Set(0)
	} else {
		p.Flags.Unset(0)
	}
}

// SetUDPReflector sets value of UDPReflector conditional field.
func (p *PhoneCallProtocol) SetUDPReflector(value bool) {
	if value {
		p.Flags.Set(1)
	} else {
		p.Flags.Unset(1)
	}
}

// Decode implements bin.Decoder.
func (p *PhoneCallProtocol) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode phoneCallProtocol#fc878fc8 to nil")
	}
	if err := b.ConsumeID(PhoneCallProtocolTypeID); err != nil {
		return fmt.Errorf("unable to decode phoneCallProtocol#fc878fc8: %w", err)
	}
	{
		if err := p.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phoneCallProtocol#fc878fc8: field flags: %w", err)
		}
	}
	p.UDPP2P = p.Flags.Has(0)
	p.UDPReflector = p.Flags.Has(1)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode phoneCallProtocol#fc878fc8: field min_layer: %w", err)
		}
		p.MinLayer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode phoneCallProtocol#fc878fc8: field max_layer: %w", err)
		}
		p.MaxLayer = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode phoneCallProtocol#fc878fc8: field library_versions: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode phoneCallProtocol#fc878fc8: field library_versions: %w", err)
			}
			p.LibraryVersions = append(p.LibraryVersions, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for PhoneCallProtocol.
var (
	_ bin.Encoder = &PhoneCallProtocol{}
	_ bin.Decoder = &PhoneCallProtocol{}
)
