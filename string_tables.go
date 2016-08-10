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

type stringTable []stringTableEntry

func (t stringTable) create(br *bit.BufReader, byteSize, bitSize int) {
	idx := -1
	for i, base := 0, uint64(0); i < len(t); i++ {
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
				t[idx].key = t[base+br.ReadBits(5)].key[:br.ReadBits(5)] + bit.ReadString(br)
			} else {
				t[idx].key = bit.ReadString(br)
			}
		}

		// value flag
		if bit.ReadBool(br) {
			if byteSize != -1 {
				t[idx].value = make([]byte, byteSize)
				br.Read(t[idx].value)
			} else {
				size, _ := br.ReadBits(14), br.ReadBits(3)
				t[idx].value = make([]byte, size)
				br.Read(t[idx].value)
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
	case *dota.CSVCMsg_ClearAllStringTables:
		// prettyPrint(m)
	case *dota.CDemoStringTables:
		// prettyPrint(m)
	}
}

// type CSVCMsg_CreateStringTable struct {
// 	Name              *string
// 	NumEntries        *int32
// 	UserDataFixedSize *bool
// 	UserDataSize      *int32
// 	UserDataSizeBits  *int32
// 	Flags             *int32
// 	StringData        []byte
// 	UncompressedSize  *int32
// 	DataCompressed    *bool
// }

func (s *stringTables) handleCreate(m *dota.CSVCMsg_CreateStringTable) {
	fmt.Printf("create %s\n", m.GetName())
	s.tables = append(s.tables, make(stringTable, m.GetNumEntries()))
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
		table.create(s.br, int(m.GetUserDataSize()), int(m.GetUserDataSizeBits()))
	} else {
		table.create(s.br, -1, -1)
	}
}
