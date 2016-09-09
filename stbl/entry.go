package stbl

import (
	"fmt"
	"unicode/utf8"
)

// Entry represents a single record in a string table. It's not called "Record"
// because it's called "Entry" in the protobufs.
type Entry struct {
	key   string
	value []byte
}

func (e Entry) String() string {
	if e.value == nil {
		return fmt.Sprintf("{%s nil}", e.key)
	}

	if utf8.Valid(e.value) {
		return fmt.Sprintf("{%s %s}", e.key, e.value)
	}

	if len(e.value) > 32 {
		return fmt.Sprintf("{%s 0x%x}", e.key, e.value[:32])
	}
	return fmt.Sprintf("{%s 0x%x}", e.key, e.value)
}
