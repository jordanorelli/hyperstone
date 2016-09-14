package ent

import (
	"strconv"
)

type cutlVector struct {
	slots     []interface{}
	_slotType string
	decoder
}

func (v *cutlVector) slotName(slot int) string       { return strconv.Itoa(slot) }
func (v *cutlVector) slotValue(slot int) interface{} { return v.slots[slot] }
func (v *cutlVector) slotType(slot int) string       { return v._slotType }
func (v *cutlVector) slotDecoder(slot int) decoder   { return v.decoder }

func (v *cutlVector) setSlotValue(slot int, val interface{}) {
	v.slots[slot] = val
}
