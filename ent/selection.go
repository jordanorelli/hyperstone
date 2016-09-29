package ent

import (
	"fmt"
	"github.com/jordanorelli/hyperstone/bit"
)

type selectionOp func(*selectionReader, bit.Reader)

// selection represents the selection of an individual entity slot. a count of
// 1 indicates that the selection is to be made directly on an entity. counts
// higher than 1 indiciate that the selection is to be made on either a member
// entity or an element of a member array.
type selection struct {
	count int
	vals  [6]int
}

func (s selection) String() string { return fmt.Sprint(s.path()) }
func (s selection) path() []int    { return s.vals[:s.count] }

func (s selection) fillSlots(v slotted, r bit.Reader) error {
	return s.fillSlotsIter(0, v, v.tÿpe().typeName(), r)
}

func (s selection) fillSlotsIter(offset int, dest slotted, path string, r bit.Reader) error {
	slot := s.vals[offset]
	if s.count-offset <= 0 {
		return fmt.Errorf("unable to fill selection %v having count %d at offset %d", s, s.count, offset)
	}
	switch s.count - offset {
	case 1:
		v := dest.slotType(slot).nü()
		if err := v.read(r); err != nil {
			return fmt.Errorf("unable to fill selection: %v", err)
		}
		old := dest.getSlotValue(slot)
		dest.setSlotValue(slot, v)
		Debug.Printf("%v %s.%s (%s) %v -> %v", s, path, dest.slotName(slot), dest.slotType(slot).typeName(), old, v)
		return nil
	default:
		v := dest.getSlotValue(slot)
		if v == nil {
			v = dest.slotType(slot).nü()
			dest.setSlotValue(slot, v)
		}
		vs, ok := v.(slotted)
		if !ok {
			return fmt.Errorf("dest %s (%s) isn't slotted", dest.slotName(slot), dest.slotType(slot).typeName())
		}
		return s.fillSlotsIter(offset+1, vs, fmt.Sprintf("%s.%s", path, dest.slotName(slot)), r)
	}
}

// selectionReader reads a set of field selections off of the wire. the
// selections are represented as arrays of slot positions to be traversed in
// order to select an entity slot.
type selectionReader struct {
	// current selection; the selection being read off the wire right now
	cur selection

	// the number of valid selections made thus far
	count int

	// a list of past selections. values up to the index count-1 are considered
	// valid.
	all [1024]selection
}

func (r *selectionReader) readSelections(br bit.Reader) ([]selection, error) {
	r.cur.count = 1
	r.cur.vals[0] = -1
	r.count = 0
	for fn := walk(huffRoot, br); fn != nil; fn = walk(huffRoot, br) {
		if err := br.Err(); err != nil {
			return nil, fmt.Errorf("unable to read selection: bit reader error: %v", err)
		}
		fn(r, br)
		r.keep()
	}
	if err := br.Err(); err != nil {
		return nil, fmt.Errorf("unable to read selection: bit reader error: %v", err)
	}
	return r.all[:r.count], nil
}

// adds i to the current selection
func (r *selectionReader) add(i int) {
	r.cur.vals[r.cur.count-1] += i
}

// pushes the value i to the end of the current selection
func (r *selectionReader) push(i int) {
	r.cur.vals[r.cur.count] = i
	r.cur.count++
}

// pops n elements from the selection. if n is negative, pops all but n
// elements from the selection.
func (r *selectionReader) pop(n int) {
	if n < 0 {
		r.cur.count = -n
	} else {
		r.cur.count -= n
	}
}

// maps a function over the current set of values in the current selection
func (r *selectionReader) måp(fn func(int) int) {
	for i := 0; i < r.cur.count; i++ {
		r.cur.vals[i] = fn(r.cur.vals[i])
	}
}

// keep the current selection and move on to the next one
func (r *selectionReader) keep() {
	r.all[r.count] = r.cur
	r.count++
}
