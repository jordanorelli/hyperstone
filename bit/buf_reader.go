package bit

import (
	"io"
)

type bufReader struct {
	src  []byte // source of data
	n    uint64 // bit buffer
	bits uint   // number of valid bits in n
	err  error  // stored error
}

func (r *bufReader) ReadBits(bits uint) (n uint64) {
	for bits > r.bits {
		if len(r.src) == 0 {
			r.err = io.EOF
			return 0
		}
		r.n |= uint64(r.src[0]) << r.bits
		r.src = r.src[1:]
		r.bits += 8
	}
	n = r.n & (1<<bits - 1)
	r.n >>= bits
	r.bits -= bits
	return
}

func (r *bufReader) ReadByte() byte {
	if r.bits == 0 {
		if len(r.src) == 0 {
			r.err = io.EOF
			return 0
		}
		b := r.src[0]
		r.src = r.src[1:]
		return b
	}
	return byte(r.ReadBits(8))
}

func (r *bufReader) Read(buf []byte) int {
	for i := 0; i < len(buf); i++ {
		b := r.ReadByte()
		if r.err != nil {
			return 0
		}
		buf[i] = b
	}
	return len(buf)
}

func (r *bufReader) DiscardBytes(n int) {
	for i := 0; i < n; i++ {
		r.ReadByte()
	}
}

func (r *bufReader) Err() error { return r.err }
