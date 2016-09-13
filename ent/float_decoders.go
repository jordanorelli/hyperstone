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

func floatDecoder(f *Field) decoder {
	if f.bits <= 0 || f.bits >= 32 {
		return ieeeFloat32Decoder
	}

	// a quantized field value must have some range specified, otherwise the
	// quantization makes no sense.
	if f.low == 0 && f.high == 0 {
		panic("quantization rules make no sense")
	}

	bits := f.bits
	low := f.low
	high := f.high

	flags := f.flags

	// there's a flag that's -8 and i don't know what to do with it. I'm just
	// gonna mask away everything except the three least significant bits and
	// pray for the best.
	flags = flags & 7

	// number of input steps
	steps := int(1<<f.bits - 1)

	// keep the inverse to mult instead of divide later
	inv_steps := 1.0 / float32(steps)

	// total range of values
	span := f.high - f.low

	if span < 0 {
		panic("quantization span is backwards")
	}

	// output width of each step
	step_width := span * inv_steps

	var special *float32
	switch {
	case flags&f_min > 0:
		special = new(float32)
		*special = low
	case flags&f_max > 0:
		special = new(float32)
		*special = high
	case flags&f_center > 0:
		special = new(float32)
		middle := (high + low) * 0.5
		// if we're within a step of zero just return zero.
		if middle > 0 && middle-step_width < 0 || middle < 0 && middle+step_width > 0 {
			middle = 0
		}
		*special = middle
	}

	return func(br bit.Reader) interface{} {
		if special != nil && bit.ReadBool(br) {
			Debug.Printf("decode float type: %s low: %f high: %f bits: %d steps: %d span: %f flags: %d special: %v", f._type.String(), low, high, bits, steps, span, flags, *special)
			return *special
		}
		u := br.ReadBits(bits)
		out := low + float32(u)*inv_steps*span
		Debug.Printf("decode float type: %s low: %f high: %f bits: %d bitVal: %d steps: %d span: %f flags: %d output: %v", f._type.String(), low, high, bits, u, steps, span, flags, out)
		return out
	}
}

// reads an IEEE 754 binary float value off of the stream
func ieeeFloat32Decoder(br bit.Reader) interface{} {
	return math.Float32frombits(uint32(br.ReadBits(32)))
}
