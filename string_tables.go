package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/golang/snappy"
	"github.com/jordanorelli/hyperstone/bit"
	"github.com/jordanorelli/hyperstone/dota"
	"os"
)

const (
	sTableRingSize = 32
)

type stringTables struct {
	tables  []stringTable
	idx     map[string]*stringTable
	br      *bit.BufReader
	scratch []byte
}

func (s *stringTables) String() string {
	if s.scratch == nil {
		return fmt.Sprintf("{%v %v %v nil}", s.tables, s.idx, s.br)
	}
	if len(s.scratch) > 32 {
		return fmt.Sprintf("{%v %v %v %x...}", s.tables, s.idx, s.br, s.scratch[:32])
	}
	return fmt.Sprintf("{%v %v %v %x}", s.tables, s.idx, s.br, s.scratch)
}

func newStringTables() *stringTables {
	return &stringTables{
		tables:  make([]stringTable, 0, 64),
		idx:     make(map[string]*stringTable, 64),
		br:      new(bit.BufReader),
		scratch: make([]byte, 1<<16),
	}
}

type stringTable struct {
    entries []stringTableEntry
}

func (s stringTable) String() string {
    return fmt.Sprintf("{%s}", s.entries)
}

func (t *stringTable) create(br *bit.BufReader, entries, byteSize, bitSize int) {
    t.entries = make([]stringTableEntry, entries)
	idx := -1
	for i, base := 0, uint64(0); i < entries; i++ {
		if i > 32 {
			base++
		}

		// continue flag
		if bit.ReadBool(br) {
			idx++
		} else {
			// in practice, the one replay I'm using never hits this branch, so
			// I do not know if it works. The base referenced from above might
			// be wrong in this branch.
			idx = int(bit.ReadVarInt(br))
		}

		// key flag
		if bit.ReadBool(br) {
			// backreading flag
			if bit.ReadBool(br) {
				t.entries[idx].key = t.entries[base+br.ReadBits(5)].key[:br.ReadBits(5)] + bit.ReadString(br)
			} else {
				t.entries[idx].key = bit.ReadString(br)
			}
		}

		// value flag
		if bit.ReadBool(br) {
			if byteSize != -1 {
				t.entries[idx].value = make([]byte, byteSize)
				br.Read(t.entries[idx].value)
			} else {
				size, _ := br.ReadBits(14), br.ReadBits(3)
				t.entries[idx].value = make([]byte, size)
				br.Read(t.entries[idx].value)
			}
		}
	}
}

type stringTableEntry struct {
	key   string
	value []byte
}

func (s stringTableEntry) String() string {
	if s.value == nil {
		return fmt.Sprintf("{%s nil}", s.key)
	}
	if len(s.value) > 32 {
		return fmt.Sprintf("{%s %x}", s.key, s.value[:32])
	}
	return fmt.Sprintf("{%s %x}", s.key, s.value)
}

func (s *stringTables) handle(m proto.Message) {
	switch v := m.(type) {
	case *dota.CSVCMsg_CreateStringTable:
		prettyPrint(m)
		s.handleCreate(v)
		fmt.Println(s)
	case *dota.CSVCMsg_UpdateStringTable:
		// prettyPrint(m)
		// s.handleUpdate(v)
	case *dota.CSVCMsg_ClearAllStringTables:
		// prettyPrint(m)
	case *dota.CDemoStringTables:
		// prettyPrint(m)
	}
}

func (s *stringTables) handleCreate(m *dota.CSVCMsg_CreateStringTable) {
	fmt.Printf("create %s\n", m.GetName())
    s.tables = append(s.tables, stringTable{})
	s.idx[m.GetName()] = &s.tables[len(s.tables)-1]
	table := &s.tables[len(s.tables)-1]

	sd := m.GetStringData()
	if sd == nil || len(sd) == 0 {
		return
	}

	if m.GetDataCompressed() {
		switch string(sd[:4]) {
		case "LZSS":
			// TODO: not this
			panic("no lzss support!")
		default:
			var err error
			sd, err = snappy.Decode(s.scratch, sd)
			if err != nil {
				fmt.Fprintf(os.Stderr, "stringtable decode error: %v", err)
				return
			}
		}
	}

	s.br.SetSource(sd)
	if m.GetUserDataFixedSize() {
		table.create(s.br, int(m.GetNumEntries()), int(m.GetUserDataSize()), int(m.GetUserDataSizeBits()))
	} else {
		table.create(s.br, int(m.GetNumEntries()), -1, -1)
	}
    fmt.Println(table)
}

// type CSVCMsg_UpdateStringTable struct {
// 	TableId           *int32
// 	NumChangedEntries *int32
// 	StringData        []byte
// }

func (s *stringTables) handleUpdate(m *dota.CSVCMsg_UpdateStringTable) {
    // hazard
    table := s.tables[m.GetTableId()]
    s.br.SetSource(m.GetStringData())
    table.update(s.br, int(m.GetNumChangedEntries()))
}

func (t *stringTable) update(br *bit.BufReader, changed int) {

}
