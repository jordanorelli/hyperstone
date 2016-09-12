package ent

import (
	"fmt"
	"math"

	"github.com/golang/protobuf/proto"

	"github.com/jordanorelli/hyperstone/bit"
	"github.com/jordanorelli/hyperstone/dota"
)

// Namespace registers the names of known classess, associating their ids to
// their network types.
type Namespace struct {
	SymbolTable
	idBits int

	// maps ClassInfo ids to class names
	classIds map[int]string

	// maps classes to their {name, version} id pairs
	classes map[classId]*Class

	// maps a class name to every version of the class
	classesByName map[string]map[int]*Class
}

// Merges in the ClassInfo data found in the replay protobufs
func (n *Namespace) mergeClassInfo(ci *dota.CDemoClassInfo) {
	if n.classIds == nil {
		n.classIds = make(map[int]string, len(ci.GetClasses()))
	}
	for _, class := range ci.GetClasses() {
		n.classIds[int(class.GetClassId())] = class.GetNetworkName()
	}
	n.idBits = int(math.Floor(math.Log2(float64(len(n.classIds))))) + 1
}

func (n *Namespace) hasClassinfo() bool {
	return n.classIds != nil && len(n.classIds) > 0
}

// merges the send table data found in the replay protobufs. The send table
// data contains a specification for an entity type system.
func (n *Namespace) mergeSendTables(st *dota.CDemoSendTables) error {
	// sendtables only has one field, a binary data field.
	Debug.Printf("merge send tables")
	data := st.GetData()
	br := bit.NewBytesReader(data)

	// body is length-prefixed
	size := int(bit.ReadVarInt(br))

	buf := make([]byte, size)
	br.Read(buf)

	flat := dota.CSVCMsg_FlattenedSerializer{}
	if err := proto.Unmarshal(buf, &flat); err != nil {
		return fmt.Errorf("unable to merge send tables: %v", err)
	}

	n.SymbolTable = SymbolTable(flat.GetSymbols())

	// the full set of fields that may appear on the classes is read first.
	// each class will have a list of fields.
	fields := make([]Field, len(flat.GetFields()))
	for i, f := range flat.GetFields() {
		fields[i].fromProto(f, &n.SymbolTable)
	}

	n.classes = make(map[classId]*Class, len(flat.GetSerializers()))
	n.classesByName = make(map[string]map[int]*Class, len(flat.GetSerializers()))

	// each serializer in the source data generates a class.
	for _, c := range flat.GetSerializers() {
		name := n.Symbol(int(c.GetSerializerNameSym()))
		version := int(c.GetSerializerVersion())

		Debug.Printf("new class: %s %v", name, version)
		class := Class{Name: name, Version: version}
		class.fromProto(c, fields)

		id := classId{name: name, version: version}
		n.classes[id] = &class

		if n.classesByName[name.String()] == nil {
			n.classesByName[name.String()] = map[int]*Class{version: &class}
		} else {
			n.classesByName[name.String()][version] = &class
		}
	}

	// some fields explicitly reference their origin class (P). that is is, if
	// a given field F is included in some class C, the field F having an
	// origin class P indicates that the class C has the class P as an
	// ancestor. since these references are circular, we unpacked the fields
	// first, then the classes, and now we re-visit the fields to set their
	// origin class pointers, now that the classes exist.
	for i := range fields {
		f := &fields[i]
		if f.serializer != nil {
			if f.serializerVersion != nil {
				f.class = n.classesByName[f.serializer.String()][int(*f.serializerVersion)]
			} else {
				f.class = n.NewestClass(f.serializer.String())
			}
		}

		// we also wait until after we've discovered all of the classes to
		// build the field decoder functions, because some fields are
		// themselves entities.
		f.decoder = newFieldDecoder(n, f)
	}

	return br.Err()
}

func (n *Namespace) readClassId(r bit.Reader) int {
	return int(r.ReadBits(uint(n.idBits)))
}

func (n *Namespace) Class(name string, version int) *Class {
	return n.classesByName[name][version]
}

// retrieves the newest version of a class, as referenced by name.
func (n *Namespace) NewestClass(name string) *Class {
	versions, newest := n.classesByName[name], -1
	for v, _ := range versions {
		if v > newest {
			newest = v
		}
	}
	if newest == -1 {
		Info.Fatalf("class %s has no known versions in its version map", name)
	}
	return versions[newest]
}

func (n *Namespace) ClassByNetId(id int) *Class {
	name, ok := n.classIds[id]
	if !ok {
		Info.Fatalf("can't find class name for net id %d", id)
	}
	return n.NewestClass(name)
}

func (n *Namespace) HasClass(name string) bool {
	_, ok := n.classesByName[name]
	return ok
}
