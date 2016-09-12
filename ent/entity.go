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
	if err := fp.read(br, htree, e.Class); err != nil {
		return fmt.Errorf("unable to read entity: %v", err)
	}
	for i := 0; i <= fp.hlast; i++ {
		if fp.history[i][0] == 0 {
			Debug.Printf("direct selection: %v", fp.history[i][1])
			Debug.Printf("field: %v", e.Class.Fields[fp.history[i][1]])
		} else {
			Debug.Printf("child selection: %v (%v)", fp.history[i],
				fp.history[i][1:fp.history[i][0]+2])
		}
	}
	return nil
}
