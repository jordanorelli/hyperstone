package ent

import (
	"github.com/jordanorelli/hyperstone/bit"
)

func hSeqType(spec *typeSpec, env *Env) tÿpe {
	if spec.typeName != "HSequence" {
		return nil
	}

	Debug.Printf("  hsequence type")
	return typeLiteral{
		"HSequence",
		func(r bit.Reader) (value, error) {
			return bit.ReadVarInt(r) - 1, r.Err()
		},
	}
}
