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
	name     string
	entries  []stringTableEntry
	byteSize int
	bitSize  int // this is in the protobuf message but I don't know what it does.
}

func (s stringTable) String() string {
	return fmt.Sprintf("{%s %s}", s.name, s.entries)
}

func (t *stringTable) create(br *bit.BufReader, entries int) {
	t.entries = make([]stringTableEntry, entries)
	var (
		base  uint64
		entry *stringTableEntry
	)

	for i := range t.entries {
		entry = &t.entries[i]
		if i > 32 {
			base++
		}

		// sequential index flag should always be true in create
		if !bit.ReadBool(br) {
			panic("weird")
		}

		// key flag
		if bit.ReadBool(br) {
			// backreading flag
			if bit.ReadBool(br) {
				entry.key = t.entries[base+br.ReadBits(5)].key[:br.ReadBits(5)] + bit.ReadString(br)
			} else {
				entry.key = bit.ReadString(br)
			}
		}

		// value flag
		if bit.ReadBool(br) {
			if t.byteSize != 0 {
				entry.value = make([]byte, t.byteSize)
				br.Read(entry.value)
			} else {
				size, _ := br.ReadBits(14), br.ReadBits(3)
				entry.value = make([]byte, size)
				br.Read(entry.value)
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
		// fmt.Println(s)
	case *dota.CSVCMsg_UpdateStringTable:
		prettyPrint(m)
		s.handleUpdate(v)
	case *dota.CSVCMsg_ClearAllStringTables:
		// prettyPrint(m)
	case *dota.CDemoStringTables:
		// prettyPrint(m)
	}
}

func (s *stringTables) handleCreate(m *dota.CSVCMsg_CreateStringTable) {
	fmt.Printf("create %s\n", m.GetName())
	if m.GetUserDataFixedSize() {
		s.tables = append(s.tables, stringTable{name: m.GetName(), byteSize: int(m.GetUserDataSize()), bitSize: int(m.GetUserDataSizeBits())})
	} else {
		s.tables = append(s.tables, stringTable{name: m.GetName()})
	}
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
	table.create(s.br, int(m.GetNumEntries()))
}

func (s *stringTables) handleUpdate(m *dota.CSVCMsg_UpdateStringTable) {
	// hazard
	table := &s.tables[m.GetTableId()]
	s.br.SetSource(m.GetStringData())
	fmt.Printf("update %s\n", table.name)
	table.update(s.br, int(m.GetNumChangedEntries()))
}

func (t *stringTable) update(br *bit.BufReader, changed int) {
	var (
		idx   = -1
		entry *stringTableEntry
	)
	h := newIntRing(32)
	for i := 0; i < changed; i++ {
		// sequential index flag should rarely be true in update
		if bit.ReadBool(br) {
			idx++
		} else {
			idx = int(bit.ReadVarInt(br)) + 1
		}

		h.add(idx)

		for idx > len(t.entries)-1 {
			t.entries = append(t.entries, stringTableEntry{})
		}

		entry = &t.entries[idx]
        fmt.Printf("%s -> ", entry)

		// key flag
		if bit.ReadBool(br) {
			// backreading flag
			if bit.ReadBool(br) {
				prev, pLen := h.at(int(br.ReadBits(5))), int(br.ReadBits(5))
				if prev < len(t.entries) {
					prevEntry := &t.entries[prev]
                    entry.key = prevEntry.key[:pLen] + bit.ReadString(br)
				} else {
					panic("backread error")
				}
			} else {
                entry.key = bit.ReadString(br)
			}
		}

		// value flag
		if bit.ReadBool(br) {
			if t.byteSize != 0 {
                if entry.value == nil {
                    entry.value = make([]byte, t.byteSize)
                }
			} else {
				size, _ := int(br.ReadBits(14)), br.ReadBits(3)
                if len(entry.value) < size {
                    entry.value = make([]byte, size)
                } else {
                    entry.value = entry.value[:size]
                }
			}
            br.Read(entry.value)
		}
        fmt.Printf("%s\n", entry)
	}
}
