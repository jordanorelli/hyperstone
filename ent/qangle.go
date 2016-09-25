package ent

import (
	"github.com/jordanorelli/hyperstone/bit"
	"math"
)

type qangle struct{ pitch, yaw, roll float32 }

func qAngleType(spec *typeSpec, env *Env) tÿpe {
	if spec.typeName != "QAngle" {
		return nil
	}
	if spec.encoder == "qangle_pitch_yaw" {
		switch {
		case spec.bits <= 0 || spec.bits > 32:
			return typeError("qangle pitch_yaw has invalid bit size: %d", spec.bits)
		case spec.bits == 32:
			return typeFn(pitchYaw_t)
		default:
			return pitchYawAngles_t(spec.bits)
		}
	}
	switch spec.bits {
	case 0:
		Debug.Printf("  qangle type")
		return typeFn(func(r bit.Reader) (value, error) {
			var q qangle
			pitch, yaw, roll := bit.ReadBool(r), bit.ReadBool(r), bit.ReadBool(r)
			if pitch {
				q.pitch = bit.ReadCoord(r)
			}
			if yaw {
				q.yaw = bit.ReadCoord(r)
			}
			if roll {
				q.roll = bit.ReadCoord(r)
			}
			return q, nil
		})
	case 32:
		return nil
	default:
		return nil
	}
}

func pitchYaw_t(r bit.Reader) (value, error) {
	var q qangle
	q.pitch = math.Float32frombits(uint32(r.ReadBits(32)))
	q.yaw = math.Float32frombits(uint32(r.ReadBits(32)))
	return q, r.Err()
}

type pitchYawAngles_t uint

func (t pitchYawAngles_t) read(r bit.Reader) (value, error) {
	var q qangle
	q.pitch = bit.ReadAngle(r, uint(t))
	q.yaw = bit.ReadAngle(r, uint(t))
	return q, r.Err()
}
