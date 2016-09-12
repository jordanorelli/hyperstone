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

func (s selection) path() []int { return s.vals[:s.count] }

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

func newSelectionReader() *selectionReader {
	r := new(selectionReader)
	r.cur.count = 1
	r.cur.vals[0] = -1
	return r
}

func (r *selectionReader) read(br bit.Reader, n node) error {
	for fn := walk(n, br); fn != nil; fn = walk(n, br) {
		if err := br.Err(); err != nil {
			return fmt.Errorf("unable to read selection: bit reader error: %v", err)
		}
		fn(r, br)
		Debug.Printf("selection: %v", r.cur.path())
		r.keep()
	}
	return nil
}

func (r *selectionReader) selections() []selection { return r.all[:r.count] }

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
