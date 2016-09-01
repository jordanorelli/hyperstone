package dt

import (
	"github.com/jordanorelli/hyperstone/dota"
)

type Serializer struct {
	Name    Symbol
	Version int
	Fields  []*Field
}

func (s *Serializer) fromProto(v *dota.ProtoFlattenedSerializerT, st *SymbolTable, fields []Field) {
	s.Name = st.Symbol(int(v.GetSerializerNameSym()))
	s.Version = int(v.GetSerializerVersion())
	s.Fields = make([]*Field, len(v.GetFieldsIndex()))
	for i, fi := range v.GetFieldsIndex() {
		s.Fields[i] = &fields[fi]
	}
}
