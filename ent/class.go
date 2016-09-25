package ent

import (
	"fmt"
	"github.com/jordanorelli/hyperstone/bit"
)

type class struct {
	name    string
	version int
	fields  []field
}

func (c class) String() string {
	return fmt.Sprintf("<%s.%d>", c.name, c.version)
}

func (c *class) read(r bit.Reader) (value, error) {
	return nil, fmt.Errorf("fart")
}

type classHistory map[int]*class

func classType(spec *typeSpec, env *Env) tÿpe {
	if spec.serializer != "" {
		h := env.classes[spec.serializer]
		if h != nil {
			class := h[spec.serializerV]
			if class != nil {
				return class
			}
			return typeError("class %s exists for spec serializer but can't find version %d", spec.serializer, spec.serializerV)
		}
	}
	return nil
}
