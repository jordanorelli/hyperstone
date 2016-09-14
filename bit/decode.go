package bit

import (
	"bytes"
)

// normalized values are represented with 11 significant bits. we pre-compute a
// divisor so that we can use a multiply instruction and avoid using
// floating-point division during the lifecycle of the program.
const normal_divisor = float32(1.0) / float32(2047)

const (
	coord_ibits = 14 // number of integer bits in a coord value
	coord_fbits = 5  // number of fractional bits in a coord value
	coord_res   = 1.0 / 1 << coord_fbits
)

// ReadUbitVar reads a prefixed uint value. A prefix is 2 bits wide, followed
// by the 4 least-significant bits, then a variable number of most-significant
// bits based on the prefix.
//
// 00 - 4
// 01 - 8
// 10 - 12 (why 12? this really baffles me)
// 11 - 32
func ReadUBitVar(r Reader) uint64 {
	switch b := r.ReadBits(6); b >> 4 {
	case 0:
		return b & 0xf
	case 1:
		return b&0xf | r.ReadBits(4)<<4
	case 2:
		return b&0xf | r.ReadBits(8)<<4
	case 3:
		return b&0xf | r.ReadBits(28)<<4
	default:
		panic("not reached")
	}
}

// reads some sort of uint in a variable length encoding that appears in
// fieldpath. this encoding is deeply baffling.
func ReadUBitVarFP(r Reader) uint64 {
	if ReadBool(r) {
		return r.ReadBits(2)
	}
	if ReadBool(r) {
		return r.ReadBits(4)
	}
	if ReadBool(r) {
		return r.ReadBits(10)
	}
	if ReadBool(r) {
		return r.ReadBits(17)
	}
	return r.ReadBits(31)
}

// ReadVarInt reads a variable length int value as a uint64. This is the binary
// representation used by Protobuf. Each byte contributes 7 bits to the value
// in little-endian order. The most-significant bit of each byte represents a
// continuation bit.
func ReadVarInt(r Reader) uint64 {
	var (
		x     uint64
		b     uint64
		shift uint
	)
	for ; shift < 64; shift += 7 {
		b = r.ReadBits(8)
		if r.Err() != nil {
			return 0
		}
		x |= b & 0x7f << shift
		if b&0x80 == 0 {
			return x
		}
	}
	return x
}

// reads a 32bit varint
func ReadVarInt32(r Reader) uint32 {
	var (
		x     uint64
		b     uint64
		shift uint
	)
	for ; shift < 32; shift += 7 {
		b = r.ReadBits(8)
		if r.Err() != nil {
			return 0
		}
		x |= b & 0x7f << shift
		if b&0x80 == 0 {
			return uint32(x)
		}
	}
	return uint32(x)
}

func ReadBool(r Reader) bool {
	return r.ReadBits(1) != 0
}

// reads a null-terminated string
func ReadString(r Reader) string {
	var buf bytes.Buffer
	for b := r.ReadByte(); b != 0; b = r.ReadByte() {
		buf.WriteByte(b)
	}
	return buf.String()
}

func ReadZigZag(r Reader) int64 {
	u := ReadVarInt(r)
	if u&1 > 0 {
		return ^int64(u >> 1)
	}
	return int64(u >> 1)
}

// reads a ZigZag-encoded 32 bit signed integer
func ReadZigZag32(r Reader) int32 {
	u := ReadVarInt32(r)
	if u&1 > 0 {
		return ^int32(u >> 1)
	}
	return int32(u >> 1)
}

// reads a 12-bit normalized float. The first bit represents a sign bit, the
// next sequence of 11 bits represents some normalized value between 0 and
// 2047, allowing up to 4096 positions between -1.0 and 1.0. the resulting
// float will always be between -1.0 and 1.0 (that's why it's normal)
func ReadNormal(r Reader) float32 {
	// sign bit
	if ReadBool(r) {
		return float32(r.ReadBits(11)) * normal_divisor
	} else {
		return -float32(r.ReadBits(11)) * normal_divisor
	}
}

// an angle is just a quantized float between 0 and 360
func ReadAngle(r Reader, bits uint) float32 {
	return float32(r.ReadBits(bits)) * 360.0 / float32(uint(1)<<bits-1)
}

func ReadCoord(r Reader) float32 {
	i := ReadBool(r)
	f := ReadBool(r)
	if !(i || f) {
		return 0
	}

	neg := ReadBool(r)
	var v float32
	if i {
		v = float32(r.ReadBits(coord_ibits) + 1)
	}
	if f {
		v += float32(r.ReadBits(coord_fbits)) * coord_res
	}
	if neg {
		return -v
	}
	return v
}
