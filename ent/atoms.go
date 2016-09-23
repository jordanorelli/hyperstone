package ent

import (
	"github.com/jordanorelli/hyperstone/bit"
)

var atom_types = map[string]tÿpe{
	"uint16": {"uint16", func(...interface{}) value { return new(uint16_v) }},
}

type uint16_v uint16

func (u *uint16_v) read(r bit.Reader) error {
	*u = uint16_v(bit.ReadVarInt(r))
	return r.Err()
}
