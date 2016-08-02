package bit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	// 1000 1011 1010 1101 1111 0000 0000 1101
	badFood = []byte{0x8b, 0xad, 0xf0, 0x0d}
)

// test the bit-level reads
func TestReadBits(t *testing.T) {
	assert := assert.New(t)

	var r *Reader

	// aligned reading
	r = NewBytesReader(badFood)
	assert.Equal(uint64(0x8b), r.ReadBits(8))
	assert.Equal(uint64(0xad), r.ReadBits(8))
	assert.Equal(uint64(0xf0), r.ReadBits(8))
	assert.Equal(uint64(0x0d), r.ReadBits(8))

	// misaligned reading
	r = NewBytesReader(badFood)

	// 1000 1011 1010 1101 1111 0000 0000 1101
	// ^
	assert.Equal(uint64(0x01), r.ReadBits(1))

	// 1000 1011 1010 1101 1111 0000 0000 1101
	//  ^-^
	assert.Equal(uint64(0), r.ReadBits(3))

	// 1000 1011 1010 1101 1111 0000 0000 1101
	//      ^--^
	assert.Equal(uint64(0xb), r.ReadBits(4))

	// 1000 1011 1010 1101 1111 0000 0000 1101
	//           ^----^
	assert.Equal(uint64(0x15), r.ReadBits(5))

	// 1000 1011 1010 1101 1111 0000 0000 1101
	//                 ^---------^
	assert.Equal(uint64(0x17c), r.ReadBits(9))

	// 1000 1011 1010 1101 1111 0000 0000 1101
	//                            ^----------^
	assert.Equal(uint64(0xd), r.ReadBits(10))
}

// test the Read calls, satisfying io.Reader
func TestRead(t *testing.T) {
	assert := assert.New(t)
	var r *Reader

	r = NewBytesReader(badFood)
	// 1000 1011 1010 1101 1111 0000 0000 1101
	r.ReadBits(1)
	// 0001 0111 0101 1011 1110 0000 0001 101
	buf := make([]byte, 3)

	n, err := r.Read(buf)
	assert.NoError(err)
	assert.Equal(3, n)

	expected := []byte{0x17, 0x5b, 0xe0}
	assert.Equal(expected, buf)
}
