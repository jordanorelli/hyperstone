package dt

import (
	"fmt"
	"io"

	"github.com/jordanorelli/hyperstone/dota"
)

// TableSet represents a collection of tables.
type TableSet struct {
	SymbolTable
	Fields      []Field
	Serializers []Serializer
}

func (t *TableSet) DebugPrint(w io.Writer) {
	fmt.Fprintln(w, "Symbols:")
	for _, sym := range t.SymbolTable {
		fmt.Fprintf(w, "\t%s\n", sym)
	}
	fmt.Fprintln(w, "Fields:")
	for _, f := range t.Fields {
		fmt.Fprintf(w, "\t%s\n", f)
	}
	fmt.Fprintln(w, "Serializers:")
	for _, s := range t.Serializers {
		fmt.Fprintf(w, "\t%s (%d):\n", s.Name, s.Version)
		for _, f := range s.Fields {
			fmt.Fprintf(w, "\t\t%s\n", f)
		}
	}
}

// ParseFlattened parses a flattened TableSet definition, as defined by the
// Dota replay protobufs.
func ParseFlattened(m *dota.CSVCMsg_FlattenedSerializer) *TableSet {
	ts := &TableSet{SymbolTable: SymbolTable(m.GetSymbols())}
	ts.parseFields(m.GetFields())
	ts.parseSerializers(m.GetSerializers())
	return ts
}

func (ts *TableSet) parseFields(flat []*dota.ProtoFlattenedSerializerFieldT) {
	ts.Fields = make([]Field, len(flat))
	for i, f := range flat {
		ts.Fields[i].fromProto(f, &ts.SymbolTable)
	}
}

func (ts *TableSet) parseSerializers(flat []*dota.ProtoFlattenedSerializerT) {
	ts.Serializers = make([]Serializer, len(flat))
	for i, s := range flat {
		ts.Serializers[i].fromProto(s, &ts.SymbolTable, ts.Fields)
	}
}
