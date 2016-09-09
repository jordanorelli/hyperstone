package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/jordanorelli/hyperstone/bit"
	"github.com/jordanorelli/hyperstone/dota"
	"github.com/jordanorelli/hyperstone/ent"
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

var ctx = ent.NewContext()

func dumpEntities(m proto.Message) {
	switch v := m.(type) {
	case *dota.CDemoSendTables:
		ctx.MergeSendTables(v)

	case *dota.CDemoClassInfo:
		ctx.MergeClassInfo(v)

	case *dota.CSVCMsg_PacketEntities:
		data := v.GetEntityData()
		var datap []byte
		if len(data) > 32 {
			datap = data[:32]
		} else {
			datap = data
		}

		fmt.Printf("{MaxEntries: %d UpdatedEntries: %v IsDelta: %t UpdateBaseline: %t Baseline: %d DeltaFrom: %d EntityData: %x PendingFullFrame: %t ActiveSpawngroupHandle: %d}\n", v.GetMaxEntries(), v.GetUpdatedEntries(), v.GetIsDelta(), v.GetUpdateBaseline(), v.GetBaseline(), v.GetDeltaFrom(), datap, v.GetPendingFullFrame(), v.GetActiveSpawngroupHandle())

		br := bit.NewBytesReader(data)
		id := -1
		// for i := 0; i < int(v.GetUpdatedEntries()); i++ {
		for i := 0; i < 1; i++ {
			id++
			// there may be a jump indicator, indicating how many id positions
			// to skip.
			id += int(bit.ReadUBitVar(br))

			// next two bits encode one of four entity mutate operations
			switch br.ReadBits(2) {
			case 0:
				ctx.UpdateEntity(id, br)
			case 1:
				ctx.LeaveEntity(id)
			case 2:
				ctx.CreateEntity(id, br)
			case 3:
				ctx.DeleteEntity(id)
			}
		}
	}
}
