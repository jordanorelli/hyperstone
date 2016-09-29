package ent

type slotted interface {
	value
	slotType(int) tÿpe
	slotName(int) string
	setSlotValue(int, value)
	getSlotValue(int) value
}
