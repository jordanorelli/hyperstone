package main

import (
	"fmt"
	"os"

	"github.com/golang/protobuf/proto"

	"github.com/jordanorelli/hyperstone/bit"
	"github.com/jordanorelli/hyperstone/dota"
	"github.com/jordanorelli/hyperstone/dt"
)

func sendTables(m proto.Message) {
	v, ok := m.(*dota.CDemoSendTables)
	if !ok {
		return
	}

	// sendtables only has one field, a binary data field.
	data := v.GetData()
	br := bit.NewBytesReader(data)

	// body is length-prefixed
	size := int(bit.ReadVarInt(br))

	buf := make([]byte, size)
	br.Read(buf)

	serializer := dota.CSVCMsg_FlattenedSerializer{}
	if err := proto.Unmarshal(buf, &serializer); err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	ts := dt.ParseFlattened(&serializer)
	ts.DebugPrint(os.Stdout)
}
