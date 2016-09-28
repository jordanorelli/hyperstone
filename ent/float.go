package ent

import (
	"math"

	"github.com/jordanorelli/hyperstone/bit"
)

const (
	f_min = 1 << iota
	f_max
	f_center
)

func floatType(spec *typeSpec, env *Env) tÿpe {
	switch spec.typeName {
	case "CNetworkedQuantizedFloat":
		return qFloatType(spec, env)
	case "float32":
	case "Vector":
	default:
		return nil
	}
	if spec.encoder == "coord" {
		Debug.Printf("  coord float type")
		return typeLiteral{
			"float:coord",
			func(r bit.Reader) (value, error) {
				return bit.ReadCoord(r), r.Err()
			},
		}
	}
	if spec.serializer == "simulationtime" {
		return nil
	}
	switch spec.bits {
	case 0, 32:
		Debug.Printf("  std float type")
		return typeLiteral{
			"float:std",
			func(r bit.Reader) (value, error) {
				// TODO: check uint32 overflow here?
				return math.Float32frombits(uint32(r.ReadBits(32))), r.Err()
			},
		}
	default:
		return qFloatType(spec, env)
	}
}

func qFloatType(spec *typeSpec, env *Env) tÿpe {
	if spec.bits < 0 {
		return typeError("quantized float has invalid negative bit count specifier")
	}
	if spec.high-spec.low < 0 {
		return typeError("quantized float has invalid negative range")
	}

	t := qfloat_t{typeSpec: *spec}
	t.span = t.high - t.low
	t.intervals = uint(1<<t.bits - 1)
	t.interval = t.span / float32(t.intervals)

	if t.flags > 0 {
		t.special = new(float32)
		switch t.flags {
		case f_min:
			*t.special = t.low
		case f_max:
			*t.special = t.high
		case f_center:
			*t.special = t.low + (t.high+t.low)*0.5
		default:
			return typeError("dunno how to handle qfloat flag value: %d", t.flags)
		}
	}

	Debug.Printf("  qfloat type")
	return t
}

type qfloat_t struct {
	typeSpec
	span      float32 // total range of values
	intervals uint    // number of intervals in the quantization range
	interval  float32 // width of one interval
	special   *float32
}

func (t qfloat_t) typeName() string { return "qfloat" }

func (t qfloat_t) read(r bit.Reader) (value, error) {
	if t.special != nil && bit.ReadBool(r) {
		return *t.special, nil
	}
	return t.low + float32(r.ReadBits(t.bits))*t.interval, r.Err()
}
