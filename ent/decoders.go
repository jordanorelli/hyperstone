package ent

import (
	// "strings"

	"github.com/jordanorelli/hyperstone/bit"
)

// a decoder decodes an entity value off of a bit reader
type decoder func(bit.Reader) interface{}

// creates a new field decoder for the field f.
func newFieldDecoder(n *Namespace, f *Field) decoder {
	Debug.Printf("new decoder: type: %s name: %s sendNode: %s\n\tbits: %d low: %v high: %v\n\tflags: %d serializer: %v serializerVersion: %v\n\tclass: %v encoder: %v", f._type, f.name, f.sendNode, f.bits, f.low, f.high, f.flags, f.serializer, f.serializerVersion, f.class, f.encoder)

	typeName := f._type.String()

	switch typeName {
	case "bool":
		return decodeBool
	case "uint8", "uint16", "uint32", "uint64":
		return decodeVarInt64
	case "Color":
		return decodeColor
	case "int8", "int16", "int32", "int64":
		return decodeZigZag
	case "CNetworkedQuantizedFloat", "float32":
		return floatDecoder(f)
	case "Vector":
		return vectorDecoder(f)
	case "QAngle":
		return qangleDecoder(f)
	case "CGameSceneNodeHandle":
		// ehhh maybe no?
		return decodeVarInt32
	case "CUtlStringToken":
		return symbolDecoder(n)
	}

	// the field is itself an entity contained within the outer entity.
	if f.class != nil {
		return entityDecoder(f.class)
	}

	// for compound types such as handles, vectors (in the c++ std::vector
	// sense), and arrays
	switch f.typeSpec.kind {
	case t_element:
		Debug.Printf("weird typespec: we shouldn't have elements here: %v", f.typeSpec)
		return func(bit.Reader) interface{} {
			Info.Fatalf("unable to decode element of type: %v", f.typeSpec.name)
			return nil
		}
	case t_object:
		return func(br bit.Reader) interface{} {
			Debug.Printf("unable to decode object of type: %v", f.typeSpec.name)
			return decodeVarInt32(br)
		}
	case t_array:
		return arrayDecoder(n, f)
	case t_template:
		return templateDecoder(f)
	case t_pointer:
		return decodeBool
	}

	panic("fart")
}

func decodeBool(br bit.Reader) interface{}     { return bit.ReadBool(br) }
func decodeVarInt32(br bit.Reader) interface{} { return bit.ReadVarInt32(br) }
func decodeVarInt64(br bit.Reader) interface{} { return bit.ReadVarInt(br) }
func decodeZigZag(br bit.Reader) interface{}   { return bit.ReadZigZag(br) }

type color struct{ r, g, b, a uint8 }

func decodeColor(br bit.Reader) interface{} {
	u := bit.ReadVarInt(br)
	return color{
		r: uint8(u & 0xff000000),
		g: uint8(u & 0x00ff0000),
		b: uint8(u & 0x0000ff00),
		a: uint8(u & 0x000000ff),
	}
}

func entityDecoder(c *Class) decoder {
	return func(br bit.Reader) interface{} {
		bit.ReadBool(br) // what does this do
		return c.New(-1, false)
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

type vector struct {
	x float32
	y float32
	z float32
}

func decodeNormalVector(br bit.Reader) interface{} {
	var v vector
	x, y := bit.ReadBool(br), bit.ReadBool(br)
	if x {
		v.x = bit.ReadNormal(br)
	}
	if y {
		v.y = bit.ReadNormal(br)
	}

	// here comes a shitty hack!
	bit.ReadBool(br)
	// discard this flag, it's concerned with the sign of the z value, which
	// we're skipping.
	return v

	// ok, so. I have a suspicion that what we're interested in here is a
	// surface normal, and that it's describing something about the geometry of
	// the scene. I can't for the life of me see why in the hell we'd *ever*
	// care about this in the context of parsing a replay, so I'm just turning
	// off the z calculation. What follows is the original implementation as
	// adapted from Manta and Clarity. The reason I care to skip this is that
	// it involves a sqrt, which is a super slow operatation, and one worth
	// dispensing with if you don't need the data that it produces.
	//
	// p := v.x*v.x + v.y*v.y
	// if p < 1.0 {
	// 	v.z = float32(math.Sqrt(float64(1.0 - p)))
	// }
	// if bit.ReadBool(br) {
	// 	// we might wind up with the float representation of negative zero here,
	// 	// but as far as I can tell, negative zero is fine because negative
	// 	// zero is equivalent to positive zero. They'll print differently, but
	// 	// I don't think we care about that.
	// 	v.z = -v.z
	// }
	// return v
}

// decodes a qangle, which is a representation of an euler angle. that is, it's
// a three-angle encoding of orientation.
// https://developer.valvesoftware.com/wiki/QAngle
//
// the x, y, and z in this case do not refer to positions along a set of
// cartesian axes, but degress of rotation in an Euler angle.
//
//   (-45,10,0) means 45° up, 10° left and 0° roll.
func qangleDecoder(f *Field) decoder {
	if f.encoder != nil && f.encoder.String() == "qangle_pitch_yaw" {
		return nil
	}
	bits := f.bits
	if bits == 32 {
		return nil
	}
	if bits == 0 {
		return func(br bit.Reader) interface{} {
			var (
				v vector
				x = bit.ReadBool(br)
				y = bit.ReadBool(br)
				z = bit.ReadBool(br)
			)
			if x {
				v.x = bit.ReadCoord(br)
			}
			if y {
				v.y = bit.ReadCoord(br)
			}
			if z {
				v.z = bit.ReadCoord(br)
			}
			return v
		}
	}
	return func(br bit.Reader) interface{} {
		return vector{
			x: bit.ReadAngle(br, bits),
			y: bit.ReadAngle(br, bits),
			z: bit.ReadAngle(br, bits),
		}
	}
}

func arrayDecoder(n *Namespace, f *Field) decoder {
	return func(br bit.Reader) interface{} {
		Debug.Printf("dunno what this int32 val does in array decoder: %d", bit.ReadVarInt32(br))
		if f.initializer != nil {
			return f.initializer()
		}
		return nil
	}
}

func templateDecoder(f *Field) decoder {
	switch f.typeSpec.template {
	case "CHandle":
		return decodeVarInt32
	case "CStrongHandle":
		return decodeVarInt64
	case "CUtlVector":
		return func(br bit.Reader) interface{} {
			v := decodeVarInt32(br)
			Debug.Printf("dunno what this varint is for in the cutlvector decoder: %v", v)
			return v
		}
	}
	return nil
}

// so far a sanity check on the values I'm seeing out of this seem wrong.
func symbolDecoder(n *Namespace) decoder {
	return func(br bit.Reader) interface{} {
		u := bit.ReadVarInt32(br)
		return n.Symbol(int(u))
	}
}
