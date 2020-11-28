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

// PaymentsBankCardData represents TL type `payments.bankCardData#3e24e573`.
type PaymentsBankCardData struct {
	// Title field of PaymentsBankCardData.
	Title string
	// OpenUrls field of PaymentsBankCardData.
	OpenUrls []BankCardOpenUrl
}

// PaymentsBankCardDataTypeID is TL type id of PaymentsBankCardData.
const PaymentsBankCardDataTypeID = 0x3e24e573

// Encode implements bin.Encoder.
func (b *PaymentsBankCardData) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode payments.bankCardData#3e24e573 as nil")
	}
	buf.PutID(PaymentsBankCardDataTypeID)
	buf.PutString(b.Title)
	buf.PutVectorHeader(len(b.OpenUrls))
	for idx, v := range b.OpenUrls {
		if err := v.Encode(buf); err != nil {
			return fmt.Errorf("unable to encode payments.bankCardData#3e24e573: field open_urls element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (b *PaymentsBankCardData) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode payments.bankCardData#3e24e573 to nil")
	}
	if err := buf.ConsumeID(PaymentsBankCardDataTypeID); err != nil {
		return fmt.Errorf("unable to decode payments.bankCardData#3e24e573: %w", err)
	}
	{
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode payments.bankCardData#3e24e573: field title: %w", err)
		}
		b.Title = value
	}
	{
		headerLen, err := buf.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode payments.bankCardData#3e24e573: field open_urls: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value BankCardOpenUrl
			if err := value.Decode(buf); err != nil {
				return fmt.Errorf("unable to decode payments.bankCardData#3e24e573: field open_urls: %w", err)
			}
			b.OpenUrls = append(b.OpenUrls, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for PaymentsBankCardData.
var (
	_ bin.Encoder = &PaymentsBankCardData{}
	_ bin.Decoder = &PaymentsBankCardData{}
)
