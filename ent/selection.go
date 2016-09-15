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

func (s selection) fill(offset int, displayPath string, dest slotted, br bit.Reader) error {
	slot := s.vals[offset]
	if s.count-offset <= 0 {
		panic("selection makes no sense")
	}
	switch s.count - offset {
	case 1:
		fn := dest.slotDecoder(slot)
		if fn == nil {
			switch dest.(type) {
			case *Entity:
				Debug.Printf("%s %s (%s)", s, fmt.Sprintf("%s.%s", displayPath, dest.slotName(slot)), dest.slotType(slot))
				return nil
				// Info.Fatalf("%v entity has no decoder for slot %d (%v)", v.Class, slot, v.Class.Fields[slot])
			default:
				Info.Printf("slotted value %v has no decoder for slot %d", dest, slot)
				return nil
			}
		}
		old := dest.slotValue(slot)
		val := fn(br)
		dest.setSlotValue(slot, val)
		Debug.Printf("%s %s (%s): %v -> %v", s, fmt.Sprintf("%s.%s", displayPath, dest.slotName(slot)), dest.slotType(slot), old, val)
		return nil
	default:
		v := dest.slotValue(slot)
		vs, ok := v.(slotted)
		if !ok {
			return fmt.Errorf("selection %s at offset %d (%d) refers to a slot (%s) that contains a non-slotted type (%s) with value %v", s, offset, slot, fmt.Sprintf("%s.%s", displayPath, dest.slotName(slot)), dest.slotType(slot), v)
		}
		return s.fill(offset+1, fmt.Sprintf("%s.%s", displayPath, dest.slotName(slot)), vs, br)
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

func (r *selectionReader) readSelections(br bit.Reader, n node) ([]selection, error) {
	r.cur.count = 1
	r.cur.vals[0] = -1
	r.count = 0
	for fn := walk(n, br); fn != nil; fn = walk(n, br) {
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

// pops the last value off of the current selection
func (r *selectionReader) pop() { r.cur.count-- }

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
