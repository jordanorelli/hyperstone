package ent

import (
	"github.com/jordanorelli/hyperstone/bit"
	"github.com/jordanorelli/hyperstone/dota"
)

// a handle represents a soft pointer to an entity. handles are represented by
// IDs and can cross the client-server divide.
type handle int

func handleType(flat *dota.ProtoFlattenedSerializerFieldT, env *Env) tÿpe {
	if env.symbol(int(flat.GetVarTypeSym())) != "CGameSceneNodeHandle" {
		return nil
	}

	return typeFn(func(r bit.Reader) (value, error) {
		return handle(bit.ReadVarInt(r)), r.Err()
	})
}
