package ent

type Entity struct {
	*Class
	serial int
	slots  []interface{}
}

func (e *Entity) getSlotValue(n int) interface{}    { return e.slots[n] }
func (e *Entity) setSlotValue(n int, v interface{}) { e.slots[n] = v }
func (e *Entity) getSlotDecoder(n int) decoder      { return e.Class.Fields[n].decoder }
