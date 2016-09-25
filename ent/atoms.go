package ent

import (
	"github.com/jordanorelli/hyperstone/bit"
)

var atom_types = map[string]typeFn{
	"bool": func(r bit.Reader) (value, error) {
		return bit.ReadBool(r), r.Err()
	},
	"uint8": func(r bit.Reader) (value, error) {
		// TODO: bounds check here
		return uint8(bit.ReadVarInt(r)), r.Err()
	},
	"uint16": func(r bit.Reader) (value, error) {
		// TODO: bounds check here
		return uint16(bit.ReadVarInt(r)), r.Err()
	},
	"uint64": func(r bit.Reader) (value, error) {
		return bit.ReadVarInt(r), r.Err()
	},
	"int8": func(r bit.Reader) (value, error) {
		// TODO: bounds check here
		return int8(bit.ReadZigZag32(r)), r.Err()
	},
	"int32": func(r bit.Reader) (value, error) {
		return bit.ReadZigZag32(r), r.Err()
	},
	"CUtlStringToken": func(r bit.Reader) (value, error) {
		return bit.ReadVarInt(r), r.Err()
	},
}

func atomType(spec *typeSpec, env *Env) tÿpe {
	if t, ok := atom_types[spec.typeName]; ok {
		Debug.Printf("  atom type")
		return t
	}
	return nil
}
