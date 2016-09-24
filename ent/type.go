package ent

import (
	"fmt"

	"github.com/jordanorelli/hyperstone/bit"
	"github.com/jordanorelli/hyperstone/dota"
)

type tÿpe interface {
	read(bit.Reader) (value, error)
}

type typeFn func(bit.Reader) (value, error)

func (fn typeFn) read(r bit.Reader) (value, error) { return fn(r) }

type typeParseFn func(*dota.ProtoFlattenedSerializerFieldT, *Env) tÿpe

func parseType(flat *dota.ProtoFlattenedSerializerFieldT, env *Env) tÿpe {
	coalesce := func(fns ...typeParseFn) tÿpe {
		for _, fn := range fns {
			if t := fn(flat, env); t != nil {
				return t
			}
		}
		return nil
	}
	return coalesce(atomType, floatType, handleType, qAngleType)
}

// a type error is both an error and a type. It represents a type that we were
// unable to correctly parse. It can be interpreted as an error or as a type;
// when interpreted as a type, it errors every time it tries to read a value.
func typeError(t string, args ...interface{}) tÿpe {
	Debug.Printf("  type error: %s", fmt.Sprintf(t, args...))
	return error_t(fmt.Sprintf(t, args...))
}

type error_t string

func (e error_t) Error() string { return string(e) }
func (e error_t) read(r bit.Reader) (value, error) {
	return nil, fmt.Errorf("type error: %s", string(e))
}
