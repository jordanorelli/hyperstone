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

func (r *Reader) Err() error { return r.err }
