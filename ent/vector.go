package ent

import (
	"fmt"
	"github.com/jordanorelli/hyperstone/bit"
)

type vector struct{ x, y, z float32 }

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
	return vector_t{elem: t}
}

type vector_t struct {
	elem tÿpe
}

func (t vector_t) typeName() string {
	return fmt.Sprintf("vector:%s", t.elem.typeName())
}

func (t vector_t) read(r bit.Reader) (value, error) {
	var err error
	var v interface{}
	read := func(f *float32) {
		if err != nil {
			return
		}
		v, err = t.elem.read(r)
		*f = v.(float32)
	}
	var out vector
	read(&out.x)
	read(&out.y)
	read(&out.z)
	return out, err
}
