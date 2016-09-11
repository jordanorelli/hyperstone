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
	for fn := walk(htree, br); fn != nil; fn = walk(htree, br) {
		if err := br.Err(); err != nil {
			return fmt.Errorf("unable to read entity: reader error: %v", err)
		}
		fn(fp, br)
	}
	Debug.Printf("fieldpath %s", fp.pathString())
	return nil
}
