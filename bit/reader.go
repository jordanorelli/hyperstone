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
		r.n |= uint64(b) << r.bits
		r.bits += 8
	}
	n = r.n & (1<<bits - 1)
	r.n >>= bits
	r.bits -= bits
	return
}

// discards up to bits bits. returns a bool indicating wheter any errors occured.
func (r *Reader) DiscardBits(n int) {
	r.ReadBits(uint(n))
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

// discards N byte of data on the reader or until EOF
func (r *Reader) DiscardBytes(n int) {
	for i := 0; i < n; i++ {
		_, err := r.ReadByte()
		if err != nil {
			r.err = err
			return
		}
	}
}

func (r *Reader) Err() error { return r.err }

// ReadUbitVar reads a prefixed uint value. A prefix is 2 bits wide, followed
// by the 4 least-significant bits, then a variable number of most-significant
// bits based on the prefix.
//
// 00 - 4
// 01 - 8
// 10 - 12 (why 12? this really baffles me)
// 11 - 32
func ReadUBitVar(r *Reader) uint64 {
	switch b := r.ReadBits(6); b >> 4 {
	case 0:
		return b & 0xf
	case 1:
		return b&0xf | r.ReadBits(4)<<4
	case 2:
		return b&0xf | r.ReadBits(8)<<4
	case 3:
		return b&0xf | r.ReadBits(28)<<4
	default:
		panic("not reached")
	}
}

// ReadVarInt reads a variable length int value as a uint64. This is the binary
// representation used by Protobuf. Each byte contributes 7 bits to the value
// in little-endian order. The most-significant bit of each byte represents a
// continuation bit.
func ReadVarInt(r *Reader) uint64 {
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
