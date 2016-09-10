package stbl

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/golang/snappy"

	"github.com/jordanorelli/hyperstone/bit"
	"github.com/jordanorelli/hyperstone/dota"
)

// Dict represents a dictionary of string tables. Each table may be referenced
// by either a numeric ID or its name.
type Dict struct {
	byId      []Table
	byName    map[string]*Table
	br        *bit.BufReader
	scratch   []byte
	observers map[string][]func(*Table)
}

func NewDict() *Dict {
	return &Dict{
		byId:      make([]Table, 0, 64),
		byName:    make(map[string]*Table, 64),
		br:        new(bit.BufReader),
		scratch:   make([]byte, 1<<16),
		observers: make(map[string][]func(*Table)),
	}
}

// creates a new table, appending it to our list of tables and adding an entry
// in the table name index
func (d *Dict) newTable(name string) *Table {
	d.byId = append(d.byId, Table{name: name})
	t := &d.byId[len(d.byId)-1]
	d.byName[name] = t
	return t
}

// Creates a string table based on the provided protobuf message. The table is
// retained in the dict, but a pointer to the table is also returned in case
// the newly-created table is of use to the caller.
func (d *Dict) Create(m *dota.CSVCMsg_CreateStringTable) (*Table, error) {
	defer d.notifyObservers(m.GetName())
	Debug.Printf("create table %s", m.GetName())
	t := d.newTable(m.GetName())

	if m.GetUserDataFixedSize() {
		t.byteSize = int(m.GetUserDataSize())
		t.bitSize = int(m.GetUserDataSizeBits())
	}

	data := m.GetStringData()
	if data == nil || len(data) == 0 {
		Debug.Printf("table %s created as empty table", m.GetName())
		return t, nil
	}

	if m.GetDataCompressed() {
		switch string(data[:4]) {
		case "LZSS":
			return nil, fmt.Errorf("stbl: LZSS compression is not supported")
		default:
			var err error
			data, err = snappy.Decode(d.scratch, data)
			if err != nil {
				return nil, fmt.Errorf("stbl: decode error: %v", err)
			}
		}
	}

	d.br.SetSource(data)
	if err := t.createEntries(d.br, int(m.GetNumEntries())); err != nil {
		return nil, err
	}
	return t, nil
}

// updates a string table in the dict based on the data found in a protobuf
// message
func (d *Dict) Update(m *dota.CSVCMsg_UpdateStringTable) error {
	Debug.Printf("dict: update %d entries in table having id %d", m.GetNumChangedEntries(), m.GetTableId())
	t := d.TableForId(int(m.GetTableId()))
	if t == nil {
		return fmt.Errorf("no known string table for id %d", m.GetTableId())
	}
	defer d.notifyObservers(t.name)

	d.br.SetSource(m.GetStringData())
	return t.updateEntries(d.br, int(m.GetNumChangedEntries()))
}

func (d *Dict) TableForId(id int) *Table {
	if id >= len(d.byId) {
		Debug.Printf("bad dict access: id %d is greater than the max table id %d", id, len(d.byId)-1)
		return nil
	}
	return &d.byId[id]
}

func (d *Dict) TableForName(name string) *Table {
	return d.byName[name]
}

func (d *Dict) Handle(m proto.Message) {
	switch v := m.(type) {
	case *dota.CSVCMsg_CreateStringTable:
		d.Create(v)
	case *dota.CSVCMsg_UpdateStringTable:
		d.Update(v)
	case *dota.CSVCMsg_ClearAllStringTables:
		Debug.Println("clear all string tables")
	case *dota.CDemoStringTables:
		Debug.Println("ignoring a full stringtable dump")
	}
}

func (d *Dict) WatchTable(name string, fn func(*Table)) {
	if d.observers[name] == nil {
		d.observers[name] = make([]func(*Table), 0, 8)
	}
	d.observers[name] = append(d.observers[name], fn)
}

func (d *Dict) notifyObservers(name string) {
	for _, fn := range d.observers[name] {
		fn(d.TableForName(name))
	}
}
