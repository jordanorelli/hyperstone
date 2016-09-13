package ent

import (
	"math"

	"github.com/jordanorelli/hyperstone/bit"
)

const (
	f_round_down = 1 << iota
	f_round_up
	f_encode_zero
	f_encode_ints
)

func floatDecoder(f *Field) decoder {
	if f.bits <= 0 || f.bits >= 32 {
		return ieeeFloat32Decoder
	}

	// a quantized field value must have some range specified, otherwise the
	// quantization makes no sense.
	if f.low == 0 && f.high == 0 {
		panic("quantization rules make no sense")
	}

	flags := f.flags

	// number of input steps
	// steps := int(1<<f.bits - 1)

	// keep the inverse to mult instead of divide later
	// inv_steps := 1.0 / float32(steps)

	// total range of values
	span := f.high - f.low

	if span < 0 {
		panic("quantization span is backwards")
	}

	if flags&f_round_down&f_round_up > 0 {
		panic("how can you round down and up at the same time")
	}

	// output width of each step
	// step_width := span * inv_steps

	return func(br bit.Reader) interface{} {
		if flags&f_round_down > 0 {
			return nil
		}
		if flags&f_round_up > 0 {
			panic("round up flag not done yet")
		}
		if flags&f_encode_zero > 0 {
			panic("encode zero flag not done yet")
		}
		if flags&f_encode_ints > 0 {
			panic("encode ints flag not done yet")
		}
		return nil
	}
}

// reads an IEEE 754 binary float value off of the stream
func ieeeFloat32Decoder(br bit.Reader) interface{} {
	return math.Float32frombits(uint32(br.ReadBits(32)))
}
