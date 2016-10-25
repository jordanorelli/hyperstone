package ent

import (
	"fmt"
	"github.com/jordanorelli/hyperstone/bit"
)

type entity struct {
	class *class
	slots []value
}

func (e *entity) read(r bit.Reader) error {
	bit.ReadBool(r) // ???
	return nil
}

func (e *entity) className() string {
	if e.class != nil {
		return e.class.name
	}
	return "<None>"
}

func (e *entity) String() string {
	return fmt.Sprintf("%s{id: ?}", e.class.typeName())
}

func (e *entity) tÿpe() tÿpe { return e.class }
func (e *entity) slotType(i int) tÿpe {
	if i >= len(e.class.fields) {
		return typeError("index out of range in slotType: %d is beyond capacity %d", i, len(e.class.fields))
	}
	return e.class.fields[i].tÿpe
}
func (e *entity) slotName(i int) string       { return e.class.fields[i].name }
func (e *entity) setSlotValue(i int, v value) { e.slots[i] = v }
func (e *entity) getSlotValue(i int) value    { return e.slots[i] }
