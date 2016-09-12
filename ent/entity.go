package ent

import (
	"fmt"

	"github.com/jordanorelli/hyperstone/bit"
)

type Entity struct {
	*Class
}

func (e *Entity) Read(br bit.Reader, sr *selectionReader) error {
	if e.Class == nil {
		return fmt.Errorf("unable to read entity: entity has no class")
	}
	Debug.Printf("entity %v read", e)

	if err := sr.read(br, htree); err != nil {
		return fmt.Errorf("unable to read entity: %v", err)
	}

	for _, s := range sr.selections() {
		switch s.count {
		case 0:
			Debug.Printf("FUCK!")
		case 1:
			Debug.Printf("direct selection: %v", s.path())
		default:
			Debug.Printf("child selection: %v", s.path())
		}
	}
	return nil
}
