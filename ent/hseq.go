package ent

import (
	"github.com/jordanorelli/hyperstone/bit"
	"github.com/jordanorelli/hyperstone/dota"
)

func hSeqType(flat *dota.ProtoFlattenedSerializerFieldT, env *Env) tÿpe {
	if env.symbol(int(flat.GetVarTypeSym())) != "HSequence" {
		return nil
	}

	Debug.Printf("  hsequence type")
	return typeFn(func(r bit.Reader) (value, error) {
		return bit.ReadVarInt(r) - 1, r.Err()
	})
}
