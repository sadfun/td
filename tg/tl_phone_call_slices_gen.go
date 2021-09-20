//go:build !no_gotd_slices
// +build !no_gotd_slices

// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
)

// PhoneCallClassArray is adapter for slice of PhoneCallClass.
type PhoneCallClassArray []PhoneCallClass

// Sort sorts slice of PhoneCallClass.
func (s PhoneCallClassArray) Sort(less func(a, b PhoneCallClass) bool) PhoneCallClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of PhoneCallClass.
func (s PhoneCallClassArray) SortStable(less func(a, b PhoneCallClass) bool) PhoneCallClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of PhoneCallClass.
func (s PhoneCallClassArray) Retain(keep func(x PhoneCallClass) bool) PhoneCallClassArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s PhoneCallClassArray) First() (v PhoneCallClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s PhoneCallClassArray) Last() (v PhoneCallClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *PhoneCallClassArray) PopFirst() (v PhoneCallClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero PhoneCallClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *PhoneCallClassArray) Pop() (v PhoneCallClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByID sorts slice of PhoneCallClass by ID.
func (s PhoneCallClassArray) SortByID() PhoneCallClassArray {
	return s.Sort(func(a, b PhoneCallClass) bool {
		return a.GetID() < b.GetID()
	})
}

// SortStableByID sorts slice of PhoneCallClass by ID.
func (s PhoneCallClassArray) SortStableByID() PhoneCallClassArray {
	return s.SortStable(func(a, b PhoneCallClass) bool {
		return a.GetID() < b.GetID()
	})
}

// FillPhoneCallEmptyMap fills only PhoneCallEmpty constructors to given map.
func (s PhoneCallClassArray) FillPhoneCallEmptyMap(to map[int64]*PhoneCallEmpty) {
	for _, elem := range s {
		value, ok := elem.(*PhoneCallEmpty)
		if !ok {
			continue
		}
		to[value.GetID()] = value
	}
}

// PhoneCallEmptyToMap collects only PhoneCallEmpty constructors to map.
func (s PhoneCallClassArray) PhoneCallEmptyToMap() map[int64]*PhoneCallEmpty {
	r := make(map[int64]*PhoneCallEmpty, len(s))
	s.FillPhoneCallEmptyMap(r)
	return r
}

// AsPhoneCallEmpty returns copy with only PhoneCallEmpty constructors.
func (s PhoneCallClassArray) AsPhoneCallEmpty() (to PhoneCallEmptyArray) {
	for _, elem := range s {
		value, ok := elem.(*PhoneCallEmpty)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// FillPhoneCallWaitingMap fills only PhoneCallWaiting constructors to given map.
func (s PhoneCallClassArray) FillPhoneCallWaitingMap(to map[int64]*PhoneCallWaiting) {
	for _, elem := range s {
		value, ok := elem.(*PhoneCallWaiting)
		if !ok {
			continue
		}
		to[value.GetID()] = value
	}
}

// PhoneCallWaitingToMap collects only PhoneCallWaiting constructors to map.
func (s PhoneCallClassArray) PhoneCallWaitingToMap() map[int64]*PhoneCallWaiting {
	r := make(map[int64]*PhoneCallWaiting, len(s))
	s.FillPhoneCallWaitingMap(r)
	return r
}

// AsPhoneCallWaiting returns copy with only PhoneCallWaiting constructors.
func (s PhoneCallClassArray) AsPhoneCallWaiting() (to PhoneCallWaitingArray) {
	for _, elem := range s {
		value, ok := elem.(*PhoneCallWaiting)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// FillPhoneCallRequestedMap fills only PhoneCallRequested constructors to given map.
func (s PhoneCallClassArray) FillPhoneCallRequestedMap(to map[int64]*PhoneCallRequested) {
	for _, elem := range s {
		value, ok := elem.(*PhoneCallRequested)
		if !ok {
			continue
		}
		to[value.GetID()] = value
	}
}

// PhoneCallRequestedToMap collects only PhoneCallRequested constructors to map.
func (s PhoneCallClassArray) PhoneCallRequestedToMap() map[int64]*PhoneCallRequested {
	r := make(map[int64]*PhoneCallRequested, len(s))
	s.FillPhoneCallRequestedMap(r)
	return r
}

// AsPhoneCallRequested returns copy with only PhoneCallRequested constructors.
func (s PhoneCallClassArray) AsPhoneCallRequested() (to PhoneCallRequestedArray) {
	for _, elem := range s {
		value, ok := elem.(*PhoneCallRequested)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// FillPhoneCallAcceptedMap fills only PhoneCallAccepted constructors to given map.
func (s PhoneCallClassArray) FillPhoneCallAcceptedMap(to map[int64]*PhoneCallAccepted) {
	for _, elem := range s {
		value, ok := elem.(*PhoneCallAccepted)
		if !ok {
			continue
		}
		to[value.GetID()] = value
	}
}

// PhoneCallAcceptedToMap collects only PhoneCallAccepted constructors to map.
func (s PhoneCallClassArray) PhoneCallAcceptedToMap() map[int64]*PhoneCallAccepted {
	r := make(map[int64]*PhoneCallAccepted, len(s))
	s.FillPhoneCallAcceptedMap(r)
	return r
}

// AsPhoneCallAccepted returns copy with only PhoneCallAccepted constructors.
func (s PhoneCallClassArray) AsPhoneCallAccepted() (to PhoneCallAcceptedArray) {
	for _, elem := range s {
		value, ok := elem.(*PhoneCallAccepted)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// FillPhoneCallMap fills only PhoneCall constructors to given map.
func (s PhoneCallClassArray) FillPhoneCallMap(to map[int64]*PhoneCall) {
	for _, elem := range s {
		value, ok := elem.(*PhoneCall)
		if !ok {
			continue
		}
		to[value.GetID()] = value
	}
}

// PhoneCallToMap collects only PhoneCall constructors to map.
func (s PhoneCallClassArray) PhoneCallToMap() map[int64]*PhoneCall {
	r := make(map[int64]*PhoneCall, len(s))
	s.FillPhoneCallMap(r)
	return r
}

// AsPhoneCall returns copy with only PhoneCall constructors.
func (s PhoneCallClassArray) AsPhoneCall() (to PhoneCallArray) {
	for _, elem := range s {
		value, ok := elem.(*PhoneCall)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// FillPhoneCallDiscardedMap fills only PhoneCallDiscarded constructors to given map.
func (s PhoneCallClassArray) FillPhoneCallDiscardedMap(to map[int64]*PhoneCallDiscarded) {
	for _, elem := range s {
		value, ok := elem.(*PhoneCallDiscarded)
		if !ok {
			continue
		}
		to[value.GetID()] = value
	}
}

// PhoneCallDiscardedToMap collects only PhoneCallDiscarded constructors to map.
func (s PhoneCallClassArray) PhoneCallDiscardedToMap() map[int64]*PhoneCallDiscarded {
	r := make(map[int64]*PhoneCallDiscarded, len(s))
	s.FillPhoneCallDiscardedMap(r)
	return r
}

// AsPhoneCallDiscarded returns copy with only PhoneCallDiscarded constructors.
func (s PhoneCallClassArray) AsPhoneCallDiscarded() (to PhoneCallDiscardedArray) {
	for _, elem := range s {
		value, ok := elem.(*PhoneCallDiscarded)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// FillNotEmptyMap fills only NotEmpty constructors to given map.
func (s PhoneCallClassArray) FillNotEmptyMap(to map[int64]NotEmptyPhoneCall) {
	for _, elem := range s {
		value, ok := elem.AsNotEmpty()
		if !ok {
			continue
		}
		to[value.GetID()] = value
	}
}

// NotEmptyToMap collects only NotEmpty constructors to map.
func (s PhoneCallClassArray) NotEmptyToMap() map[int64]NotEmptyPhoneCall {
	r := make(map[int64]NotEmptyPhoneCall, len(s))
	s.FillNotEmptyMap(r)
	return r
}

// AppendOnlyNotEmpty appends only NotEmpty constructors to
// given slice.
func (s PhoneCallClassArray) AppendOnlyNotEmpty(to []NotEmptyPhoneCall) []NotEmptyPhoneCall {
	for _, elem := range s {
		value, ok := elem.AsNotEmpty()
		if !ok {
			continue
		}
		to = append(to, value)
	}

	return to
}

// AsNotEmpty returns copy with only NotEmpty constructors.
func (s PhoneCallClassArray) AsNotEmpty() (to []NotEmptyPhoneCall) {
	return s.AppendOnlyNotEmpty(to)
}

// FirstAsNotEmpty returns first element of slice (if exists).
func (s PhoneCallClassArray) FirstAsNotEmpty() (v NotEmptyPhoneCall, ok bool) {
	value, ok := s.First()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// LastAsNotEmpty returns last element of slice (if exists).
func (s PhoneCallClassArray) LastAsNotEmpty() (v NotEmptyPhoneCall, ok bool) {
	value, ok := s.Last()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// PopFirstAsNotEmpty returns element of slice (if exists).
func (s *PhoneCallClassArray) PopFirstAsNotEmpty() (v NotEmptyPhoneCall, ok bool) {
	value, ok := s.PopFirst()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// PopAsNotEmpty returns element of slice (if exists).
func (s *PhoneCallClassArray) PopAsNotEmpty() (v NotEmptyPhoneCall, ok bool) {
	value, ok := s.Pop()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// PhoneCallEmptyArray is adapter for slice of PhoneCallEmpty.
type PhoneCallEmptyArray []PhoneCallEmpty

// Sort sorts slice of PhoneCallEmpty.
func (s PhoneCallEmptyArray) Sort(less func(a, b PhoneCallEmpty) bool) PhoneCallEmptyArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of PhoneCallEmpty.
func (s PhoneCallEmptyArray) SortStable(less func(a, b PhoneCallEmpty) bool) PhoneCallEmptyArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of PhoneCallEmpty.
func (s PhoneCallEmptyArray) Retain(keep func(x PhoneCallEmpty) bool) PhoneCallEmptyArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s PhoneCallEmptyArray) First() (v PhoneCallEmpty, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s PhoneCallEmptyArray) Last() (v PhoneCallEmpty, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *PhoneCallEmptyArray) PopFirst() (v PhoneCallEmpty, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero PhoneCallEmpty
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *PhoneCallEmptyArray) Pop() (v PhoneCallEmpty, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByID sorts slice of PhoneCallEmpty by ID.
func (s PhoneCallEmptyArray) SortByID() PhoneCallEmptyArray {
	return s.Sort(func(a, b PhoneCallEmpty) bool {
		return a.GetID() < b.GetID()
	})
}

// SortStableByID sorts slice of PhoneCallEmpty by ID.
func (s PhoneCallEmptyArray) SortStableByID() PhoneCallEmptyArray {
	return s.SortStable(func(a, b PhoneCallEmpty) bool {
		return a.GetID() < b.GetID()
	})
}

// FillMap fills constructors to given map.
func (s PhoneCallEmptyArray) FillMap(to map[int64]PhoneCallEmpty) {
	for _, value := range s {
		to[value.GetID()] = value
	}
}

// ToMap collects constructors to map.
func (s PhoneCallEmptyArray) ToMap() map[int64]PhoneCallEmpty {
	r := make(map[int64]PhoneCallEmpty, len(s))
	s.FillMap(r)
	return r
}

// PhoneCallWaitingArray is adapter for slice of PhoneCallWaiting.
type PhoneCallWaitingArray []PhoneCallWaiting

// Sort sorts slice of PhoneCallWaiting.
func (s PhoneCallWaitingArray) Sort(less func(a, b PhoneCallWaiting) bool) PhoneCallWaitingArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of PhoneCallWaiting.
func (s PhoneCallWaitingArray) SortStable(less func(a, b PhoneCallWaiting) bool) PhoneCallWaitingArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of PhoneCallWaiting.
func (s PhoneCallWaitingArray) Retain(keep func(x PhoneCallWaiting) bool) PhoneCallWaitingArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s PhoneCallWaitingArray) First() (v PhoneCallWaiting, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s PhoneCallWaitingArray) Last() (v PhoneCallWaiting, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *PhoneCallWaitingArray) PopFirst() (v PhoneCallWaiting, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero PhoneCallWaiting
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *PhoneCallWaitingArray) Pop() (v PhoneCallWaiting, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByID sorts slice of PhoneCallWaiting by ID.
func (s PhoneCallWaitingArray) SortByID() PhoneCallWaitingArray {
	return s.Sort(func(a, b PhoneCallWaiting) bool {
		return a.GetID() < b.GetID()
	})
}

// SortStableByID sorts slice of PhoneCallWaiting by ID.
func (s PhoneCallWaitingArray) SortStableByID() PhoneCallWaitingArray {
	return s.SortStable(func(a, b PhoneCallWaiting) bool {
		return a.GetID() < b.GetID()
	})
}

// SortByDate sorts slice of PhoneCallWaiting by Date.
func (s PhoneCallWaitingArray) SortByDate() PhoneCallWaitingArray {
	return s.Sort(func(a, b PhoneCallWaiting) bool {
		return a.GetDate() < b.GetDate()
	})
}

// SortStableByDate sorts slice of PhoneCallWaiting by Date.
func (s PhoneCallWaitingArray) SortStableByDate() PhoneCallWaitingArray {
	return s.SortStable(func(a, b PhoneCallWaiting) bool {
		return a.GetDate() < b.GetDate()
	})
}

// FillMap fills constructors to given map.
func (s PhoneCallWaitingArray) FillMap(to map[int64]PhoneCallWaiting) {
	for _, value := range s {
		to[value.GetID()] = value
	}
}

// ToMap collects constructors to map.
func (s PhoneCallWaitingArray) ToMap() map[int64]PhoneCallWaiting {
	r := make(map[int64]PhoneCallWaiting, len(s))
	s.FillMap(r)
	return r
}

// PhoneCallRequestedArray is adapter for slice of PhoneCallRequested.
type PhoneCallRequestedArray []PhoneCallRequested

// Sort sorts slice of PhoneCallRequested.
func (s PhoneCallRequestedArray) Sort(less func(a, b PhoneCallRequested) bool) PhoneCallRequestedArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of PhoneCallRequested.
func (s PhoneCallRequestedArray) SortStable(less func(a, b PhoneCallRequested) bool) PhoneCallRequestedArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of PhoneCallRequested.
func (s PhoneCallRequestedArray) Retain(keep func(x PhoneCallRequested) bool) PhoneCallRequestedArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s PhoneCallRequestedArray) First() (v PhoneCallRequested, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s PhoneCallRequestedArray) Last() (v PhoneCallRequested, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *PhoneCallRequestedArray) PopFirst() (v PhoneCallRequested, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero PhoneCallRequested
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *PhoneCallRequestedArray) Pop() (v PhoneCallRequested, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByID sorts slice of PhoneCallRequested by ID.
func (s PhoneCallRequestedArray) SortByID() PhoneCallRequestedArray {
	return s.Sort(func(a, b PhoneCallRequested) bool {
		return a.GetID() < b.GetID()
	})
}

// SortStableByID sorts slice of PhoneCallRequested by ID.
func (s PhoneCallRequestedArray) SortStableByID() PhoneCallRequestedArray {
	return s.SortStable(func(a, b PhoneCallRequested) bool {
		return a.GetID() < b.GetID()
	})
}

// SortByDate sorts slice of PhoneCallRequested by Date.
func (s PhoneCallRequestedArray) SortByDate() PhoneCallRequestedArray {
	return s.Sort(func(a, b PhoneCallRequested) bool {
		return a.GetDate() < b.GetDate()
	})
}

// SortStableByDate sorts slice of PhoneCallRequested by Date.
func (s PhoneCallRequestedArray) SortStableByDate() PhoneCallRequestedArray {
	return s.SortStable(func(a, b PhoneCallRequested) bool {
		return a.GetDate() < b.GetDate()
	})
}

// FillMap fills constructors to given map.
func (s PhoneCallRequestedArray) FillMap(to map[int64]PhoneCallRequested) {
	for _, value := range s {
		to[value.GetID()] = value
	}
}

// ToMap collects constructors to map.
func (s PhoneCallRequestedArray) ToMap() map[int64]PhoneCallRequested {
	r := make(map[int64]PhoneCallRequested, len(s))
	s.FillMap(r)
	return r
}

// PhoneCallAcceptedArray is adapter for slice of PhoneCallAccepted.
type PhoneCallAcceptedArray []PhoneCallAccepted

// Sort sorts slice of PhoneCallAccepted.
func (s PhoneCallAcceptedArray) Sort(less func(a, b PhoneCallAccepted) bool) PhoneCallAcceptedArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of PhoneCallAccepted.
func (s PhoneCallAcceptedArray) SortStable(less func(a, b PhoneCallAccepted) bool) PhoneCallAcceptedArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of PhoneCallAccepted.
func (s PhoneCallAcceptedArray) Retain(keep func(x PhoneCallAccepted) bool) PhoneCallAcceptedArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s PhoneCallAcceptedArray) First() (v PhoneCallAccepted, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s PhoneCallAcceptedArray) Last() (v PhoneCallAccepted, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *PhoneCallAcceptedArray) PopFirst() (v PhoneCallAccepted, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero PhoneCallAccepted
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *PhoneCallAcceptedArray) Pop() (v PhoneCallAccepted, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByID sorts slice of PhoneCallAccepted by ID.
func (s PhoneCallAcceptedArray) SortByID() PhoneCallAcceptedArray {
	return s.Sort(func(a, b PhoneCallAccepted) bool {
		return a.GetID() < b.GetID()
	})
}

// SortStableByID sorts slice of PhoneCallAccepted by ID.
func (s PhoneCallAcceptedArray) SortStableByID() PhoneCallAcceptedArray {
	return s.SortStable(func(a, b PhoneCallAccepted) bool {
		return a.GetID() < b.GetID()
	})
}

// SortByDate sorts slice of PhoneCallAccepted by Date.
func (s PhoneCallAcceptedArray) SortByDate() PhoneCallAcceptedArray {
	return s.Sort(func(a, b PhoneCallAccepted) bool {
		return a.GetDate() < b.GetDate()
	})
}

// SortStableByDate sorts slice of PhoneCallAccepted by Date.
func (s PhoneCallAcceptedArray) SortStableByDate() PhoneCallAcceptedArray {
	return s.SortStable(func(a, b PhoneCallAccepted) bool {
		return a.GetDate() < b.GetDate()
	})
}

// FillMap fills constructors to given map.
func (s PhoneCallAcceptedArray) FillMap(to map[int64]PhoneCallAccepted) {
	for _, value := range s {
		to[value.GetID()] = value
	}
}

// ToMap collects constructors to map.
func (s PhoneCallAcceptedArray) ToMap() map[int64]PhoneCallAccepted {
	r := make(map[int64]PhoneCallAccepted, len(s))
	s.FillMap(r)
	return r
}

// PhoneCallArray is adapter for slice of PhoneCall.
type PhoneCallArray []PhoneCall

// Sort sorts slice of PhoneCall.
func (s PhoneCallArray) Sort(less func(a, b PhoneCall) bool) PhoneCallArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of PhoneCall.
func (s PhoneCallArray) SortStable(less func(a, b PhoneCall) bool) PhoneCallArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of PhoneCall.
func (s PhoneCallArray) Retain(keep func(x PhoneCall) bool) PhoneCallArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s PhoneCallArray) First() (v PhoneCall, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s PhoneCallArray) Last() (v PhoneCall, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *PhoneCallArray) PopFirst() (v PhoneCall, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero PhoneCall
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *PhoneCallArray) Pop() (v PhoneCall, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByID sorts slice of PhoneCall by ID.
func (s PhoneCallArray) SortByID() PhoneCallArray {
	return s.Sort(func(a, b PhoneCall) bool {
		return a.GetID() < b.GetID()
	})
}

// SortStableByID sorts slice of PhoneCall by ID.
func (s PhoneCallArray) SortStableByID() PhoneCallArray {
	return s.SortStable(func(a, b PhoneCall) bool {
		return a.GetID() < b.GetID()
	})
}

// SortByDate sorts slice of PhoneCall by Date.
func (s PhoneCallArray) SortByDate() PhoneCallArray {
	return s.Sort(func(a, b PhoneCall) bool {
		return a.GetDate() < b.GetDate()
	})
}

// SortStableByDate sorts slice of PhoneCall by Date.
func (s PhoneCallArray) SortStableByDate() PhoneCallArray {
	return s.SortStable(func(a, b PhoneCall) bool {
		return a.GetDate() < b.GetDate()
	})
}

// FillMap fills constructors to given map.
func (s PhoneCallArray) FillMap(to map[int64]PhoneCall) {
	for _, value := range s {
		to[value.GetID()] = value
	}
}

// ToMap collects constructors to map.
func (s PhoneCallArray) ToMap() map[int64]PhoneCall {
	r := make(map[int64]PhoneCall, len(s))
	s.FillMap(r)
	return r
}

// PhoneCallDiscardedArray is adapter for slice of PhoneCallDiscarded.
type PhoneCallDiscardedArray []PhoneCallDiscarded

// Sort sorts slice of PhoneCallDiscarded.
func (s PhoneCallDiscardedArray) Sort(less func(a, b PhoneCallDiscarded) bool) PhoneCallDiscardedArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of PhoneCallDiscarded.
func (s PhoneCallDiscardedArray) SortStable(less func(a, b PhoneCallDiscarded) bool) PhoneCallDiscardedArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of PhoneCallDiscarded.
func (s PhoneCallDiscardedArray) Retain(keep func(x PhoneCallDiscarded) bool) PhoneCallDiscardedArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s PhoneCallDiscardedArray) First() (v PhoneCallDiscarded, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s PhoneCallDiscardedArray) Last() (v PhoneCallDiscarded, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *PhoneCallDiscardedArray) PopFirst() (v PhoneCallDiscarded, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero PhoneCallDiscarded
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *PhoneCallDiscardedArray) Pop() (v PhoneCallDiscarded, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByID sorts slice of PhoneCallDiscarded by ID.
func (s PhoneCallDiscardedArray) SortByID() PhoneCallDiscardedArray {
	return s.Sort(func(a, b PhoneCallDiscarded) bool {
		return a.GetID() < b.GetID()
	})
}

// SortStableByID sorts slice of PhoneCallDiscarded by ID.
func (s PhoneCallDiscardedArray) SortStableByID() PhoneCallDiscardedArray {
	return s.SortStable(func(a, b PhoneCallDiscarded) bool {
		return a.GetID() < b.GetID()
	})
}

// FillMap fills constructors to given map.
func (s PhoneCallDiscardedArray) FillMap(to map[int64]PhoneCallDiscarded) {
	for _, value := range s {
		to[value.GetID()] = value
	}
}

// ToMap collects constructors to map.
func (s PhoneCallDiscardedArray) ToMap() map[int64]PhoneCallDiscarded {
	r := make(map[int64]PhoneCallDiscarded, len(s))
	s.FillMap(r)
	return r
}
