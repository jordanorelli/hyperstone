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

type typeParseFn func(*typeSpec, *Env) tÿpe

func parseFieldType(flat *dota.ProtoFlattenedSerializerFieldT, env *Env) tÿpe {
	spec := new(typeSpec)
	spec.fromProto(flat, env)
	return parseTypeSpec(spec, env)
}

func parseTypeSpec(spec *typeSpec, env *Env) tÿpe {
	coalesce := func(fns ...typeParseFn) tÿpe {
		for _, fn := range fns {
			if t := fn(spec, env); t != nil {
				return t
			}
		}
		return nil
	}
	return coalesce(atomType, floatType, handleType, qAngleType, hSeqType)
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

type typeSpec struct {
	name        string
	typeName    string
	bits        uint
	low         float32
	high        float32
	flags       int
	serializer  string
	serializerV int
	send        string
	encoder     string
}

func (s *typeSpec) fromProto(flat *dota.ProtoFlattenedSerializerFieldT, env *Env) {
	s.name = env.symbol(int(flat.GetVarNameSym()))
	s.typeName = env.symbol(int(flat.GetVarTypeSym()))
	if flat.GetBitCount() < 0 {
		// this would cause ridiculously long reads later if we let it overflow
		panic("negative bit count: data is likely corrupt")
	}
	s.bits = uint(flat.GetBitCount())
	s.low = flat.GetLowValue()
	s.high = flat.GetHighValue()
	s.flags = int(flat.GetEncodeFlags())
	if flat.FieldSerializerNameSym != nil {
		s.serializer = env.symbol(int(*flat.FieldSerializerNameSym))
	}
	s.serializerV = int(flat.GetFieldSerializerVersion())
	if flat.SendNodeSym != nil {
		s.send = env.symbol(int(*flat.SendNodeSym))
	}
	if flat.VarEncoderSym != nil {
		s.encoder = env.symbol(int(*flat.VarEncoderSym))
	}
}
