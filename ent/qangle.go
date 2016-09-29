package ent

import (
	"fmt"
	"math"

	"github.com/jordanorelli/hyperstone/bit"
)

var qangle_t = &typeLiteral{
	name: "qangle",
	newFn: func() value {
		return new(qangle)
	},
}

type qangle struct{ pitch, yaw, roll float32 }

func (q qangle) tÿpe() tÿpe { return qangle_t }
func (q *qangle) read(r bit.Reader) error {
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
	return r.Err()
}

func (q qangle) String() string {
	return fmt.Sprintf("qangle{%f %f %f}", q.pitch, q.yaw, q.roll)
}

func qAngleType(spec *typeSpec, env *Env) tÿpe {
	if spec.typeName != "QAngle" {
		return nil
	}
	if spec.encoder == "qangle_pitch_yaw" {
		switch {
		case spec.bits <= 0 || spec.bits > 32:
			return typeError("qangle pitch_yaw has invalid bit size: %d", spec.bits)
		case spec.bits == 32:
			return pitchYaw_t
		default:
			t := pitchYawAngles_t(spec.bits)
			return &t
		}
	}
	switch spec.bits {
	case 0:
		Debug.Printf("  qangle type")
		return qangle_t
	case 32:
		return nil
	default:
		return nil
	}
}

var pitchYaw_t = &typeLiteral{
	name: "qangle:pitchYaw",
	newFn: func() value {
		return new(pitchYaw_v)
	},
}

type pitchYaw_v qangle

func (v pitchYaw_v) tÿpe() tÿpe { return pitchYaw_t }
func (v *pitchYaw_v) read(r bit.Reader) error {
	v.pitch = math.Float32frombits(uint32(r.ReadBits(32)))
	v.yaw = math.Float32frombits(uint32(r.ReadBits(32)))
	return r.Err()
}

func (v pitchYaw_v) String() string {
	return fmt.Sprintf("qangle:pitchYaw{%f %f}", v.pitch, v.yaw)
}

type pitchYawAngles_t uint

func (t pitchYawAngles_t) typeName() string { return "qangle:pitchYawAngles" }
func (t *pitchYawAngles_t) nü() value {
	return &pitchYawAngles_v{t: t}
}

type pitchYawAngles_v struct {
	t *pitchYawAngles_t
	qangle
}

func (v pitchYawAngles_v) tÿpe() tÿpe { return v.t }
func (v *pitchYawAngles_v) read(r bit.Reader) error {
	v.pitch = bit.ReadAngle(r, uint(*v.t))
	v.yaw = bit.ReadAngle(r, uint(*v.t))
	return r.Err()
}

func (v pitchYawAngles_v) String() string {
	return fmt.Sprintf("qangle:pitchYawAngles{%f %f}", v.pitch, v.yaw)
}
