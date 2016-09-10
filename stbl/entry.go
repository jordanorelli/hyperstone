package stbl

import (
	"fmt"
	"unicode/utf8"
)

// Entry represents a single record in a string table. It's not called "Record"
// because it's called "Entry" in the protobufs.
type Entry struct {
	Key   string
	Value []byte
}

func (e Entry) String() string {
	if e.Value == nil {
		return fmt.Sprintf("{%s nil}", e.Key)
	}

	if utf8.Valid(e.Value) {
		return fmt.Sprintf("{%s %s}", e.Key, e.Value)
	}

	if len(e.Value) > 32 {
		return fmt.Sprintf("{%s 0x%x}", e.Key, e.Value[:32])
	}
	return fmt.Sprintf("{%s 0x%x}", e.Key, e.Value)
}
