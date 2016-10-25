package ent

import (
	"fmt"
	"github.com/jordanorelli/hyperstone/bit"
)

func vectorType(spec *typeSpec, env *Env) tÿpe {
	if spec.encoder != "" {
		return nil
	}
	t := floatType(spec, env)
	if _, ok := t.(error); ok {
		return t
	}
	if t == nil {
		return nil
	}
	return &vector_t{elem: t}
}

type vector_t struct {
	elem tÿpe
}

func (t vector_t) typeName() string {
	return fmt.Sprintf("vector<%s>", t.elem.typeName())
}

func (t *vector_t) nü() value {
	return &vector{t: t}
}

type vector struct {
	t       tÿpe
	x, y, z value
}

func (v vector) tÿpe() tÿpe { return v.t }

func (v *vector) read(r bit.Reader) error {
	if v.x == nil {
		v.x = v.t.nü()
	}
	if v.y == nil {
		v.y = v.t.nü()
	}
	if v.z == nil {
		v.z = v.t.nü()
	}

	type fn func(bit.Reader) error
	coalesce := func(fns ...fn) error {
		for _, f := range fns {
			if err := f(r); err != nil {
				return err
			}
		}
		return nil
	}

	return coalesce(v.x.read, v.y.read, v.z.read)
}

func (v vector) String() string {
	return fmt.Sprintf("vector<%s>{%s %s %s}", v.t.typeName, v.x, v.y, v.z)
}
