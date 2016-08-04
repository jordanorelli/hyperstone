package main

import (
	"fmt"
)

type entity struct {
	t    entityType
	size uint32
	body []byte
}

func (e entity) String() string {
	if len(e.body) < 30 {
		return fmt.Sprintf("{entity type: %s size: %d data: %x}", e.t, e.size, e.body)
	}
	return fmt.Sprintf("{entity type: %s size: %d data: %x...}", e.t, e.size, e.body[:27])
}
