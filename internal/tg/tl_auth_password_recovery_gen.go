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

// AuthPasswordRecovery represents TL type `auth.passwordRecovery#137948a5`.
type AuthPasswordRecovery struct {
	// EmailPattern field of AuthPasswordRecovery.
	EmailPattern string
}

// AuthPasswordRecoveryTypeID is TL type id of AuthPasswordRecovery.
const AuthPasswordRecoveryTypeID = 0x137948a5

// Encode implements bin.Encoder.
func (p *AuthPasswordRecovery) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode auth.passwordRecovery#137948a5 as nil")
	}
	b.PutID(AuthPasswordRecoveryTypeID)
	b.PutString(p.EmailPattern)
	return nil
}

// Decode implements bin.Decoder.
func (p *AuthPasswordRecovery) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode auth.passwordRecovery#137948a5 to nil")
	}
	if err := b.ConsumeID(AuthPasswordRecoveryTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.passwordRecovery#137948a5: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode auth.passwordRecovery#137948a5: field email_pattern: %w", err)
		}
		p.EmailPattern = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AuthPasswordRecovery.
var (
	_ bin.Encoder = &AuthPasswordRecovery{}
	_ bin.Decoder = &AuthPasswordRecovery{}
)
