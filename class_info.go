package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/jordanorelli/hyperstone/dota"
)

// classInfo container contains info about entity classes found in a given
// replay
type classInfo struct {
	names map[int]string
}

func (c *classInfo) handle(m proto.Message) {
	v, ok := m.(*dota.CDemoClassInfo)
	if !ok {
		return
	}

	for _, class := range v.GetClasses() {
		fmt.Printf("class-id: %d network-name: %s table-name: %s\n", class.GetClassId(), class.GetNetworkName(), class.GetTableName())
		c.names[int(class.GetClassId())] = class.GetNetworkName()
	}
}
