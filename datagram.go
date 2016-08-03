package main

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/golang/snappy"
	"github.com/jordanorelli/hyperstone/bit"
	"github.com/jordanorelli/hyperstone/dota"
)

// datagram represents the top-level envelope in the dota replay format. All
// data in the replay file is packed into datagram frames of at most 65kb.
type dataGram struct {
	cmd        dota.EDemoCommands
	tick       int64
	compressed bool
	body       []byte
}

func (g dataGram) String() string {
	if len(g.body) > 30 {
		return fmt.Sprintf("{dataGram cmd: %v tick: %v compressed: %t size: %d body: %q...}", g.cmd, g.tick, g.compressed, len(g.body), g.body[:27])
	}
	return fmt.Sprintf("{dataGram cmd: %v tick: %v compressed: %t size: %d body: %q}", g.cmd, g.tick, g.compressed, len(g.body), g.body)
}

func (g *dataGram) check(dump bool) error {
	if g.cmd != dota.EDemoCommands_DEM_Packet {
		return fmt.Errorf("wrong command type in openPacket: %v", g.cmd)
	}

	if g.compressed {
		buf, err := snappy.Decode(nil, g.body)
		if err != nil {
			return wrap(err, "open packet error: could not decode body")
		}
		g.body = buf
		g.compressed = false
	}

	packet := new(dota.CDemoPacket)
	if err := proto.Unmarshal(g.body, packet); err != nil {
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
			fmt.Printf("\t%v\n", entity{t: uint32(t), size: uint32(s), body: b})
			e := entFactory.BuildMessage(int(t))
			if e == nil {
				fmt.Printf("\tno known entity for type id %d\n", int(t))
				continue
			}
			err := proto.Unmarshal(b, e)
			if err != nil {
				fmt.Printf("entity unmarshal error: %v\n", err)
			} else {
				fmt.Printf("\t\t%v\n", e)
			}
		}
	}
	return nil
}
