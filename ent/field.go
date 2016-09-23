package ent

import (
	"fmt"
	"github.com/jordanorelli/hyperstone/dota"
)

type field struct {
	name string
	tÿpe
}

func (f *field) fromProto(flat *dota.ProtoFlattenedSerializerFieldT, env *Env) error {
	var_name := env.symbol(int(flat.GetVarNameSym()))
	var_type := env.symbol(int(flat.GetVarTypeSym()))

	if t, ok := atom_types[var_type]; ok {
		f.name = var_name
		f.tÿpe = t
		return nil
	}

	return fmt.Errorf("unable to parse type: %s", prettyFlatField(flat, env))
}
