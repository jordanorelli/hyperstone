package bit

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
