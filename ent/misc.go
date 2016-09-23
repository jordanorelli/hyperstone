package ent

import (
	"github.com/golang/protobuf/proto"

	"github.com/jordanorelli/hyperstone/bit"
	"github.com/jordanorelli/hyperstone/dota"
)

func getSerializers(m *dota.CDemoSendTables) (*dota.CSVCMsg_FlattenedSerializer, error) {
	data := m.GetData()
	r := bit.NewBytesReader(data)

	size := int(bit.ReadVarInt(r))
	buf := make([]byte, size)

	r.Read(buf)
	if r.Err() != nil {
		return nil, wrap(r.Err(), "error reading serializers body")
	}

	flat := dota.CSVCMsg_FlattenedSerializer{}
	if err := proto.Unmarshal(buf, &flat); err != nil {
		return nil, wrap(err, "error unmarshaling serializers body")
	}
	return &flat, nil
}
