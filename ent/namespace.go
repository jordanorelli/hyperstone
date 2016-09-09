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
func (n *Namespace) MergeClassInfo(ci *dota.CDemoClassInfo) {
	if n.classIds == nil {
		n.classIds = make(map[int]string, len(ci.GetClasses()))
	}
	for _, class := range ci.GetClasses() {
		n.classIds[int(class.GetClassId())] = class.GetNetworkName()
	}
	n.idBits = int(math.Floor(math.Log2(float64(len(n.classIds))))) + 1
}

func (n *Namespace) MergeSendTables(st *dota.CDemoSendTables) {
	// sendtables only has one field, a binary data field.
	data := st.GetData()
	br := bit.NewBytesReader(data)

	// body is length-prefixed
	size := int(bit.ReadVarInt(br))

	buf := make([]byte, size)
	br.Read(buf)

	flat := dota.CSVCMsg_FlattenedSerializer{}
	if err := proto.Unmarshal(buf, &flat); err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	n.SymbolTable = SymbolTable(flat.GetSymbols())

	fields := make([]Field, len(flat.GetFields()))
	for i, f := range flat.GetFields() {
		fields[i].fromProto(f, &n.SymbolTable)
	}

	n.classes = make(map[classId]*Class, len(flat.GetSerializers()))
	n.classesByName = make(map[string]map[int]*Class, len(flat.GetSerializers()))

	for _, c := range flat.GetSerializers() {
		class := Class{}
		class.fromProto(c, fields)

		name := n.Symbol(int(c.GetSerializerNameSym()))
		class.Name = name
		version := int(c.GetSerializerVersion())
		id := classId{name: name, version: version}
		n.classes[id] = &class

		if n.classesByName[name.String()] == nil {
			n.classesByName[name.String()] = map[int]*Class{version: &class}
		} else {
			n.classesByName[name.String()][version] = &class
		}
	}
}

func (n *Namespace) readClassId(r bit.Reader) int {
	return int(r.ReadBits(uint(n.idBits)))
}

func (n *Namespace) Class(name string, version int) *Class {
	return n.classesByName[name][version]
}
