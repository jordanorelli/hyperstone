package ent

import (
	"github.com/jordanorelli/hyperstone/bit"
	"github.com/jordanorelli/hyperstone/dota"
)

func handleType(flat *dota.ProtoFlattenedSerializerFieldT, env *Env) tÿpe {
	if env.symbol(int(flat.GetVarTypeSym())) != "CGameSceneNodeHandle" {
		return nil
	}
	return handle_t{}
}

// a handle represents a soft pointer to an entity. handles are represented by
// IDs and can cross the client-server divide.
type handle int

type handle_t struct{}

func (t handle_t) read(r bit.Reader) (value, error) {
	return handle(bit.ReadVarInt(r)), nil
}
