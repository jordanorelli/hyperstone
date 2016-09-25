package ent

import (
	"github.com/golang/protobuf/proto"

	"github.com/jordanorelli/hyperstone/bit"
	"github.com/jordanorelli/hyperstone/dota"
)

type Env struct {
	symbols symbolTable
	source  bit.BufReader
	classes map[string]classHistory
	fields  []field
}

func (e *Env) Handle(m proto.Message) error {
	switch v := m.(type) {
	case *dota.CDemoSendTables:
		return e.mergeSendTables(v)
	}
	return nil
}

func (e *Env) setSource(buf []byte) {
	e.source.SetSource(buf)
}

func (e *Env) mergeSendTables(m *dota.CDemoSendTables) error {
	Debug.Printf("merge send tables")

	flat, err := getSerializers(m)
	if err != nil {
		return wrap(err, "unable to get serializers in sendtables")
	}
	e.symbols = symbolTable(flat.GetSymbols())
	e.stubClasses(flat)
	if err := e.parseFields(flat); err != nil {
		return wrap(err, "unable to parse serializer fields")
	}
	e.fillClasses(flat)
	return nil
}

// stubs out the classes to be created later. we do this to create empty class
// structs that fields may point to.
func (e *Env) stubClasses(flat *dota.CSVCMsg_FlattenedSerializer) {
	serializers := flat.GetSerializers()
	if e.classes == nil {
		e.classes = make(map[string]classHistory, len(serializers))
	}
	for _, s := range serializers {
		name := e.symbol(int(s.GetSerializerNameSym()))
		v := int(s.GetSerializerVersion())
		if e.classes[name] == nil {
			e.classes[name] = make(classHistory, 4)
		}
		e.classes[name][v] = &class{name: name, version: v}
	}
}

// parses the type definitions for each field. some fields have types that
// refer to class types defined in the replay file. classes must be declared up
// front via stubclasses prior to parseFields working correctly.
func (e *Env) parseFields(flat *dota.CSVCMsg_FlattenedSerializer) error {
	e.fields = make([]field, len(flat.GetFields()))
	for i, ff := range flat.GetFields() {
		f := &e.fields[i]
		if err := f.fromProto(ff, e); err != nil {
			return err
		}
	}
	return nil
}

// associates each class with its list of fields. parseFields must be run
// before fillClasses in order for the field definitions to exist.
func (e *Env) fillClasses(flat *dota.CSVCMsg_FlattenedSerializer) {
	for _, s := range flat.GetSerializers() {
		name := e.symbol(int(s.GetSerializerNameSym()))
		v := int(s.GetSerializerVersion())
		class := e.classes[name][v]

		class.fields = make([]field, len(s.GetFieldsIndex()))
		for i, id := range s.GetFieldsIndex() {
			class.fields[i] = e.fields[id]
		}
	}
}

func (e *Env) symbol(id int) string { return e.symbols[id] }
