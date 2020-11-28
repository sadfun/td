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

// PaymentsPaymentReceipt represents TL type `payments.paymentReceipt#500911e1`.
type PaymentsPaymentReceipt struct {
	// Flags field of PaymentsPaymentReceipt.
	Flags bin.Fields
	// Date field of PaymentsPaymentReceipt.
	Date int
	// BotID field of PaymentsPaymentReceipt.
	BotID int
	// Invoice field of PaymentsPaymentReceipt.
	Invoice Invoice
	// ProviderID field of PaymentsPaymentReceipt.
	ProviderID int
	// Info field of PaymentsPaymentReceipt.
	//
	// Use SetInfo and GetInfo helpers.
	Info PaymentRequestedInfo
	// Shipping field of PaymentsPaymentReceipt.
	//
	// Use SetShipping and GetShipping helpers.
	Shipping ShippingOption
	// Currency field of PaymentsPaymentReceipt.
	Currency string
	// TotalAmount field of PaymentsPaymentReceipt.
	TotalAmount int64
	// CredentialsTitle field of PaymentsPaymentReceipt.
	CredentialsTitle string
	// Users field of PaymentsPaymentReceipt.
	Users []UserClass
}

// PaymentsPaymentReceiptTypeID is TL type id of PaymentsPaymentReceipt.
const PaymentsPaymentReceiptTypeID = 0x500911e1

// Encode implements bin.Encoder.
func (p *PaymentsPaymentReceipt) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode payments.paymentReceipt#500911e1 as nil")
	}
	b.PutID(PaymentsPaymentReceiptTypeID)
	if err := p.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode payments.paymentReceipt#500911e1: field flags: %w", err)
	}
	b.PutInt(p.Date)
	b.PutInt(p.BotID)
	if err := p.Invoice.Encode(b); err != nil {
		return fmt.Errorf("unable to encode payments.paymentReceipt#500911e1: field invoice: %w", err)
	}
	b.PutInt(p.ProviderID)
	if p.Flags.Has(0) {
		if err := p.Info.Encode(b); err != nil {
			return fmt.Errorf("unable to encode payments.paymentReceipt#500911e1: field info: %w", err)
		}
	}
	if p.Flags.Has(1) {
		if err := p.Shipping.Encode(b); err != nil {
			return fmt.Errorf("unable to encode payments.paymentReceipt#500911e1: field shipping: %w", err)
		}
	}
	b.PutString(p.Currency)
	b.PutLong(p.TotalAmount)
	b.PutString(p.CredentialsTitle)
	b.PutVectorHeader(len(p.Users))
	for idx, v := range p.Users {
		if v == nil {
			return fmt.Errorf("unable to encode payments.paymentReceipt#500911e1: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode payments.paymentReceipt#500911e1: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// SetInfo sets value of Info conditional field.
func (p *PaymentsPaymentReceipt) SetInfo(value PaymentRequestedInfo) {
	p.Flags.Set(0)
	p.Info = value
}

// GetInfo returns value of Info conditional field and
// boolean which is true if field was set.
func (p *PaymentsPaymentReceipt) GetInfo() (value PaymentRequestedInfo, ok bool) {
	if !p.Flags.Has(0) {
		return value, false
	}
	return p.Info, true
}

// SetShipping sets value of Shipping conditional field.
func (p *PaymentsPaymentReceipt) SetShipping(value ShippingOption) {
	p.Flags.Set(1)
	p.Shipping = value
}

// GetShipping returns value of Shipping conditional field and
// boolean which is true if field was set.
func (p *PaymentsPaymentReceipt) GetShipping() (value ShippingOption, ok bool) {
	if !p.Flags.Has(1) {
		return value, false
	}
	return p.Shipping, true
}

// Decode implements bin.Decoder.
func (p *PaymentsPaymentReceipt) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode payments.paymentReceipt#500911e1 to nil")
	}
	if err := b.ConsumeID(PaymentsPaymentReceiptTypeID); err != nil {
		return fmt.Errorf("unable to decode payments.paymentReceipt#500911e1: %w", err)
	}
	{
		if err := p.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode payments.paymentReceipt#500911e1: field flags: %w", err)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode payments.paymentReceipt#500911e1: field date: %w", err)
		}
		p.Date = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode payments.paymentReceipt#500911e1: field bot_id: %w", err)
		}
		p.BotID = value
	}
	{
		if err := p.Invoice.Decode(b); err != nil {
			return fmt.Errorf("unable to decode payments.paymentReceipt#500911e1: field invoice: %w", err)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode payments.paymentReceipt#500911e1: field provider_id: %w", err)
		}
		p.ProviderID = value
	}
	if p.Flags.Has(0) {
		if err := p.Info.Decode(b); err != nil {
			return fmt.Errorf("unable to decode payments.paymentReceipt#500911e1: field info: %w", err)
		}
	}
	if p.Flags.Has(1) {
		if err := p.Shipping.Decode(b); err != nil {
			return fmt.Errorf("unable to decode payments.paymentReceipt#500911e1: field shipping: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode payments.paymentReceipt#500911e1: field currency: %w", err)
		}
		p.Currency = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode payments.paymentReceipt#500911e1: field total_amount: %w", err)
		}
		p.TotalAmount = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode payments.paymentReceipt#500911e1: field credentials_title: %w", err)
		}
		p.CredentialsTitle = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode payments.paymentReceipt#500911e1: field users: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode payments.paymentReceipt#500911e1: field users: %w", err)
			}
			p.Users = append(p.Users, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for PaymentsPaymentReceipt.
var (
	_ bin.Encoder = &PaymentsPaymentReceipt{}
	_ bin.Decoder = &PaymentsPaymentReceipt{}
)
