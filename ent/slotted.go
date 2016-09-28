package ent

type slotted interface {
	slotType(int) tÿpe
	setSlotValue(int, value)
	getSlotValue(int) value
}
