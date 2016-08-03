package main

import (
	"fmt"
)

type entity struct {
	t    uint32
	size uint32
	body []byte
}

func (e entity) String() string {
	if len(e.body) < 32 {
		return fmt.Sprintf("{entity type: %d size: %d data: %x}", e.t, e.size, e.body)
	}
	return fmt.Sprintf("{entity type: %d size: %d data: %x...}", e.t, e.size, e.body[:32])
}
