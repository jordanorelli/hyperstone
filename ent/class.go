package ent

import (
	"fmt"
	"github.com/jordanorelli/hyperstone/dota"
)

// Class represents a set of constraints around an Entity.
type Class struct {
	name    Symbol
	Version int
	Fields  []*Field

	// all other entities for this class use this instance as a prototype
	baseline *Entity

	// maps field names back to their indexes. Entities use this to access
	// their own fields by name instead of by slot.
	fieldNames map[string]int
}

func (c *Class) Name() string  { return c.name.String() }
func (c *Class) Slotted() bool { return true }
func (c *Class) Id() classId   { return classId{name: c.name, version: c.Version} }

func (c *Class) New(serial int, baseline bool) *Entity {
	e := &Entity{
		Class:      c,
		slots:      make([]interface{}, len(c.Fields)),
		serial:     serial,
		isBaseline: baseline,
	}
	for slot := range e.slots {
		e.slots[slot] = c.Fields[slot].initializer()
	}
	return e
}

func (c Class) String() string {
	return fmt.Sprintf("{%s %d}", c.Name, c.Version)
}

// A class is identified by the union of its name and version.
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
