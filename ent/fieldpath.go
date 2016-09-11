package ent

import (
	"bytes"
	"fmt"
)

// a fieldpath is a list of integers that is used to walk the type hierarchy to
// identify a given field on a given type.
type fieldPath struct {
	vals []int
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

func (f *fieldPath) pathString() string {
	var buf bytes.Buffer
	for i := 0; i <= f.last; i++ {
		fmt.Fprintf(&buf, "/%d", f.vals[i])
	}
	return buf.String()
}
