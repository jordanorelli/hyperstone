package ent

import (
	"bytes"
	"fmt"

	"github.com/jordanorelli/hyperstone/dota"
)

/*
type Field struct {
	Name string
	Type
}

*/

type Field struct {
	_type             Symbol // type of data held by the field
	typeSpec          typeSpec
	name              Symbol  // name of the field
	sendNode          Symbol  // not sure what this is
	bits              uint    // number of bits used to encode field?
	low               float32 // lower limit of field values
	high              float32 // upper limit of field values
	flags             int     // used by float decoder
	serializer        *Symbol // the field is an entity with this class
	serializerVersion *int32
	class             *Class  // source class on which the field was originally defined
	encoder           *Symbol // binary encoder, named explicitly in protobuf
	decoder                   // decodes field values from a bit stream
	initializer       func() interface{}
	isTemplate        bool // whether or not the field is a template type
	templateType      string
	elemType          string
}

func (f Field) String() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "{type: %s name: %s send: %s", f._type, f.name, f.sendNode)
	if f.bits > 0 {
		fmt.Fprintf(&buf, " bits: %d", f.bits)
	}
	if f.flags > 0 {
		fmt.Fprintf(&buf, " flags: %d", f.flags)
	}
	fmt.Fprintf(&buf, " low: %f", f.low)
	fmt.Fprintf(&buf, " high: %f", f.high)
	if f.serializer != nil {
		fmt.Fprintf(&buf, " serializer: %s", *f.serializer)
	}
	if f.serializerVersion != nil {
		fmt.Fprintf(&buf, " serializer_v: %d", *f.serializerVersion)
	}
	if f.encoder != nil {
		fmt.Fprintf(&buf, " encoder: %s", *f.encoder)
	}
	fmt.Fprint(&buf, "}")
	return buf.String()
}

func (f *Field) fromProto(flat *dota.ProtoFlattenedSerializerFieldT, t *SymbolTable) {
	f._type = t.Symbol(int(flat.GetVarTypeSym()))
	f.name = t.Symbol(int(flat.GetVarNameSym()))
	f.bits = uint(flat.GetBitCount())
	f.flags = int(flat.GetEncodeFlags())
	f.low = flat.GetLowValue()
	f.high = flat.GetHighValue()
	f.initializer = nilInitializer

	if flat.FieldSerializerNameSym == nil {
		f.serializer = nil
	} else {
		f.serializer = new(Symbol)
		*f.serializer = t.Symbol(int(flat.GetFieldSerializerNameSym()))
	}

	f.serializerVersion = flat.FieldSerializerVersion
	// panic if we don't have a send node cause that shit is corrupt yo
	f.sendNode = t.Symbol(int(*flat.SendNodeSym))
	Debug.Printf("new field: %v", f)
}

// creates a new field which is a sort of virtual field that represents what a
// field woudl look like if we had one for a container field's elements.
// honestly this is a really shitty hack it just seems easier than rewriting
// the newFieldDecoder logic.
func (f *Field) memberField() *Field {
	mf := new(Field)
	*mf = *f
	mf.typeSpec = *f.typeSpec.member
	// yeahhhh
	mf._type = Symbol{0, &SymbolTable{mf.typeSpec.name}}
	return mf
}

func (f *Field) isContainer() bool {
	if f.typeSpec.kind == t_array {
		return true
	}
	if f.typeSpec.kind == t_template {
		if f.typeSpec.template == "CUtlVector" {
			return true
		}
	}
	return false
}

func nilInitializer() interface{} { return nil }
