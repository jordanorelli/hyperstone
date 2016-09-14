package ent

import (
	"strconv"
)

type array struct {
	slots     []interface{}
	_slotType string
	decoder
}

func (a *array) slotName(slot int) string               { return strconv.Itoa(slot) }
func (a *array) slotValue(slot int) interface{}         { return a.slots[slot] }
func (a *array) slotType(slot int) string               { return a._slotType }
func (a *array) slotDecoder(slot int) decoder           { return a.decoder }
func (a *array) setSlotValue(slot int, val interface{}) { a.slots[slot] = val }
