package bit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	// 1000 1011 1010 1101 1111 0000 0000 1101
	badFood = []byte{0x8b, 0xad, 0xf0, 0x0d}

	// 0000 1110 0001 1110 1110 0111 1011 1110 1110 1111
	eLeetBeef = []byte{0x0e, 0x1e, 0xe7, 0xbe, 0xef}
)

// test the bit-level reads
func TestReadBits(t *testing.T) {
	var (
		assert = assert.New(t)
		r      *Reader
	)

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
	//                 ^----+----^
	assert.Equal(uint64(0x17c), r.ReadBits(9))

	// 1000 1011 1010 1101 1111 0000 0000 1101
	//                            ^----+----+^
	assert.Equal(uint64(0xd), r.ReadBits(10))
}

// test the Read calls, satisfying io.Reader
func TestRead(t *testing.T) {
	var (
		assert = assert.New(t)
		r      *Reader
		buf    []byte
		n      int
		err    error
	)

	r = NewBytesReader(badFood)
	buf = make([]byte, 4)
	n, err = r.Read(buf)
	assert.NoError(err)
	assert.Equal(4, n)
	assert.Equal(badFood, buf)

	r = NewBytesReader(badFood)
	// 1000 1011 1010 1101 1111 0000 0000 1101
	r.ReadBits(1)
	// 0001 0111 0101 1011 1110 0000 0001 101
	buf = make([]byte, 3)

	n, err = r.Read(buf)
	assert.NoError(err)
	assert.Equal(3, n)

	expected := []byte{0x17, 0x5b, 0xe0}
	assert.Equal(expected, buf)
}

func TestUbitVar(t *testing.T) {
	var (
		assert = assert.New(t)
		r      *Reader
		u      uint64
	)

	// 1000 1011 1010 1101 1111 0000 0000 1101
	//  ||^---^                                : data (0101)
	//  ^^                                     : prefix (00)
	r = NewBytesReader(badFood)
	r.ReadBits(1)
	u = r.ReadUBitVar()
	assert.Equal(uint64(5), u)

	r.ReadBits(2)
	// 1000 1011 1010 1101 1111 0000 0000 1101
	//            |||   |  ^--^                : most significant (1111)
	//            ||^---^                      : least significant (0110)
	//            ^^                           : prefix (01)
	u = r.ReadUBitVar()
	assert.Equal(uint64(0xf6), u)

	r = NewBytesReader(badFood)
	// 1000 1011 1010 1101 1111 0000 0000 1101
	// |||   |^----+---^                       : most significant (1110 1011)
	// ||^---^                                 : least significant (0010)
	// ^^                                      : prefix (10)
	u = r.ReadUBitVar()
	assert.Equal(uint64(0xeb2), u)

	r = NewBytesReader(eLeetBeef)
	r.ReadBits(4)
	// 0000 1110 0001 1110 1110 0111 1011 1110 1110 1111
	//      |||   |^----+----+----+----+----+----+---^   : msb (0111 1011 1001 1110
	//      |||   |                                             1111 1011 1011)
	//      ||^---^                                      : lsb (1000)
	//      ^^                                           : prefix
	u = r.ReadUBitVar()
	assert.Equal(uint64(0x7b9efbb8), u)
}
