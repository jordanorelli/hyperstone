package bit

import (
	"bufio"
	"bytes"
	"io"
)

// bit.Reader allows for bit-level reading of arbitrary source data. This is
// based on the bit reader found in the standard library's bzip2 package.
// https://golang.org/src/compress/bzip2/bit_reader.go
type Reader struct {
	src  io.ByteReader // source of data
	n    uint64        // bit buffer
	bits uint          // number of valid bits in n
	err  error         // stored error
}

// NewReader creates a new bit.Reader for any arbitrary reader.
func NewReader(r io.Reader) *Reader {
	br, ok := r.(io.ByteReader)
	if !ok {
		br = bufio.NewReader(r)
	}
	return &Reader{src: br}
}

// NewByteReader creates a bit.Reader for a static slice of bytes. It's just
// using a bytes.Reader internally.
func NewBytesReader(b []byte) *Reader {
	return NewReader(bytes.NewReader(b))
}

// ReadBits reads the given number of bits and returns them in the
// least-significant part of a uint64.
func (r *Reader) ReadBits(bits uint) (n uint64) {
	for bits > r.bits {
		b, err := r.src.ReadByte()
		if err != nil {
			r.err = err
			return 0
		}
		r.n <<= 8
		r.n |= uint64(b)
		r.bits += 8
	}
	n = (r.n >> (r.bits - bits)) & ((1 << bits) - 1)
	r.bits -= bits
	return
}

// ReadByte reads a single byte, regardless of alignment.
func (r *Reader) ReadByte() (byte, error) {
	if r.bits == 0 {
		return r.src.ReadByte()
	}
	b := byte(r.ReadBits(8))
	if err := r.Err(); err != nil {
		return 0, err
	}
	return b, nil
}

// Read reads like an io.Reader, taking care of alignment internally.
func (r *Reader) Read(buf []byte) (int, error) {
	for i := 0; i < len(buf); i++ {
		b, err := r.ReadByte()
		if err != nil {
			return 0, err
		}
		buf[i] = b
	}
	return len(buf), nil
}

// ReadUbitVar reads a prefixed uint value. A prefix is 2 bits wide, followed
// by the 4 least-significant bits, then a variable number of most-significant
// bits based on the prefix.
//
// 00 - 0
// 01 - 4
// 10 - 8
// 11 - 28
func (r *Reader) ReadUBitVar() uint64 {
	switch prefix := r.ReadBits(2); prefix {
	case 0:
		return r.ReadBits(4)
	case 1:
		return r.ReadBits(4) | r.ReadBits(4)<<4
	case 2:
		return r.ReadBits(4) | r.ReadBits(8)<<4
	case 3:
		return r.ReadBits(4) | r.ReadBits(28)<<4
	default:
		panic("not reached")
	}
}

// ReadVarInt reads a variable length int value as a uint64. This is the binary
// representation used by Protobuf. Each byte contributes 7 bits to the value
// in little-endian order. The most-significant bit of each byte represents a
// continuation bit.
func (r *Reader) ReadVarInt() uint64 {
	var (
		x     uint64
		b     uint64
		shift uint
	)
	for ; shift < 64; shift += 7 {
		b = r.ReadBits(8)
		if r.Err() != nil {
			return 0
		}
		x |= b & 0x7f << shift
		if b&0x80 == 0 {
			return x
		}
	}
	return x
}

func (r *Reader) Err() error { return r.err }
