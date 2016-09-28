package ent

import (
	"github.com/jordanorelli/hyperstone/bit"
)

type entity struct {
	class *class
	slots []value
}

func (e *entity) read(r bit.Reader) error {
	Debug.Printf("entity %s read", e.className())
	sr := new(selectionReader)
	selections, err := sr.readSelections(r, htree)
	if err != nil {
		return wrap(err, "entity of type %s failed to read selections", e.className())
	}
	for _, s := range selections {
		if err := s.fillSlots(e, r); err != nil {
			return err
		}
	}
	return nil
}

func (e *entity) className() string {
	if e.class != nil {
		return e.class.name
	}
	return "<None>"
}

func (e *entity) slotType(i int) tÿpe         { return e.class.fields[i].tÿpe }
func (e *entity) setSlotValue(i int, v value) { e.slots[i] = v }
func (e *entity) getSlotValue(i int) value    { return e.slots[i] }
