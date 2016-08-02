package bit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	// 1000 1011 1010 1101 1111 0000 0000 1101
	badFood = []byte{0x8b, 0xad, 0xf0, 0x0d}

	// 1110 1100 0111 0000 1100 0000 0000 0001 1110 0010
	ectoCooler = []byte{0xec, 0x70, 0xc0, 0x01, 0xe2}
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
	//         ^
	assert.Equal(uint64(0x01), r.ReadBits(1))

	// 1000 1011 1010 1101 1111 0000 0000 1101
	//      ^-^
	assert.Equal(uint64(0x5), r.ReadBits(3))

	// 1000 1011 1010 1101 1111 0000 0000 1101
	// ^--^
	assert.Equal(uint64(0x8), r.ReadBits(4))

	// 1000 1011 1010 1101 1111 0000 0000 1101
	//              ^----^
	assert.Equal(uint64(0xd), r.ReadBits(5))

	// 1000 1011 1010 1101 1111 0000 0000 1101
	//           ^-^.........^-----^
	// data bits: 1 1000 0101
	assert.Equal(uint64(0x185), r.ReadBits(9))

	// 1000 1011 1010 1101 1111 0000 0000 1101
	//                     ^^........^-------^
	// data bits: 00 0011 0111
	assert.Equal(uint64(0x37), r.ReadBits(10))
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
	r.ReadBits(1)
	// 1000 1011 1010 1101 1111 0000 0000 1101
	// ^------^..........^                     : 1100 0101
	//           ^------^..........^           : 0101 0110
	//                     ^------^..........^ : 1111 1000
	buf = make([]byte, 3)

	n, err = r.Read(buf)
	assert.NoError(err)
	assert.Equal(3, n)

	expected := []byte{0xc5, 0x56, 0xf8}
	assert.Equal(expected, buf)
}

func TestUbitVar(t *testing.T) {
	var (
		assert = assert.New(t)
		r      *Reader
		u      uint64
	)

	// 1000 1011 1010 1101 1111 0000 0000 1101
	r = NewBytesReader(badFood)

	// 1000 1011 1010 1101 1111 0000 0000 1101
	//         ^
	r.ReadBits(1)

	// 1000 1011 1010 1101 1111 0000 0000 1101
	//  ^^                                     : prefix 00
	//    ^---^                                : data 0101
	u = r.ReadUBitVar()
	assert.Equal(uint64(5), u)

	// 1000 1011 1010 1101 1111 0000 0000 1101
	// ^.................^
	r.ReadBits(2)

	// 1000 1011 1010 1101 1111 0000 0000 1101
	//            ^^                           : prefix 01
	//           ^...............^-^           : msb 0001
	//              ^---^                      : lsb 0110
	u = r.ReadUBitVar()
	assert.Equal(uint64(0x16), u)

	// 1000 1011 1010 1101 1111 0000 0000 1101
	r = NewBytesReader(badFood)

	// 1000 1011 1010 1101 1111 0000 0000 1101
	//        ^^
	r.ReadBits(2)

	// 1000 1011 1010 1101 1111 0000 0000 1101
	// ^^                                      : prefix 10
	//           ^-------^                     : msb 1010 1101
	//   ^---^                                 : lsb 0010
	u = r.ReadUBitVar()
	assert.Equal(uint64(0xad2), u)

	// 1110 1100 0111 0000 1100 0000 0000 0001 1110 0010
	r = NewBytesReader(ectoCooler)

	// 1110 1100 0111 0000 1100 0000 0000 0001 1110 0010
	//        ^^
	r.ReadBits(2)

	// 1110 1100 0111 0000 1100 0000 0000 0001 1110 0010
	// ^^                                                : prefix 11
	//                                              ^--^ : msb 0010
	//                               ^-------^           : 0000 0001
	//                     ^-------^                     : 1110 0000
	//           ^-------^                               : 0111 0000
	//   ^---^                                           : lsb 1011
	// data bits:
	// 0010 0000 0001 1110 0000 0111 0000 1011
	u = r.ReadUBitVar()
	assert.Equal(uint64(0x201c070b), u)
}

func TestVarInt(t *testing.T) {
	var (
		assert = assert.New(t)
		r      *Reader
		u      uint64
	)

	r = NewBytesReader(badFood)
	r.ReadBits(24)
	// 1000 1011 1010 1101 1111 0000 0000 1101
	//                                ^------^ : data
	//                               ^         : stop
	u = r.ReadVarInt()
	assert.Equal(uint64(0xd), u)

	r = NewBytesReader(badFood)
	r.ReadBits(16)
	// 1000 1011 1010 1101 1111 0000 0000 1101
	//                                ^------^ : msb
	//                               ^         : stop
	//                      ^------^           : lsb
	//                     ^                   : continue
	// data bits:
	// 0000 0110 1111 0000
	// 0    6    f    0
	u = r.ReadVarInt()
	assert.Equal(uint64(0x6f0), u)

	r = NewBytesReader(badFood)
	// 1000 1011 1010 1101 1111 0000 0000 1101
	//                                ^------^ : msb      000 1101
	//                               ^         : stop
	//                      ^------^           : data     111 0000
	//                     ^                   : continue
	//            ^------^                     : data     010 1101
	//           ^                             : continue
	//  ^------^                               : lsb      000 1011
	// ^                                       : continue
	// data bits:
	// 0001 1011 1100 0001 0110 1000 1011
	// 1    b    c    1    6    8    b
	u = r.ReadVarInt()
	assert.Equal(uint64(0x1bc168b), u)
}
