package ent

import (
	"github.com/jordanorelli/hyperstone/bit"
)

type slotted interface {
	getSlotValue(int) interface{}
	setSlotValue(int, interface{})
	getSlotDecoder(int) decoder
}

func fillSlots(dest slotted, sr *selectionReader, br bit.Reader) error {
	selections, err := sr.readSelections(br, htree)
	if err != nil {
		return err
	}

	for _, s := range selections {
		if err := s.fill(0, dest, br); err != nil {
			return err
		}
	}
	return nil
}
