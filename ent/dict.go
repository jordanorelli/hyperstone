package ent

import (
	"fmt"
	"strconv"

	"github.com/golang/protobuf/proto"

	"github.com/jordanorelli/hyperstone/bit"
	"github.com/jordanorelli/hyperstone/dota"
	"github.com/jordanorelli/hyperstone/stbl"
)

// https://developer.valvesoftware.com/wiki/Entity_limit
//
// There can be up to 4096 entities. This total is split into two groups of 2048:
//   Non-networked entities, which exist only on the client or server (e.g.
//   death ragdolls on client, logicals on server).
//   Entities with associated edicts, which can cross the client/server divide.
//
// If the game tries to assign a 2049th edict it will exit with an error
// message, but if it tries to create a 2049th non-networked entity it will
// merely refuse and print a warning to the console. The logic behind this
// may be that an entity spawned dynamically (i.e. not present in the map)
// but not assigned an edict probably isn't too important.
const e_limit = 2048

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
	entities []*Entity
	br       *bit.BufReader
	sr       *selectionReader

	// a reference to our string table of entity baseline data. For whatever
	// reason, the first set of baselines sometimes come in before the classes
	// are defined.
	base *stbl.Table
}

func NewDict(sd *stbl.Dict) *Dict {
	d := &Dict{
		Namespace: new(Namespace),
		entities:  make([]*Entity, e_limit),
		br:        new(bit.BufReader),
		sr:        new(selectionReader),
		base:      sd.TableForName("instancebaseline"),
	}
	sd.WatchTable("instancebaseline", d.updateBaselines)
	return d
}

// creates an entity with the provided id. the entity's contents data are read
// off of the Dict's internal bit stream br
func (d *Dict) createEntity(id int) error {
	classId := int(d.readClassId(d.br))
	if len(d.Namespace.classes) == 0 {
		return fmt.Errorf("unable to create entity %d: namespace has no classes", id)
	}
	serial := int(d.br.ReadBits(17))
	classV := int(bit.ReadVarInt(d.br))
	className := d.classIds[classId]
	class := d.Class(className, classV)
	if class == nil {
		return fmt.Errorf("unable to create entity %d: no class found for class name %s, version %d", className, classV)
	}
	e := class.New(serial, false)
	d.entities[id] = e
	Debug.Printf("create entity id: %d serial: %d classId: %d className: %v class: %v\n", id, serial, classId, className, class)
	return fillSlots(e, class.Name.String(), d.sr, d.br)
}

func (d *Dict) getEntity(id int) *Entity {
	if id < 0 || id >= e_limit {
		Debug.Printf("edict refused getEntity request for invalid id %d", id)
		return nil
	}
	return d.entities[id]
}

func (d *Dict) updateEntity(id int) error {
	Debug.Printf("update entity id: %d\n", id)
	e := d.getEntity(id)
	if e == nil {
		return fmt.Errorf("update entity %d refused: no such entity", id)
	}
	// e.Read(d.br, d.sr)
	return nil
}

func (d *Dict) deleteEntity(id int) error {
	Debug.Printf("delete entity id: %d\n", id)
	if id < 0 || id >= e_limit {
		return fmt.Errorf("delete entity %d refused: no such entity", id)
	}
	d.entities[id] = nil
	return nil
}

func (d *Dict) leaveEntity(id int) error {
	Debug.Printf("leave entity id: %d\n", id)
	// what the shit does this do?
	return nil
}

func (d *Dict) Handle(m proto.Message) {
	switch v := m.(type) {
	case *dota.CDemoSendTables:
		d.mergeSendTables(v)
		d.syncBaselines()

	case *dota.CDemoClassInfo:
		d.mergeClassInfo(v)
		d.syncBaselines()

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

func (d *Dict) updateBaselines(t *stbl.Table) {
	d.base = t
	d.syncBaselines()
}

func (d *Dict) syncBaselines() {
	if !d.hasClassinfo() {
		Debug.Printf("syncBaselines skip: no classInfo yet")
		return
	}
	Debug.Printf("syncBaselines start")
	if d.base == nil {
		Debug.Printf("syncBaselines failed: reference to baseline string table is nil")
		return
	}

	for _, e := range d.base.Entries() {
		id, err := strconv.Atoi(e.Key)
		if err != nil {
			Debug.Printf("syncBaselines skipping entry with key %s: key failed to parse to integer: %v", e.Key, err)
			continue
		}

		c := d.ClassByNetId(id)
		if c == nil {
			Debug.Printf("syncBaselines skipping entry with key %s: no such class", e.Key)
			continue
		}

		if c.baseline == nil {
			c.baseline = c.New(-1, true)
		}

		if e.Value == nil || len(e.Value) == 0 {
			Debug.Printf("syncBaselines skipping entry with key %s: value is empty", e.Key)
			continue
		}

		d.br.SetSource(e.Value)
		Debug.Printf("syncBaselines has new baseline for class %v", c)
		if err := fillSlots(c.baseline, c.Name.String(), d.sr, d.br); err != nil {
			Debug.Printf("syncBaselines failed to fill a baseline: %v", err)
			continue
		}
	}
}
