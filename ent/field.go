package ent

import (
	"bytes"
	"fmt"

	"github.com/jordanorelli/hyperstone/dota"
)

type Field struct {
	_type             Symbol
	name              Symbol
	sendNode          Symbol
	bits              *int
	low               *float32
	high              *float32
	flags             *int32
	serializer        *Symbol
	serializerVersion *int32
	encoder           *Symbol
}

func (f Field) String() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "{type: %s name: %s send: %s", f._type, f.name, f.sendNode)
	if f.bits != nil {
		fmt.Fprintf(&buf, " bits: %d", *f.bits)
	}
	if f.low != nil {
		fmt.Fprintf(&buf, " low: %f", *f.low)
	}
	if f.high != nil {
		fmt.Fprintf(&buf, " high: %f", *f.high)
	}
	if f.flags != nil {
		fmt.Fprintf(&buf, " flags: %d", *f.flags)
	}
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
	if flat.BitCount == nil {
		f.bits = nil
	} else {
		f.bits = new(int)
		*f.bits = int(flat.GetBitCount())
	}
	f.low = flat.LowValue
	f.high = flat.HighValue
	f.flags = flat.EncodeFlags

	if flat.FieldSerializerNameSym == nil {
		f.serializer = nil
	} else {
		f.serializer = new(Symbol)
		*f.serializer = t.Symbol(int(flat.GetFieldSerializerNameSym()))
	}

	f.serializerVersion = flat.FieldSerializerVersion
	// panic if we don't have a send node cause that shit is corrupt yo
	f.sendNode = t.Symbol(int(*flat.SendNodeSym))
}
