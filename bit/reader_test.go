package bit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	// 1000 1011 1010 1101 1111 0000 0000 1101
	badFood = []byte{0x8b, 0xad, 0xf0, 0x0d}
)

func TestRead(t *testing.T) {
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
