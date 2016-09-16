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
		class := Class{name: name, Version: version}
		class.fromProto(c, fields)

		id := classId{name: name, version: version}
		n.classes[id] = &class

		if n.classesByName[name.String()] == nil {
			n.classesByName[name.String()] = map[int]*Class{version: &class}
		} else {
			n.classesByName[name.String()][version] = &class
		}
	}

	for i := range fields {
		f := &fields[i]
		if f.serializer != nil {
			if f.serializerVersion != nil {
				f.class = n.classesByName[f.serializer.String()][int(*f.serializerVersion)]
			} else {
				f.class = n.NewestClass(f.serializer.String())
			}
		}

		// for some fields, we want to initialize a zero value on the baseline
		// instance. specifically, we do this for arrays so that we can't try
		// to index into nil.
		f.typeSpec = parseTypeName(n, f._type.String())
		if f.isContainer() {
			mf := f.memberField()
			fn := newFieldDecoder(n, mf)
			if f.typeSpec.kind == t_array {
				f.initializer = func() interface{} {
					return &array{
						slots:     make([]interface{}, f.typeSpec.size),
						_slotType: mf._type.String(),
						decoder:   fn,
					}
				}
			} else if f.typeSpec.kind == t_template && f.typeSpec.template == "CUtlVector" {
				f.initializer = func() interface{} {
					return &cutlVector{
						slots:     make([]interface{}, f.typeSpec.size),
						_slotType: mf._type.String(),
						decoder:   fn,
					}
				}
			}
		}

		// we also wait until after we've discovered all of the classes to
		// build the field decoder functions, because some fields are
		// themselves entities.
		f.decoder = newFieldDecoder(n, f)
	}

	for _, ff := range flat.GetFields() {
		t, err := parseType(n, ff)
		if err != nil {
			Debug.Printf("  parseType error: %v", err)
		} else {
			Debug.Printf("  parseType type: %v", t)
		}
	}

	return br.Err()
}

/*
func (n *Namespace) newMergeSendTables(st *dota.CDemoSendTables) error {
	// sendtables only has one field, a binary data field. It's not immediately
	// clear why this message exists at all when it could simply contain
	// CSVMsg_FlattenedSerializer.
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
	return n.mergeSerializers(&flat)
}

func (n *Namespace) mergeSerializers(flat *dota.CSVCMsg_FlattenedSerializer) error {
	// most of the names and associated strings are interned in a symbol table.
	// We'll need these first since they're referenced throughought the class
	// and field definitions.
	n.SymbolTable = SymbolTable(flat.GetSymbols())

	n.classes = make(map[classId]*Class, len(flat.GetSerializers()))
	n.classesByName = make(map[string]map[int]*Class, len(flat.GetSerializers()))

	// some fields refer to classes, but classes are collections of fields.
	// Their references are potentially cyclical. We start by creating empty
	// classes so that fields may point to them.
	for _, s := range flat.GetSerializers() {
		name_sym := n.Symbol(int(s.GetSerializerNameSym()))
		ver := int(s.GetSerializerVersion())
		class := Class{name: name_sym, Version: ver}
		n.classes[class.Id()] = &class
		if n.classesByName[class.Name()] == nil {
			n.classesByName[class.Name()] = make(map[int]*Class)
		}
		n.classesByName[class.Name()][ver] = &class
	}

	// type ProtoFlattenedSerializerFieldT struct {
	// 	VarTypeSym             *int32
	// 	VarNameSym             *int32
	// 	VarEncoderSym          *int32
	// 	FieldSerializerNameSym *int32
	// 	FieldSerializerVersion *int32
	// 	BitCount               *int32
	// 	LowValue               *float32
	// 	HighValue              *float32
	// 	EncodeFlags            *int32
	// 	SendNodeSym            *int32
	// }

	// next we parse all of the fields along with their type definitions
	fields := make(map[int]Field, len(flat.GetFields()))
	for id, flatField := range flat.GetFields() {
		t := parseType(n, flatField)
		fields[id].Name = n.Symbol(int(flatField.GetVarNameSym())).String()
	}
	os.Exit(1)
	return nil
}
*/

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
