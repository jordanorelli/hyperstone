package ent

import (
	"github.com/jordanorelli/hyperstone/bit"
)

func stringType(spec *typeSpec, env *Env) tÿpe {
	if spec.typeName != "CUtlStringToken" {
		return nil
	}
	return typeFn(func(r bit.Reader) (value, error) {
		return bit.ReadVarInt(r), r.Err()
	})
}
