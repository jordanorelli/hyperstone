package ent

import (
	"fmt"

	"github.com/jordanorelli/hyperstone/bit"
)

type handle_t string

func (t handle_t) typeName() string { return string(t) }
func (t *handle_t) nü() value       { return &handle{t: t} }

// a handle represents a soft pointer to an entity. handles are represented by
// IDs and can cross the client-server divide.
type handle struct {
	t  tÿpe
	id uint64
}

func (h handle) tÿpe() tÿpe { return h.t }
func (h *handle) read(r bit.Reader) error {
	h.id = bit.ReadVarInt(r)
	return r.Err()
}

func (h handle) String() string {
	return fmt.Sprintf("handle<%s>: %d", h.t.typeName(), h.id)
}

func handleType(spec *typeSpec, env *Env) tÿpe {
	if spec.typeName != "CGameSceneNodeHandle" {
		return nil
	}

	Debug.Printf("  handle type")
	t := handle_t(spec.typeName)
	return &t
}
