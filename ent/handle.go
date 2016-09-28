package ent

import (
	"github.com/jordanorelli/hyperstone/bit"
)

// a handle represents a soft pointer to an entity. handles are represented by
// IDs and can cross the client-server divide.
type handle int

func handleType(spec *typeSpec, env *Env) tÿpe {
	if spec.typeName != "CGameSceneNodeHandle" {
		return nil
	}

	Debug.Printf("  handle type")
	return typeLiteral{
		"handle:CGameSceneNodeHandle",
		func(r bit.Reader) (value, error) {
			return handle(bit.ReadVarInt(r)), r.Err()
		},
	}
}
