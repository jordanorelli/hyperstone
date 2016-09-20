package ent

import (
	"github.com/jordanorelli/hyperstone/bit"
)

var primitives map[string]Primitive

// a Primitive type in the ent object system.
type Primitive struct {
	name  string
	read  decodeFn
	alloc func() interface{}
}

func (p *Primitive) New(args ...interface{}) interface{} {
	if p.alloc != nil {
		return p.alloc()
	}
	return nil
}
func (p *Primitive) Name() string  { return p.name }
func (p *Primitive) Slotted() bool { return false }

func (p *Primitive) Read(br bit.Reader, d *Dict) (interface{}, error) {
	return p.read(br, d)
}

func init() {
	types := []Primitive{
		{name: "bool", read: readBool},
		{name: "uint8", read: readUint8},
		{name: "uint16", read: readUint16},
		{name: "uint32", read: readUint32},
		{name: "uint64", read: readUint64},
		{name: "int8", read: readInt8},
		{name: "int16", read: readInt16},
		{name: "int32", read: readInt32},
		{name: "int64", read: readInt64},
		{name: "Color", read: readColor},
		{name: "HSequence", read: readIntMinusOne},
	}
	primitives = make(map[string]Primitive, len(types))
	for _, t := range types {
		primitives[t.name] = t
	}
}

func readBool(br bit.Reader, d *Dict) (interface{}, error) {
	return bit.ReadBool(br), br.Err()
}

// ------------------------------------------------------------------------------
// uints
// ------------------------------------------------------------------------------

func readUint8(br bit.Reader, d *Dict) (interface{}, error) {
	return uint8(br.ReadBits(8)), br.Err()
}

func readUint16(br bit.Reader, d *Dict) (interface{}, error) {
	return uint16(bit.ReadVarInt(br)), br.Err()
}

func readUint32(br bit.Reader, d *Dict) (interface{}, error) {
	return uint32(bit.ReadVarInt(br)), br.Err()
}

func readUint64(br bit.Reader, d *Dict) (interface{}, error) {
	return bit.ReadVarInt(br), br.Err()
}

// ------------------------------------------------------------------------------
// ints
// ------------------------------------------------------------------------------

func readInt8(br bit.Reader, d *Dict) (interface{}, error) {
	return int8(bit.ReadZigZag(br)), br.Err()
}

func readInt16(br bit.Reader, d *Dict) (interface{}, error) {
	return int16(bit.ReadZigZag(br)), br.Err()
}

func readInt32(br bit.Reader, d *Dict) (interface{}, error) {
	return int32(bit.ReadZigZag(br)), br.Err()
}

func readInt64(br bit.Reader, d *Dict) (interface{}, error) {
	return bit.ReadZigZag(br), br.Err()
}

func readColor(br bit.Reader, d *Dict) (interface{}, error) {
	u := bit.ReadVarInt(br)
	return color{
		r: uint8(u & 0xff000000),
		g: uint8(u & 0x00ff0000),
		b: uint8(u & 0x0000ff00),
		a: uint8(u & 0x000000ff),
	}, br.Err()
}

// what in the good fuck is this
func readIntMinusOne(br bit.Reader, d *Dict) (interface{}, error) {
	return bit.ReadVarInt(br) - 1, br.Err()
}
