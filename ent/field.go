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
	Debug.Printf("parse flat field: %s", prettyFlatField(flat, env))
	t := parseType(flat, env)
	if t == nil {
		return fmt.Errorf("unable to parse type %s", prettyFlatField(flat, env))
	}
	if err, ok := t.(error); ok {
		return wrap(err, "unable to parse type %s", prettyFlatField(flat, env))
	}

	f.tÿpe = t
	f.name = env.symbol(int(flat.GetVarNameSym()))
	return nil
}
