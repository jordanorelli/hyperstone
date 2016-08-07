package bit

import (
	"io"
)

// bit.Reader allows for bit-level reading of arbitrary source data. This is
// based on the bit reader found in the standard library's bzip2 package.
// https://golang.org/src/compress/bzip2/bit_reader.go
type StreamReader struct {
	src  io.ByteReader // source of data
	n    uint64        // bit buffer
	bits uint          // number of valid bits in n
	err  error         // stored error
}

// ReadBits reads the given number of bits and returns them in the
// least-significant part of a uint64.
func (r *StreamReader) ReadBits(bits uint) (n uint64) {
	if r.err != nil {
		return 0
	}

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

// ReadByte reads a single byte, regardless of alignment.
func (r *StreamReader) ReadByte() byte {
	if r.bits == 0 {
		b, err := r.src.ReadByte()
		if err != nil {
			r.err = err
			return 0
		}
		return b
	}
	return byte(r.ReadBits(8))
}

// Read reads like an io.Reader, taking care of alignment internally.
func (r *StreamReader) Read(buf []byte) int {
	for i := 0; i < len(buf); i++ {
		b := r.ReadByte()
		if r.err != nil {
			return 0
		}
		buf[i] = b
	}
	return len(buf)
}

// discards N byte of data on the reader or until EOF
func (r *StreamReader) DiscardBytes(n int) {
	for i := 0; i < n; i++ {
		r.ReadByte()
	}
}

func (r *StreamReader) Err() error { return r.err }
