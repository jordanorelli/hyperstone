package bit

import (
	"io"
)

type BufReader struct {
	src  []byte // source of data
	n    uint64 // bit buffer
	bits uint   // number of valid bits in n
	err  error  // stored error
}

func (r *BufReader) ReadBits(bits uint) (n uint64) {
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

func (r *BufReader) ReadByte() byte {
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

func (r *BufReader) Read(buf []byte) int {
	if r.bits == 0 {
		if len(r.src) < len(buf) {
			r.err = io.EOF
			return 0
		}
		copy(buf, r.src[:len(buf)])
		r.src = r.src[len(buf):]
		return len(buf)
	}

	for i := 0; i < len(buf); i++ {
		b := r.ReadByte()
		if r.err != nil {
			return 0
		}
		buf[i] = b
	}
	return len(buf)
}

func (r *BufReader) DiscardBytes(n int) {
	if r.bits == 0 {
		if len(r.src) < n {
			r.err = io.EOF
			return
		}
		r.src = r.src[n:]
		return
	}

	for r.bits > 8 {
		r.bits -= 8
		n -= 1
	}

	if len(r.src) < n {
		r.err = io.EOF
		return
	}

	r.n = uint64(r.src[n-1]) >> (8 - r.bits)
	r.src = r.src[n:]
}

func (r *BufReader) SetSource(b []byte) {
	r.src = b
	r.bits = 0
	r.err = nil
	r.n = 0
}

func (r *BufReader) Err() error { return r.err }
