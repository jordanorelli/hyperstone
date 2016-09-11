package ent

import (
	"fmt"
	"github.com/jordanorelli/hyperstone/bit"
)

// a fieldpath is a list of integers that is used to walk the type hierarchy to
// identify a given field on a given type.
type fieldPath struct {
	// slice of values, to be reused over and over
	vals []int
	// index of the last valid value. e.g., the head of the stack.
	last int
}

func newFieldPath() *fieldPath {
	f := &fieldPath{vals: make([]int, 32)}
	f.vals[f.last] = -1
	return f
}

func (f *fieldPath) add(i int) {
	f.vals[f.last] += i
}

func (f *fieldPath) push(i int) {
	f.last++
	f.vals[f.last] = i
}

func (f *fieldPath) pop() int {
	f.last--
	return f.vals[f.last+1]
}

func (f *fieldPath) replaceAll(fn func(v int) int) {
	for i := 0; i <= f.last; i++ {
		f.vals[i] = fn(f.vals[i])
	}
}

// reads the sequence of id values off of the provided bit reader given the
// huffman tree of fieldpath ops rooted at the node n
func (f *fieldPath) read(br bit.Reader, n node) error {
	f.last = 0
	for fn := walk(n, br); fn != nil; fn = walk(n, br) {
		if err := br.Err(); err != nil {
			return fmt.Errorf("unable to read fieldpath: reader error: %v", err)
		}
		fn(f, br)
	}
	return nil
}

// the subslice of valid index values that has been read on the fieldpath
func (f *fieldPath) path() []int {
	return f.vals[:f.last+1]
}
