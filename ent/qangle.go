package ent

import (
	"github.com/jordanorelli/hyperstone/bit"
)

func qAngleType(spec *typeSpec, env *Env) tÿpe {
	if spec.typeName != "QAngle" {
		return nil
	}
	switch spec.bits {
	case 0:
		Debug.Printf("  qangle type")
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
