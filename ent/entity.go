package ent

import (
	"fmt"

	"github.com/jordanorelli/hyperstone/bit"
)

type Entity struct {
	*Class
	slots []interface{}
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
			panic("field selection makes no sense")
		case 1:
			Debug.Printf("direct selection: %v", s.path())
			field := e.Class.Fields[s.vals[0]]
			Debug.Printf("field: %v", e.Class.Fields[s.vals[0]])
			fn := field.decoder
			if fn == nil {
				Info.Fatalf("field has no decoder: %v", field)
			}
			v, err := fn(br), br.Err()
			Debug.Printf("value: %v err: %v", v, err)
			if err != nil {
				Info.Fatalf("field decode error: %v", err)
			}
		default:
			Debug.Printf("child selection: %v", s.path())
			return fmt.Errorf("child selections aren't done yet")
		}
	}
	return nil
}
