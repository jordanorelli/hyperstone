package ent

import (
	"fmt"
	"github.com/jordanorelli/hyperstone/bit"
)

type Handle struct{ name string }

func (h *Handle) Name() string                   { return h.name }
func (h *Handle) New(...interface{}) interface{} { return nil }
func (h *Handle) Slotted() bool                  { return false }

func (h *Handle) Read(br bit.Reader, d *Dict) (interface{}, error) {
	id := int(bit.ReadVarInt(br))
	e, ok := d.hidx[id]
	if !ok {
		if br.Err() != nil {
			return nil, br.Err()
		}
		return nil, fmt.Errorf("no entity found with handle %d", id)
	}
	return e, br.Err()
}
