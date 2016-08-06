package bit

import (
	"bufio"
	"bytes"
	"io"
)

type Reader interface {
	ReadBits(uint) uint64
	DiscardBits(int)
	ReadByte() (byte, error)
	Read([]byte) (int, error)
	DiscardBytes(int)
	Err() error
}

// NewReader creates a new bit.Reader for any arbitrary reader.
func NewReader(r io.Reader) Reader {
	br, ok := r.(io.ByteReader)
	if !ok {
		br = bufio.NewReader(r)
	}
	return &streamReader{src: br}
}

// NewByteReader creates a bit.Reader for a static slice of bytes. It's just
// using a bytes.Reader internally.
func NewBytesReader(b []byte) Reader {
	return NewReader(bytes.NewReader(b))
}
