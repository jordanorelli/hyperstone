package ent

import (
	"fmt"
	"github.com/jordanorelli/hyperstone/bit"
	"strconv"
)

// ------------------------------------------------------------------------------
// bool
// ------------------------------------------------------------------------------

var bool_t = &typeLiteral{
	name: "bool",
	newFn: func() value {
		return new(bool_v)
	},
}

type bool_v bool

func (v bool_v) tÿpe() tÿpe { return bool_t }

func (v *bool_v) read(r bit.Reader) error {
	*v = bool_v(bit.ReadBool(r))
	return r.Err()
}

func (v bool_v) String() string {
	if v {
		return "true"
	}
	return "false"
}

// ------------------------------------------------------------------------------
// uint8
// ------------------------------------------------------------------------------

var uint8_t = &typeLiteral{
	name: "uint8",
	newFn: func() value {
		return new(uint8_v)
	},
}

type uint8_v uint8

func (v uint8_v) tÿpe() tÿpe { return uint8_t }
func (v *uint8_v) read(r bit.Reader) error {
	*v = uint8_v(bit.ReadVarInt32(r))
	return r.Err()
}

func (v uint8_v) String() string {
	return strconv.FormatUint(uint64(v), 10)
}

// ------------------------------------------------------------------------------
// uint16
// ------------------------------------------------------------------------------

var uint16_t = &typeLiteral{
	name: "uint16",
	newFn: func() value {
		return new(uint16_v)
	},
}

type uint16_v uint16

func (v uint16_v) tÿpe() tÿpe { return uint16_t }
func (v *uint16_v) read(r bit.Reader) error {
	*v = uint16_v(bit.ReadVarInt32(r))
	return r.Err()
}

func (v uint16_v) String() string {
	return strconv.FormatUint(uint64(v), 10)
}

// ------------------------------------------------------------------------------
// uint32
// ------------------------------------------------------------------------------

var uint32_t = &typeLiteral{
	name: "uint32",
	newFn: func() value {
		return new(uint32_v)
	},
}

type uint32_v uint32

func (v uint32_v) tÿpe() tÿpe { return uint32_t }
func (v *uint32_v) read(r bit.Reader) error {
	*v = uint32_v(bit.ReadVarInt32(r))
	return r.Err()
}

func (v uint32_v) String() string {
	return strconv.FormatUint(uint64(v), 10)
}

// ------------------------------------------------------------------------------
// uint64
// ------------------------------------------------------------------------------

var uint64_t = &typeLiteral{
	name: "uint64",
	newFn: func() value {
		return new(uint64_v)
	},
}

type uint64_v uint64

func (v uint64_v) tÿpe() tÿpe { return uint64_t }
func (v *uint64_v) read(r bit.Reader) error {
	*v = uint64_v(bit.ReadVarInt32(r))
	return r.Err()
}

func (v uint64_v) String() string {
	return strconv.FormatUint(uint64(v), 10)
}

// ------------------------------------------------------------------------------
// uint64fixed
//
// (a uint64 value that is always represented on the wire with 64 bits)
// ------------------------------------------------------------------------------

var uint64fixed_t = &typeLiteral{
	name: "uint64fixed",
	newFn: func() value {
		return new(uint64fixed_v)
	},
}

type uint64fixed_v uint64

func (v uint64fixed_v) tÿpe() tÿpe { return uint64fixed_t }
func (v *uint64fixed_v) read(r bit.Reader) error {
	*v = uint64fixed_v(r.ReadBits(64))
	return r.Err()
}

func (v uint64fixed_v) String() string {
	return strconv.FormatUint(uint64(v), 10)
}

// ------------------------------------------------------------------------------
// int8
// ------------------------------------------------------------------------------

var int8_t = &typeLiteral{
	name: "int8",
	newFn: func() value {
		return new(int8_v)
	},
}

type int8_v int8

func (v int8_v) tÿpe() tÿpe { return int8_t }
func (v *int8_v) read(r bit.Reader) error {
	// TODO: bounds check here?
	*v = int8_v(bit.ReadZigZag32(r))
	return r.Err()
}

func (v int8_v) String() string {
	return strconv.FormatInt(int64(v), 10)
}

// ------------------------------------------------------------------------------
// int16
// ------------------------------------------------------------------------------

var int16_t = &typeLiteral{
	name: "int16",
	newFn: func() value {
		return new(int16_v)
	},
}

type int16_v int16

func (v int16_v) tÿpe() tÿpe { return int16_t }
func (v *int16_v) read(r bit.Reader) error {
	// TODO: bounds check here?
	*v = int16_v(bit.ReadZigZag32(r))
	return r.Err()
}

