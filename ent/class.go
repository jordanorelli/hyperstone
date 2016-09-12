package ent

import (
	"fmt"
	"github.com/jordanorelli/hyperstone/dota"
)

// Class represents a set of constraints around an Entity.
type Class struct {
	Name    Symbol
	Version int
	Fields  []*Field

	// all other entities for this class use this instance as a prototype
	baseline *Entity

	// maps field names back to their indexes. Entities use this to access
	// their own fields by name instead of by slot.
	fieldNames map[string]int

	fp *fieldPath
}

func (c *Class) New() *Entity {
	return &Entity{Class: c}
}

func (c Class) String() string {
	return fmt.Sprintf("{%s %d}", c.Name, c.Version)
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
