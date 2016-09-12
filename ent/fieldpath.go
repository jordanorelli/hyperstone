package ent

import (
	"fmt"
	"github.com/jordanorelli/hyperstone/bit"
)

type fieldPath struct {
	// slice of values, to be reused over and over
	vals [6]int

	// index of the last valid value. e.g., the head of the stack.
	last int

	// history of fieldpath values. that is, all selections thus far.
	// Everything is fixed width to avoid allocations as much as possible
	// inside of the fieldpath stuff. history lengths are known to go over 500,
	// but I honestly have no idea if there is or isn't a limit. this is a
	// hazard.
	history [1024][7]int

	// last valid index of a history entry
	hlast int
}

func newFieldPath() *fieldPath {
	f := new(fieldPath)
	f.vals[f.last] = -1
	f.hlast = -1
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
func (f *fieldPath) read(br bit.Reader, n node, class *Class) error {
	f.last = 0
	for fn := walk(n, br); fn != nil; fn = walk(n, br) {
		if err := br.Err(); err != nil {
			return fmt.Errorf("unable to read fieldpath: reader error: %v", err)
		}
		fn(f, br)
		Debug.Printf("fieldpath: %v", f.path())
		f.keep()
	}
	return nil
}

// adds our current vals to the history list.
// a history list entry is an array of ints, 7 ints wide. a fieldPath value is
// 6 ints wide. a history value encodes the number of valid entries in its
// first position. e.g., this:
// [0 0 0 0 0 0 0]
// would indicate the selection [0]
// [0 4 0 0 0 0 0]
// would indicate the selection [4]
// [1 4 2 0 0 0 0]
// would indicate the selection [4 2]
func (f *fieldPath) keep() {
	f.hlast++
	f.history[f.hlast][0] = f.last
	for i := 0; i <= f.last; i++ {
		f.history[f.hlast][i+1] = f.vals[i]
	}
}

func (f *fieldPath) getField(class *Class) *Field {
	if f.last > 0 {
		for i := 0; i < f.last; i++ {
			if f.vals[i] >= len(class.Fields) {
				Info.Fatalf("bad access for field %d on class %v; class has only %d fields", f.vals[i], class, len(class.Fields))
			}
			field := class.Fields[f.vals[i]]
			if field.class == nil {
				Info.Fatalf("class %s field at %d is %v, has no class", class, f.vals[i], field)
			} else {
				class = class.Fields[f.vals[i]].class
			}
		}
	}
	return class.Fields[f.vals[f.last]]
}

// the subslice of valid index values that has been read on the fieldpath
func (f *fieldPath) path() []int {
	return f.vals[:f.last+1]
}