func (v int16_v) String() string {
	return strconv.FormatInt(int64(v), 10)
}

// ------------------------------------------------------------------------------
// int32
// ------------------------------------------------------------------------------

var int32_t = &typeLiteral{
	name: "int32",
	newFn: func() value {
		return new(int32_v)
	},
}

type int32_v int32

func (v int32_v) tÿpe() tÿpe { return int32_t }
func (v *int32_v) read(r bit.Reader) error {
	*v = int32_v(bit.ReadZigZag32(r))
	return r.Err()
}

func (v int32_v) String() string {
	return strconv.FormatInt(int64(v), 10)
}

// ------------------------------------------------------------------------------
// int64
// ------------------------------------------------------------------------------

var int64_t = &typeLiteral{
	name: "int64",
	newFn: func() value {
		return new(int64_v)
	},
}

type int64_v int64

func (v int64_v) tÿpe() tÿpe { return int64_t }
func (v *int64_v) read(r bit.Reader) error {
	*v = int64_v(bit.ReadZigZag(r))
	return r.Err()
}

func (v int64_v) String() string {
	return strconv.FormatInt(int64(v), 10)
}

// ------------------------------------------------------------------------------
// CUtlStringToken
//
// weirdly, this type isn't a string; it's actually a number. The number
// presumably indicates some value on a symbol table.
// ------------------------------------------------------------------------------

var stringToken_t = &typeLiteral{
	name: "CUtlStringToken",
	newFn: func() value {
		return new(stringToken_v)
	},
}

type stringToken_v uint64

func (v stringToken_v) tÿpe() tÿpe { return stringToken_t }
func (v *stringToken_v) read(r bit.Reader) error {
	*v = stringToken_v(bit.ReadVarInt(r))
	return r.Err()
}

func (v stringToken_v) String() string {
	return fmt.Sprintf("token:%d", v)
}

// ------------------------------------------------------------------------------
// Color
// ------------------------------------------------------------------------------

var color_t = &typeLiteral{
	name: "Color",
	newFn: func() value {
		return new(color)
	},
}

type color struct{ r, g, b, a uint8 }

func (c color) tÿpe() tÿpe { return color_t }
func (c *color) read(r bit.Reader) error {
	u := bit.ReadVarInt(r)
	c.r = uint8(u >> 6 & 0xff)
	c.g = uint8(u >> 4 & 0xff)
	c.b = uint8(u >> 2 & 0xff)
	c.a = uint8(u >> 0 & 0xff)
	return r.Err()
}

func (c color) String() string {
	return fmt.Sprintf("#%x%x%x%x", c.r, c.g, c.b, c.a)
}

// ------------------------------------------------------------------------------
// CUtlSymbolLarge
// ------------------------------------------------------------------------------

var cutl_string_t = typeLiteral{
	name: "CUtlSymbolLarge",
	newFn: func() value {
		return new(cutl_string_v)
	},
}

type cutl_string_v string

func (v cutl_string_v) tÿpe() tÿpe     { return cutl_string_t }
func (v cutl_string_v) String() string { return string(v) }
func (v *cutl_string_v) read(r bit.Reader) error {
	*v = cutl_string_v(bit.ReadString(r))
	return r.Err()
}

var atom_types = []tÿpe{
	bool_t,
	uint8_t,
	uint16_t,
	uint32_t,
	int8_t,
	int16_t,
	int32_t,
	int64_t,
	stringToken_t,
	color_t,
	cutl_string_t,
}

func atomType(spec *typeSpec, env *Env) tÿpe {
	for _, t := range atom_types {
		if t.typeName() == spec.typeName {
			Debug.Printf("  atom type: %s", t.typeName())
			if spec.bits != 0 {
				return typeError("spec can't be atom type: has bit specification: %v", spec)
			}
			if spec.encoder != "" {
				return typeError("spec can't be atom type: has encoder specification: %v", spec)
			}
			if spec.flags != 0 {
				return typeError("spec can't be atom type: has flags: %v", spec)
			}
			if spec.high != 0 {
				return typeError("spec can't be atom type: has high value constraint: %v", spec)
			}
			if spec.low != 0 {
				return typeError("spec can't be atom type: has low value constraint: %v", spec)
			}
			if spec.serializer != "" {
				return typeError("spec can't be atom type: has serializer: %v", spec)
			}
			return t
		}
	}
	if spec.typeName == "uint64" {
		if spec.encoder == "fixed64" {
			return uint64fixed_t
		}
		return uint64_t
	}
	return nil
}
