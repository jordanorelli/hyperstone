package main

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/golang/snappy"
	"github.com/jordanorelli/hyperstone/bit"
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

type entity struct {
	t    uint32
	size uint32
	body []byte
}

func (e entity) String() string {
	if len(e.body) < 32 {
		return fmt.Sprintf("{%v %v %x}", e.t, e.size, e.body)
	}
	return fmt.Sprintf("{%v %v %x...}", e.t, e.size, e.body[:32])
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
		br := bit.NewBytesReader(packet.GetData())
		for {
			t := br.ReadUBitVar()
			s := br.ReadVarInt()
			b := make([]byte, s)
			br.Read(b)
			if br.Err() != nil {
				break
			}
			e := entity{t: uint32(t), size: uint32(s), body: b}
			fmt.Printf("\t%v\n", e)
		}
	}
	return nil
}
