package ent

import (
	"math"
	"strconv"

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
		return coord_t
	}
	if spec.serializer == "simulationtime" {
		return nil
	}
	switch spec.bits {
	case 0, 32:
		Debug.Printf("  std float type")
		return float_t
	default:
		return qFloatType(spec, env)
	}
}

var coord_t = &typeLiteral{
	name: "coord",
	newFn: func() value {
		return new(coord_v)
	},
}

type coord_v float32

func (v coord_v) tÿpe() tÿpe { return coord_t }
func (v *coord_v) read(r bit.Reader) error {
	*v = coord_v(bit.ReadCoord(r))
	return r.Err()
}

func (v coord_v) String() string {
	return strconv.FormatFloat(float64(v), 'f', 3, 32)
}

var float_t = &typeLiteral{
	name: "float",
	newFn: func() value {
		return new(float_v)
	},
}

type float_v float32

func (v float_v) tÿpe() tÿpe { return float_t }
func (v *float_v) read(r bit.Reader) error {
	*v = float_v(math.Float32frombits(uint32(r.ReadBits(32))))
	return r.Err()
}

func (v float_v) String() string {
	return strconv.FormatFloat(float64(v), 'f', 3, 32)
}

func qFloatType(spec *typeSpec, env *Env) tÿpe {
	if spec.bits < 0 {
		return typeError("quantized float has invalid negative bit count specifier")
	}
	if spec.high-spec.low < 0 {
		return typeError("quantized float has invalid negative range")
	}

	t := &qfloat_t{typeSpec: *spec}
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

func (t *qfloat_t) nü() value       { return &qfloat_v{t: t} }
func (t qfloat_t) typeName() string { return "qfloat" }

type qfloat_v struct {
	t *qfloat_t
	v float32
}

func (v qfloat_v) tÿpe() tÿpe { return v.t }
func (v *qfloat_v) read(r bit.Reader) error {
	if v.t.special != nil && bit.ReadBool(r) {
		v.v = *v.t.special
	} else {
		v.v = v.t.low + float32(r.ReadBits(v.t.bits))*v.t.interval
	}
	return r.Err()
}

func (v qfloat_v) String() string {
	return strconv.FormatFloat(float64(v.v), 'f', 3, 32)
}
