package ent

import (
	"fmt"
	"math"

	"github.com/jordanorelli/hyperstone/bit"
	"github.com/jordanorelli/hyperstone/dota"
)

func parseQAngleType(n *Namespace, flat *dota.ProtoFlattenedSerializerFieldT) (Type, error) {
	type_name := n.Symbol(int(flat.GetVarTypeSym())).String()
	if flat.VarEncoderSym != nil {
		encoder := n.Symbol(int(flat.GetVarEncoderSym())).String()
		switch encoder {
		case "qangle_pitch_yaw":
			return nil, fmt.Errorf("that qangle pitch yaw thing isn't done yet")
		default:
			return nil, fmt.Errorf("unknown qangle encoder: %s", encoder)
		}
	}
	if flat.BitCount == nil {
		return nil, fmt.Errorf("dunno what to do when qangle type has no bitcount")
	}
	if flat.GetBitCount() < 0 {
		return nil, fmt.Errorf("negative bit count wtf")
	}
	bits := uint(flat.GetBitCount())
	switch bits {
	case 0:
		return &Primitive{name: type_name, read: readQAngleCoords}, nil
	case 32:
		return &Primitive{name: type_name, read: readQAngleFloats}, nil
	default:
		return &Primitive{name: type_name, read: qangleReader(bits)}, nil
	}
}

func qangleReader(bits uint) decodeFn {
	return func(br bit.Reader, d *Dict) (interface{}, error) {
		return &vector{
			x: bit.ReadAngle(br, bits),
			y: bit.ReadAngle(br, bits),
			z: bit.ReadAngle(br, bits),
		}, br.Err()
	}
}

func readQAngleCoords(br bit.Reader, d *Dict) (interface{}, error) {
	var (
		v vector
		x = bit.ReadBool(br)
		y = bit.ReadBool(br)
		z = bit.ReadBool(br)
	)
	if x {
		v.x = bit.ReadCoord(br)
	}
	if y {
		v.y = bit.ReadCoord(br)
	}
	if z {
		v.z = bit.ReadCoord(br)
	}
	return v, br.Err()
}

func readQAngleFloats(br bit.Reader, d *Dict) (interface{}, error) {
	return &vector{
		x: math.Float32frombits(uint32(br.ReadBits(32))),
		y: math.Float32frombits(uint32(br.ReadBits(32))),
		z: math.Float32frombits(uint32(br.ReadBits(32))),
	}, br.Err()
}
