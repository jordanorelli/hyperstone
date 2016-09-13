package ent

type Entity struct {
	*Class
	serial int
	slots  []interface{}
}

func (e *Entity) slotName(n int) string             { return e.Class.Fields[n].name.String() }
func (e *Entity) slotType(n int) string             { return e.Class.Fields[n]._type.String() }
func (e *Entity) slotValue(n int) interface{}       { return e.slots[n] }
func (e *Entity) slotDecoder(n int) decoder         { return e.Class.Fields[n].decoder }
func (e *Entity) setSlotValue(n int, v interface{}) { e.slots[n] = v }
