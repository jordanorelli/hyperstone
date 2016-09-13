package ent

import (
	"github.com/jordanorelli/hyperstone/bit"
)

type slotted interface {
	slotName(int) string
	slotValue(int) interface{}
	slotType(int) string
	slotDecoder(int) decoder
	setSlotValue(int, interface{})
}

func fillSlots(dest slotted, displayPath string, sr *selectionReader, br bit.Reader) error {
	selections, err := sr.readSelections(br, htree)
	if err != nil {
		return err
	}

	for _, s := range selections {
		if err := s.fill(0, displayPath, dest, br); err != nil {
			return err
		}
	}
	return nil
}
