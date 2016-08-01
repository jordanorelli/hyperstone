package main

import (
	"io"
)

// a bitBuffer is a buffer of raw data that is not necessarily byte-aligned,
// for performing bitwise reads and manipulation. The bulk of this source is
// adapted from the bitReader defined in the standard library bzip2 package.
type bitBuffer struct {
	source  []byte // the source data to be read. This slice is never modified
	index   int    // position of the last byte read out of source
	scratch uint64 // scratch register of bits worthy of manipulation
	bits    uint   // bit width of the scratch register
	err     error  // stored error
}

func newBitBuffer(buf []byte) *bitBuffer {
	return &bitBuffer{source: buf}
}

func (b *bitBuffer) readBits(bits uint) uint64 {
	for bits > b.bits {
		if b.index >= len(b.source) {
			b.err = io.ErrUnexpectedEOF
			return 0
		}
		b.scratch <<= 8
		b.scratch |= uint64(b.source[b.index])
		b.index += 1
		b.bits += 8
	}

	// b.scratch looks like this (assuming that b.bits = 14 and bits = 6):
	// Bit: 111111
	//      5432109876543210
	//
	//         (6 bits, the desired output)
	//        |-----|
	//        V     V
	//      0101101101001110
	//        ^            ^
	//        |------------|
	//           b.bits (num valid bits)
	//
	// This the next line right shifts the desired bits into the
	// least-significant places and masks off anything above.
	n := (b.scratch >> (b.bits - bits)) & ((1 << bits) - 1)
	b.bits -= bits
	return n
}
