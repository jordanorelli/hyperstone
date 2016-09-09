package ent

import (
	"fmt"

	"github.com/golang/protobuf/proto"

	"github.com/jordanorelli/hyperstone/bit"
	"github.com/jordanorelli/hyperstone/dota"
)

// Dict corresponds to the edict_t in Valve's documentation for the Source
// engine. See here: https://developer.valvesoftware.com/wiki/Edict_t
//
// From the Valve docs:
//   edict_t ("entity dictionary") is an interface struct that allows entities
//   to cross the client/server divide: with one attached, an entity has the
//   same index at both ends. The edict also manages the state of the entity's
//   DataTable and provides a common representation across all DLLs. It cannot
//   be used on the client.
type Dict struct {
	*Namespace
	entities map[int]Entity
	br       *bit.BufReader
}

func NewDict() *Dict {
	return &Dict{
		Namespace: new(Namespace),
		entities:  make(map[int]Entity),
		br:        new(bit.BufReader),
	}
}

// creates an entity with the provided id. the entity's contents data are read
// off of the Dict's internal bit stream br
func (d *Dict) createEntity(id int) error {
	classId := int(d.readClassId(d.br))
	if len(d.Namespace.classes) == 0 {
		return fmt.Errorf("unable to create entity %d: namespace has no classes", id)
	}
	d.br.ReadBits(17) // ???
	classV := int(bit.ReadVarInt(d.br))
	className := d.classIds[classId]
	class := d.Class(className, classV)
	if class == nil {
		return fmt.Errorf("unable to create entity %d: no class found for class name %s, version %d", className, classV)
	}
	Debug.Printf("create entity id: %d classId: %d className: %v class: %v\n", id, classId, className, class)
	e := class.New()
	e.Read(d.br)
	return nil
}

func (d *Dict) updateEntity(id int) error {
	Debug.Printf("update entity id: %d\n", id)
	return nil
}

func (d *Dict) deleteEntity(id int) error {
	Debug.Printf("delete entity id: %d\n", id)
	return nil
}

func (d *Dict) leaveEntity(id int) error {
	Debug.Printf("leave entity id: %d\n", id)
	return nil
}

func (d *Dict) Handle(m proto.Message) {
	switch v := m.(type) {
	case *dota.CDemoSendTables:
		d.mergeSendTables(v)

	case *dota.CDemoClassInfo:
		d.mergeClassInfo(v)

	case *dota.CSVCMsg_PacketEntities:
		d.mergeEntities(v)
	}
}

func (d *Dict) mergeEntities(m *dota.CSVCMsg_PacketEntities) error {
	data := m.GetEntityData()

	Debug.Printf("packet header MaxEntries: %d UpdatedEntries: %v IsDelta: %t UpdateBaseline: %t Baseline: %d DeltaFrom: %d PendingFullFrame: %t ActiveSpawngroupHandle: %d", m.GetMaxEntries(), m.GetUpdatedEntries(), m.GetIsDelta(), m.GetUpdateBaseline(), m.GetBaseline(), m.GetDeltaFrom(), m.GetPendingFullFrame(), m.GetActiveSpawngroupHandle())

	d.br.SetSource(data)
	id := -1
	// for i := 0; i < int(m.GetUpdatedEntries()); i++ {
	for i := 0; i < 1; i++ {
		id++
		// there may be a jump indicator, indicating how many id positions
		// to skip.
		id += int(bit.ReadUBitVar(d.br))

		// next two bits encode one of four entity mutate operations
		var fn func(int) error
		switch d.br.ReadBits(2) {
		case 0:
			fn = d.updateEntity
		case 1:
			fn = d.leaveEntity
		case 2:
			fn = d.createEntity
		case 3:
			fn = d.deleteEntity
		}

		if err := fn(id); err != nil {
			return fmt.Errorf("entity merge error: %v", err)
		}
	}
	return nil
}
