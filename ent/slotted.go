package ent

import (
	"fmt"
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
		return fmt.Errorf("error filling slots: %v", err)
	}

	for _, s := range selections {
		if err := s.fill(0, displayPath, dest, br); err != nil {
			return fmt.Errorf("error filling slots: %v", err)
		}
	}
	return nil
}
