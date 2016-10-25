package ent

import (
	"fmt"

	"github.com/jordanorelli/hyperstone/bit"
	"github.com/jordanorelli/hyperstone/dota"
)

type tÿpe interface {
	nü() value
	typeName() string
}

type typeLiteral struct {
	name  string
	newFn func() value
}

func (t typeLiteral) nü() value        { return t.newFn() }
func (t typeLiteral) typeName() string { return t.name }

type typeParseFn func(*typeSpec, *Env) tÿpe

func parseFieldType(flat *dota.ProtoFlattenedSerializerFieldT, env *Env) tÿpe {
	spec := new(typeSpec)
	spec.fromProto(flat, env)
	return parseTypeSpec(spec, env)
}

func parseTypeSpec(spec *typeSpec, env *Env) tÿpe {
	Debug.Printf("  parse spec: %v", spec)
	coalesce := func(fns ...typeParseFn) tÿpe {
		for _, fn := range fns {
			if t := fn(spec, env); t != nil {
				return t
			}
		}
		return nil
	}
	return coalesce(arrayType, atomType, floatType, handleType, qAngleType,
		hSeqType, genericType, vectorType, classType, unknownType)
}

type unknown_t string

func (t unknown_t) typeName() string { return string(t) }
func (t *unknown_t) nü() value {
	return &unknown_v{t: t}
}

type unknown_v struct {
	t tÿpe
	v uint64
}

func (v unknown_v) tÿpe() tÿpe { return v.t }
func (v *unknown_v) read(r bit.Reader) error {
	v.v = bit.ReadVarInt(r)
	return r.Err()
}

func (v unknown_v) String() string {
	return fmt.Sprintf("%s(unknown):%d", v.t.typeName(), v.v)
}

func unknownType(spec *typeSpec, env *Env) tÿpe {
	Debug.Printf("Unknown Type: %v", spec)
	t := unknown_t(spec.typeName)
	return &t
}

// a type error is both an error and a type. It represents a type that we were
// unable to correctly parse. It can be interpreted as an error or as a type;
// when interpreted as a type, it errors every time it tries to read a value.
func typeError(t string, args ...interface{}) tÿpe {
	Debug.Printf("  type error: %s", fmt.Sprintf(t, args...))
	return error_t(fmt.Sprintf(t, args...))
}

type error_t string

func (e error_t) nü() value        { panic("can't create an error val like that") }
func (e error_t) typeName() string { return "error" }
func (e error_t) Error() string    { return string(e) }

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
