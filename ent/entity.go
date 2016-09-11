package ent

import (
	"fmt"

	"github.com/jordanorelli/hyperstone/bit"
)

type Entity struct {
	*Class
}

func (e *Entity) Read(br bit.Reader) error {
	if e.Class == nil {
		return fmt.Errorf("unable to read entity: entity has no class")
	}
	Debug.Printf("entity %v read", e)

	fp := newFieldPath()
	if err := fp.read(br, htree); err != nil {
		return fmt.Errorf("unable to read entity: %v", err)
	}
	Debug.Printf("fieldpath %v", fp.path())
	return nil
}
