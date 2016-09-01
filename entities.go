package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/jordanorelli/hyperstone/dota"
)

// type CSVCMsg_PacketEntities struct {
// 	MaxEntries                    *int32
// 	UpdatedEntries                *int32
// 	IsDelta                       *bool
// 	UpdateBaseline                *bool
// 	Baseline                      *int32
// 	DeltaFrom                     *int32
// 	EntityData                    []byte
// 	PendingFullFrame              *bool
// 	ActiveSpawngroupHandle        *uint32
// 	MaxSpawngroupCreationsequence *uint32
// }

func dumpEntities(m proto.Message) {
	switch v := m.(type) {
	case *dota.CSVCMsg_PacketEntities:
		data := v.GetEntityData()
		if len(data) > 32 {
			data = data[:32]
		}
		fmt.Printf("{MaxEntries: %d UpdatedEntries: %v IsDelta: %t UpdateBaseline: %t Baseline: %d DeltaFrom: %d EntityData: %x PendingFullFrame: %t ActiveSpawngroupHandle: %d}\n", v.GetMaxEntries(), v.GetUpdatedEntries(), v.GetIsDelta(), v.GetUpdateBaseline(), v.GetBaseline(), v.GetDeltaFrom(), data, v.GetPendingFullFrame(), v.GetActiveSpawngroupHandle())
	}
}
