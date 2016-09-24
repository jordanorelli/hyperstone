package ent

import (
	"fmt"

	"github.com/jordanorelli/hyperstone/bit"
	"github.com/jordanorelli/hyperstone/dota"
)

const (
	f_min = 1 << iota
	f_max
	f_center
)

func qFloatType(flat *dota.ProtoFlattenedSerializerFieldT, env *Env) tÿpe {
	if env.symbol(int(flat.GetVarTypeSym())) != "CNetworkedQuantizedFloat" {
		return nil
	}
	if flat.GetBitCount() < 0 {
		return typeError("quantized float has invalid negative bit count specifier")
	}
	if flat.GetHighValue()-flat.GetLowValue() < 0 {
		return typeError("quantized float has invalid negative range")
	}
	t := qfloat_t{
		bits:  uint(flat.GetBitCount()),
		low:   flat.GetLowValue(),
		high:  flat.GetHighValue(),
		flags: int(flat.GetEncodeFlags()) & 0x7,
	}
	t.span = t.high - t.low
	t.intervals = uint(1<<t.bits - 1)
	t.interval = t.span / float32(t.intervals)

	if t.flags > 0 {
		t.special = new(float32)
	}

	switch {
	case t.flags&f_min > 0:
		*t.special = t.low
	case t.flags&f_max > 0:
		*t.special = t.high
	case t.flags&f_center > 0:
		*t.special = t.low + (t.high+t.low)*0.5
	}
	return t
}

type qfloat_t struct {
	bits      uint
	low       float32
	high      float32
	flags     int
	span      float32 // total range of values
	intervals uint    // number of intervals in the quantization range
	interval  float32 // width of one interval
	special   *float32
}

func (t qfloat_t) read(r bit.Reader) (value, error) {
	if t.special != nil && bit.ReadBool(r) {
		return *t.special, nil
	}
	return nil, fmt.Errorf("I'll get there")
}
