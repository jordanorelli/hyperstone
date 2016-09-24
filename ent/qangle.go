package ent

import (
	"github.com/jordanorelli/hyperstone/bit"
	"github.com/jordanorelli/hyperstone/dota"
)

func qAngleType(flat *dota.ProtoFlattenedSerializerFieldT, env *Env) tÿpe {
	if env.symbol(int(flat.GetVarTypeSym())) != "QAngle" {
		return nil
	}
	switch flat.GetBitCount() {
	case 0:
		return typeFn(func(r bit.Reader) (value, error) {
			x, y, z := bit.ReadBool(r), bit.ReadBool(r), bit.ReadBool(r)
			var v vector
			if x {
				v.x = bit.ReadCoord(r)
			}
			if y {
				v.y = bit.ReadCoord(r)
			}
			if z {
				v.z = bit.ReadCoord(r)
			}
			return v, nil
		})
	case 32:
		return nil
	default:
		return nil
	}
}
