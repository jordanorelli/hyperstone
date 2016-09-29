package ent

import (
	"github.com/golang/protobuf/proto"
	"strconv"

	"github.com/jordanorelli/hyperstone/bit"
	"github.com/jordanorelli/hyperstone/dota"
	"github.com/jordanorelli/hyperstone/stbl"
)

type Env struct {
	symbols symbolTable
	source  bit.BufReader
	classes map[string]*classHistory
	netIds  map[int]string
	fields  []field
	strings *stbl.Dict
}

func NewEnv() *Env {
	e := &Env{
		classes: make(map[string]*classHistory),
		strings: stbl.NewDict(),
	}
	e.strings.WatchTable("instancebaseline", e.syncBaselineTable)
	return e
}

func (e *Env) Handle(m proto.Message) error {
	switch v := m.(type) {
	case *dota.CSVCMsg_CreateStringTable:
		_, err := e.strings.Create(v)
		return err

	case *dota.CSVCMsg_UpdateStringTable:
		return e.strings.Update(v)

	case *dota.CDemoSendTables:
		if err := e.mergeSendTables(v); err != nil {
			return err
		}

	case *dota.CDemoClassInfo:
		e.mergeClassInfo(v)
		e.syncBaseline()
	}
	return nil
}

func (e *Env) setSource(buf []byte) {
	e.source.SetSource(buf)
}

// merges the "sendTables" message. "sendTables" is a deeply misleading name,
// but it's the name the structure is given in the protobufs. sendTables
// contains the type definitions that define our entity type system.
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
	for _, s := range serializers {
		name := e.symbol(int(s.GetSerializerNameSym()))
		v := int(s.GetSerializerVersion())
		c := &class{name: name, version: v}
		Debug.Printf("new class: %s", c)
		h := e.classes[name]
		if h == nil {
			h = new(classHistory)
			e.classes[name] = h
		}
		h.add(c)
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
		class := e.classes[name].version(v)

		class.fields = make([]field, len(s.GetFieldsIndex()))
		for i, id := range s.GetFieldsIndex() {
			Debug.Printf("class %s has field %s (%s)", name, e.fields[id].name, e.fields[id].typeName())
			class.fields[i] = e.fields[id]
		}
	}
}

func (e *Env) mergeClassInfo(m *dota.CDemoClassInfo) {
	if e.netIds == nil {
		e.netIds = make(map[int]string, len(m.GetClasses()))
	}
	for _, info := range m.GetClasses() {
		id := info.GetClassId()
		name := info.GetNetworkName()
		table := info.GetTableName()
		Debug.Printf("class info id: %d name: %s table: %s", id, name, table)
		e.netIds[int(id)] = name
	}
}

func (e *Env) symbol(id int) string { return e.symbols[id] }

func (e *Env) syncBaseline() {
	t := e.strings.TableForName("instancebaseline")
	if t != nil {
		e.syncBaselineTable(t)
	}
}

func (e *Env) syncBaselineTable(t *stbl.Table) {
	if e.netIds == nil || len(e.netIds) == 0 {
		Debug.Printf("syncBaselines skipped: net ids are nil")
	}
	if e.classes == nil || len(e.classes) == 0 {
		Debug.Printf("syncBaselines skipped: classes are nil")
	}

	r := new(bit.BufReader)
	sr := new(selectionReader)
	for _, entry := range t.Entries() {
		netId, err := strconv.Atoi(entry.Key)
		if err != nil {
			Debug.Printf("syncBaselines ignored bad key %s: %v", err)
			continue
		}
		className := e.netIds[netId]
		if className == "" {
			Debug.Printf("syncBaselines couldn't find class with net id %d", netId)
			continue
		}
		c := e.class(className)
		if c == nil {
			Debug.Printf("syncBaselines couldn't find class named %s", className)
			continue
		}
		Debug.Printf("syncBaselines key: %s className: %s", entry.Key, c.name)
		ent := c.nü().(*entity)
		r.SetSource(entry.Value)
		selections, err := sr.readSelections(r)
		if err != nil {
			Debug.Printf("unable to read selections for %s: %v", className, err)
			continue
		}
		Debug.Printf("selections: %v", selections)
		for _, s := range selections {
			if err := s.fillSlots(ent, r); err != nil {
				Debug.Printf("unable to fill slots for %s: %v", className, err)
			}
		}
	}
}

func (e *Env) class(name string) *class {
	h := e.classes[name]
	if h == nil {
		return nil
	}
	return h.newest
}

func (e *Env) classVersion(name string, version int) *class {
	h := e.classes[name]
	return h.version(version)
}
