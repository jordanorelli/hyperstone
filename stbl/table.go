package stbl

import (
	"fmt"

	"github.com/jordanorelli/hyperstone/bit"
)

type Table struct {
	name     string
	entries  []Entry
	byteSize int

	// this is in the protobuf message but I don't know what it does.
	bitSize int
}

func (t *Table) Entries() []Entry { return t.entries }

// creates n entries from the bit stream br
func (t *Table) createEntries(br *bit.BufReader, n int) error {
	Debug.Printf("table %s create %d entries", t.name, n)
	t.entries = make([]Entry, n)
	var (
		base  uint64
		entry *Entry
	)

	for i := range t.entries {
		entry = &t.entries[i]
		if i > 32 {
			base++
		}

		// sequential index flag should always be true in create
		if !bit.ReadBool(br) {
			return fmt.Errorf("stbl: unexpected nonsequential index")
		}

		// key flag: indicates that a key is present
		if bit.ReadBool(br) {
			// backreading flag: indicates that the key references an earlier
			// key or a portion of an earlier key as a prefix
			if bit.ReadBool(br) {
				entry.Key = t.entries[base+br.ReadBits(5)].Key[:br.ReadBits(5)] + bit.ReadString(br)
			} else {
				entry.Key = bit.ReadString(br)
			}
		}

		// value flag: indicates that a value is present
		if bit.ReadBool(br) {
			if t.byteSize != 0 {
				entry.Value = make([]byte, t.byteSize)
				br.Read(entry.Value)
			} else {
				size := br.ReadBits(14)
				br.ReadBits(3) // ???
				entry.Value = make([]byte, size)
				br.Read(entry.Value)
			}
		}
	}
	return br.Err()
}

func (t *Table) updateEntries(br *bit.BufReader, n int) error {
	Debug.Printf("table %s update %d entries", t.name, n)
	var (
		idx   = -1
		entry *Entry
	)
	h := newIntRing(32)
	for i := 0; i < n; i++ {
		// sequential index flag should rarely be true in update
		if bit.ReadBool(br) {
			idx++
		} else {
			idx = int(bit.ReadVarInt(br)) + 1
		}

		h.add(idx)

		// there's probably a faster way to grow the table here.
		for idx > len(t.entries)-1 {
			t.entries = append(t.entries, Entry{})
		}

		entry = &t.entries[idx]

		// key flag
		if bit.ReadBool(br) {
			// backreading flag
			if bit.ReadBool(br) {
				prev, pLen := h.at(int(br.ReadBits(5))), int(br.ReadBits(5))
				if prev < len(t.entries) {
					prevEntry := &t.entries[prev]
					entry.Key = prevEntry.Key[:pLen] + bit.ReadString(br)
				} else {
					return fmt.Errorf("backread error")
				}
			} else {
				entry.Key = bit.ReadString(br)
			}
		}

		// value flag
		if bit.ReadBool(br) {
			if t.byteSize != 0 {
				if entry.Value == nil {
					entry.Value = make([]byte, t.byteSize)
				}
			} else {
				size, _ := int(br.ReadBits(14)), br.ReadBits(3)
				if len(entry.Value) < size {
					entry.Value = make([]byte, size)
				} else {
					entry.Value = entry.Value[:size]
				}
			}
			br.Read(entry.Value)
		}
		Debug.Printf("%s:%s = %x", t.name, entry.Key, entry.Value)
	}
	return nil
}
