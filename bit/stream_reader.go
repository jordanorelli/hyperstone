package bit

import (
	"io"
)

// bit.Reader allows for bit-level reading of arbitrary source data. This is
// based on the bit reader found in the standard library's bzip2 package.
// https://golang.org/src/compress/bzip2/bit_reader.go
type streamReader struct {
	src  io.ByteReader // source of data
	n    uint64        // bit buffer
	bits uint          // number of valid bits in n
	err  error         // stored error
}

// ReadBits reads the given number of bits and returns them in the
// least-significant part of a uint64.
func (r *streamReader) ReadBits(bits uint) (n uint64) {
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
func (r *streamReader) DiscardBits(n int) {
	r.ReadBits(uint(n))
}

// ReadByte reads a single byte, regardless of alignment.
func (r *streamReader) ReadByte() (byte, error) {
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
func (r *streamReader) Read(buf []byte) (int, error) {
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
func (r *streamReader) DiscardBytes(n int) {
	for i := 0; i < n; i++ {
		_, err := r.ReadByte()
		if err != nil {
			r.err = err
			return
		}
	}
}

func (r *streamReader) Err() error { return r.err }