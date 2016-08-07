package bit

import (
	"bufio"
	"io"
)

// Reader allows for the reading and skipping of bits and bytes.
type Reader interface {
	ReadBits(uint) uint64
	ReadByte() byte
	Read([]byte) int
	DiscardBytes(int)
	Err() error
}

// NewReader creates a new bit.Reader for any arbitrary reader.
func NewReader(r io.Reader) *StreamReader {
	br, ok := r.(io.ByteReader)
	if !ok {
		br = bufio.NewReader(r)
	}
	return &StreamReader{src: br}
}

// NewByteReader creates a bit.Reader for a static slice of bytes. It's just
// using a bytes.Reader internally.
func NewBytesReader(b []byte) *BufReader {
	return &BufReader{src: b}
}
