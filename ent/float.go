package ent

import (
	"fmt"
	"math"

	"github.com/jordanorelli/hyperstone/bit"
	"github.com/jordanorelli/hyperstone/dota"
)

const (
	f_min = 1 << iota
	f_max
	f_center
)

func parseFloatType(n *Namespace, flat *dota.ProtoFlattenedSerializerFieldT) (Type, error) {
	if flat.VarEncoderSym != nil {
		encoder := n.Symbol(int(flat.GetVarEncoderSym())).String()
		switch encoder {
		case "coord":
			return nil, fmt.Errorf("coord encoder isn't dont yet")
		default:
			return nil, fmt.Errorf("unknown float encoder: %s", encoder)
		}
	}

	if flat.BitCount != nil {
		bits := int(flat.GetBitCount())
		switch {
		case bits < 0:
			return nil, fmt.Errorf("invalid bit count on float field: %d", bits)
		case bits < 32:
			return quantizedFloat(n, flat)
		case bits == 0, bits == 32:
			// these seem meaningless, which is suspicious.
		default:
			return nil, fmt.Errorf("bit count is too high on float field: %d", bits)
		}
	}

	if flat.LowValue != nil || flat.HighValue != nil {
		return nil, fmt.Errorf("float32 with a low or high value isn't supported")
	}

	type_name := n.Symbol(int(flat.GetVarTypeSym())).String()
	return &Primitive{name: type_name, read: readFloat32}, nil
}

func quantizedFloat(n *Namespace, flat *dota.ProtoFlattenedSerializerFieldT) (Type, error) {
	if flat.LowValue == nil && flat.HighValue == nil {
		return nil, fmt.Errorf("quantizedFloat has no boundaries")
	}

	bits := uint(flat.GetBitCount())
	low, high := flat.GetLowValue(), flat.GetHighValue()
	flags := int(flat.GetEncodeFlags())

	flags = flags & 7                   // dunno how to handle -8 lol
	steps := uint(1<<bits - 1)          // total number of intervals
	span := high - low                  // total range of values
	step_width := span / float32(steps) // output width of each step
	if span < 0 {
		return nil, fmt.Errorf("invalid quantization span")
	}

	var special *float32
	switch {
	case flags&f_min > 0:
		special = &low
	case flags&f_max > 0:
		special = &high
	case flags&f_center > 0:
		middle := (high + low) * 0.5
		special = &middle
	}

	read := func(br bit.Reader, d *Dict) (interface{}, error) {
		if special != nil && bit.ReadBool(br) {
			return *special, nil
		}
		u := br.ReadBits(bits)
		return low + float32(u)*step_width, br.Err()
	}

	type_name := n.Symbol(int(flat.GetVarTypeSym())).String()
	return &Primitive{name: type_name, read: read}, nil
}

// reads an IEEE 754 binary float value off of the stream
func readFloat32(br bit.Reader, d *Dict) (interface{}, error) {
	return math.Float32frombits(uint32(br.ReadBits(32))), nil
}
