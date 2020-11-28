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

// Poll represents TL type `poll#86e18161`.
type Poll struct {
	// ID field of Poll.
	ID int64
	// Flags field of Poll.
	Flags bin.Fields
	// Closed field of Poll.
	Closed bool
	// PublicVoters field of Poll.
	PublicVoters bool
	// MultipleChoice field of Poll.
	MultipleChoice bool
	// Quiz field of Poll.
	Quiz bool
	// Question field of Poll.
	Question string
	// Answers field of Poll.
	Answers []PollAnswer
	// ClosePeriod field of Poll.
	//
	// Use SetClosePeriod and GetClosePeriod helpers.
	ClosePeriod int
	// CloseDate field of Poll.
	//
	// Use SetCloseDate and GetCloseDate helpers.
	CloseDate int
}

// PollTypeID is TL type id of Poll.
const PollTypeID = 0x86e18161

// Encode implements bin.Encoder.
func (p *Poll) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode poll#86e18161 as nil")
	}
	b.PutID(PollTypeID)
	b.PutLong(p.ID)
	if err := p.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode poll#86e18161: field flags: %w", err)
	}
	b.PutString(p.Question)
	b.PutVectorHeader(len(p.Answers))
	for idx, v := range p.Answers {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode poll#86e18161: field answers element with index %d: %w", idx, err)
		}
	}
	if p.Flags.Has(4) {
		b.PutInt(p.ClosePeriod)
	}
	if p.Flags.Has(5) {
		b.PutInt(p.CloseDate)
	}
	return nil
}

// SetClosed sets value of Closed conditional field.
func (p *Poll) SetClosed(value bool) {
	if value {
		p.Flags.Set(0)
	} else {
		p.Flags.Unset(0)
	}
}

// SetPublicVoters sets value of PublicVoters conditional field.
func (p *Poll) SetPublicVoters(value bool) {
	if value {
		p.Flags.Set(1)
	} else {
		p.Flags.Unset(1)
	}
}

// SetMultipleChoice sets value of MultipleChoice conditional field.
func (p *Poll) SetMultipleChoice(value bool) {
	if value {
		p.Flags.Set(2)
	} else {
		p.Flags.Unset(2)
	}
}

// SetQuiz sets value of Quiz conditional field.
func (p *Poll) SetQuiz(value bool) {
	if value {
		p.Flags.Set(3)
	} else {
		p.Flags.Unset(3)
	}
}

// SetClosePeriod sets value of ClosePeriod conditional field.
func (p *Poll) SetClosePeriod(value int) {
	p.Flags.Set(4)
	p.ClosePeriod = value
}

// GetClosePeriod returns value of ClosePeriod conditional field and
// boolean which is true if field was set.
func (p *Poll) GetClosePeriod() (value int, ok bool) {
	if !p.Flags.Has(4) {
		return value, false
	}
	return p.ClosePeriod, true
}

// SetCloseDate sets value of CloseDate conditional field.
func (p *Poll) SetCloseDate(value int) {
	p.Flags.Set(5)
	p.CloseDate = value
}

// GetCloseDate returns value of CloseDate conditional field and
// boolean which is true if field was set.
func (p *Poll) GetCloseDate() (value int, ok bool) {
	if !p.Flags.Has(5) {
		return value, false
	}
	return p.CloseDate, true
}

// Decode implements bin.Decoder.
func (p *Poll) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode poll#86e18161 to nil")
	}
	if err := b.ConsumeID(PollTypeID); err != nil {
		return fmt.Errorf("unable to decode poll#86e18161: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode poll#86e18161: field id: %w", err)
		}
		p.ID = value
	}
	{
		if err := p.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode poll#86e18161: field flags: %w", err)
		}
	}
	p.Closed = p.Flags.Has(0)
	p.PublicVoters = p.Flags.Has(1)
	p.MultipleChoice = p.Flags.Has(2)
	p.Quiz = p.Flags.Has(3)
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode poll#86e18161: field question: %w", err)
		}
		p.Question = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode poll#86e18161: field answers: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value PollAnswer
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode poll#86e18161: field answers: %w", err)
			}
			p.Answers = append(p.Answers, value)
		}
	}
	if p.Flags.Has(4) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode poll#86e18161: field close_period: %w", err)
		}
		p.ClosePeriod = value
	}
	if p.Flags.Has(5) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode poll#86e18161: field close_date: %w", err)
		}
		p.CloseDate = value
	}
	return nil
}

// Ensuring interfaces in compile-time for Poll.
var (
	_ bin.Encoder = &Poll{}
	_ bin.Decoder = &Poll{}
)
