package main

import (
	"fmt"
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

func (b *bitBuffer) readByte() (out byte) {
	if b.bits == 0 {
		if b.index >= len(b.source) {
			b.err = io.ErrUnexpectedEOF
			return
		}
		out = b.source[b.index]
		b.index += 1
		return
	}
	return byte(b.readBits(8))
}

func (b *bitBuffer) readBytes(n int) []byte {
	if b.bits == 0 {
		b.index += n
		if b.index > len(b.source) {
			b.err = io.ErrUnexpectedEOF
			return []byte{}
		}
		return b.source[b.index-n : b.index]
	}
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		buf[i] = byte(b.readBits(8))
	}
	return buf
}

// readVarUint reads a variable-length uint32, encoded with some scheme that I
// can't find a standard for. first two bits are a length prefix, followed by a
// 4, 8, 12, or 32-bit wide uint.
func (b *bitBuffer) readVarUint() uint32 {
	switch b.readBits(2) {
	case 0:
		return uint32(b.readBits(4))
	case 1:
		return uint32(b.readBits(4) | b.readBits(4)<<4)
	case 2:
		return uint32(b.readBits(4) | b.readBits(8)<<4)
	case 3:
		return uint32(b.readBits(4) | b.readBits(28)<<4)
	default:
		// this switch is already exhaustive, the compiler just can't tell.
		panic(fmt.Sprintf("invalid varuint prefix"))
	}
}

// readVarInt reads a varint-encoded value off of the front of the buffer. This
// is the varint encoding used in protobuf. That is: each byte utilizes a 7-bit
// group. the msb of each byte indicates whether there are more bytes to
// follow.
func (b *bitBuffer) readVarInt() uint64 {
	var x, n uint64
	for shift := uint(0); shift < 64; shift += 7 {
		n = b.readBits(8)
		if n < 0x80 {
			return x | n<<shift
		}
		x |= n &^ 0x80 << shift
	}
	b.err = fmt.Errorf("readVarInt never saw the end of varint")
	return 0
}
