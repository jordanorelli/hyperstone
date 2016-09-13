package ent

import (
	"math"
	"strings"

	"github.com/jordanorelli/hyperstone/bit"
)

// a decoder decodes an entity value off of a bit reader
type decoder func(bit.Reader) interface{}

// creates a new field decoder for the field f.
func newFieldDecoder(n *Namespace, f *Field) decoder {
	Debug.Printf("new decoder: type: %s name: %s sendNode: %s\n\tbits: %d low: %v high: %v\n\tflags: %d serializer: %v serializerVersion: %v\n\tclass: %v encoder: %v", f._type, f.name, f.sendNode, f.bits, f.low, f.high, f.flags, f.serializer, f.serializerVersion, f.class, f.encoder)

	switch f._type.String() {
	case "bool":
		return decodeBool
	case "uint8", "uint16", "uint32", "uint64", "Color":
		return decodeVarInt64
	case "int8", "int16", "int32", "int64":
		return decodeZigZag
	case "float32":
		return floatDecoder(f)
	case "Vector":
		return vectorDecoder(f)
	}

	// the field is itself an entity contained within the outer entity.
	if f.class != nil {
		return entityDecoder(f.class)
	}

	switch {
	case strings.HasPrefix(f._type.String(), "CHandle"):
		return decodeVarInt32
	}
	return nil
}

func decodeBool(br bit.Reader) interface{}     { return bit.ReadBool(br) }
func decodeVarInt32(br bit.Reader) interface{} { return bit.ReadVarInt32(br) }
func decodeVarInt64(br bit.Reader) interface{} { return bit.ReadVarInt(br) }
func decodeZigZag(br bit.Reader) interface{}   { return bit.ReadZigZag(br) }

func floatDecoder(f *Field) decoder {
	if f.bits <= 0 || f.bits >= 32 {
		return ieeeFloat32Decoder
	}
	return nil
}

// reads an IEEE 754 binary float value off of the stream
func ieeeFloat32Decoder(br bit.Reader) interface{} {
	return math.Float32frombits(uint32(br.ReadBits(32)))
}

func entityDecoder(c *Class) decoder {
	return func(br bit.Reader) interface{} {
		// I have no idea what this bit means.
		return bit.ReadBool(br)
	}
}

func vectorDecoder(f *Field) decoder {
	if f.encoder != nil {
		switch f.encoder.String() {
		case "normal":
			return decodeNormalVector
		default:
			return nil
		}
	}

	fn := floatDecoder(f)
	if fn == nil {
		return nil
	}
	return func(br bit.Reader) interface{} {
		return vector{fn(br).(float32), fn(br).(float32), fn(br).(float32)}
	}
}

type vector [3]float32

func decodeNormalVector(br bit.Reader) interface{} {
	var v vector
	x, y := bit.ReadBool(br), bit.ReadBool(br)
	if x {
		v[0] = bit.ReadNormal(br)
	}
	if y {
		v[1] = bit.ReadNormal(br)
	}
	// yoooooo what in the good fuck is going on here
	p := v[0]*v[0] + v[1]*v[1]
	if p < 1.0 {
		v[2] = float32(math.Sqrt(float64(1.0 - p)))
	}
	if bit.ReadBool(br) {
		v[2] = -v[2]
	}
	return v
}
