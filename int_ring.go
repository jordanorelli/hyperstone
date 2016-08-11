package main

type intRing struct {
	items []int
	next  int  // index of the next available slot
	full  bool // whether or not we've filled the ring
	base  int  // index in items of the virtual zero index
}

func newIntRing(size int) *intRing {
	return &intRing{
		items: make([]int, size),
	}
}

func (r *intRing) add(i int) {
	r.items[r.next] = i
	r.next = r.incr(r.next)
	if !r.full && r.next == 0 {
		r.full = true
	}
	if r.full {
		r.base = r.incr(r.base)
	}
}

func (r *intRing) incr(i int) int {
	if i == len(r.items)-1 {
		return 0
	}
	return i + 1
}

func (r *intRing) at(i int) int {
	idx := r.base + i
	for idx >= len(r.items) {
		idx -= len(r.items)
	}
	return r.items[idx]
}

func (r *intRing) clear() {
	r.next = 0
	r.full = false
	r.base = 0
}
