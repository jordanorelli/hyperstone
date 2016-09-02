package ent

import (
	"github.com/jordanorelli/hyperstone/dota"
)

// Class represents a set of constraints around an Entity.
type Class struct {
	Name    Symbol
	Version int
	Fields  []*Field
}

type classId struct {
	name    Symbol
	version int
}

func (c *Class) fromProto(v *dota.ProtoFlattenedSerializerT, fields []Field) {
	c.Fields = make([]*Field, len(v.GetFieldsIndex()))
	for i, fi := range v.GetFieldsIndex() {
		c.Fields[i] = &fields[fi]
	}
}
