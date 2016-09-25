package ent

import (
	"github.com/jordanorelli/hyperstone/bit"
)

var atom_types = map[string]typeFn{
	"uint16": func(r bit.Reader) (value, error) {
		// TODO: bounds check here
		return uint16(bit.ReadVarInt(r)), r.Err()
	},
	"int32": func(r bit.Reader) (value, error) {
		return int32(bit.ReadZigZag32(r)), r.Err()
	},
}

func atomType(spec *typeSpec, env *Env) tÿpe {
	if t, ok := atom_types[spec.typeName]; ok {
		Debug.Printf("  atom type")
		return t
	}
	return nil
}
