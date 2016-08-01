package main

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/golang/snappy"
	"github.com/jordanorelli/hyperstone/dota"
)

type message struct {
	cmd        dota.EDemoCommands
	tick       int64
	compressed bool
	body       []byte
}

func (m message) String() string {
	if len(m.body) > 30 {
		return fmt.Sprintf("{cmd: %v tick: %v compressed: %t size: %d body: %q...}", m.cmd, m.tick, m.compressed, len(m.body), m.body[:27])
	}
	return fmt.Sprintf("{cmd: %v tick: %v compressed: %t size: %d body: %q}", m.cmd, m.tick, m.compressed, len(m.body), m.body)
}

func (m *message) check(dump bool) error {
	if m.cmd != dota.EDemoCommands_DEM_Packet {
		return fmt.Errorf("wrong command type in openPacket: %v", m.cmd)
	}

	if m.compressed {
		buf, err := snappy.Decode(nil, m.body)
		if err != nil {
			return wrap(err, "open packet error: could not decode body")
		}
		m.body = buf
		m.compressed = false
	}

	packet := new(dota.CDemoPacket)
	if err := proto.Unmarshal(m.body, packet); err != nil {
		return wrap(err, "onPacket unable to unmarshal message body")
	}

	if dump {
		shit := packet.GetData()[:4]
		bb := newBitBuffer(packet.GetData())
		type T struct {
			t int32
		}
		var v T
		v.t = int32(bb.readVarUint())
		if bb.err != nil {
			fmt.Printf("packet error: %v\n", bb.err)
		} else {
			fmt.Printf("{in: %d out: %d data: %v shit: %x}\n", packet.GetSequenceIn(), packet.GetSequenceOutAck(), v, shit)
		}
	}
	return nil
}
