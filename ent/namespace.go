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

	fields := make([]Field, len(flat.GetFields()))
	for i, f := range flat.GetFields() {
		fields[i].fromProto(f, &n.SymbolTable)
	}

	n.classes = make(map[classId]*Class, len(flat.GetSerializers()))
	n.classesByName = make(map[string]map[int]*Class, len(flat.GetSerializers()))

	for _, c := range flat.GetSerializers() {
		name := n.Symbol(int(c.GetSerializerNameSym()))
		version := int(c.GetSerializerVersion())

		class := Class{Name: name, Version: version}
		class.fromProto(c, fields)

		Debug.Printf("new class: %v", class)

		id := classId{name: name, version: version}
		n.classes[id] = &class

		if n.classesByName[name.String()] == nil {
			n.classesByName[name.String()] = map[int]*Class{version: &class}
		} else {
			n.classesByName[name.String()][version] = &class
		}
	}

	return br.Err()
}

func (n *Namespace) readClassId(r bit.Reader) int {
	return int(r.ReadBits(uint(n.idBits)))
}

func (n *Namespace) Class(name string, version int) *Class {
	return n.classesByName[name][version]
}

func (n *Namespace) ClassByNetId(id int) *Class {
	name, ok := n.classIds[id]
	if !ok {
		Debug.Printf("can't find class name for net id %d", id)
		return nil
	}
	versions, newest := n.classesByName[name], -1
	for v, _ := range versions {
		if v > newest {
			newest = v
		}
	}
	if newest == -1 {
		Debug.Printf("class %s has no known versions in its version map", name)
		return nil
	}
	return versions[newest]
}
