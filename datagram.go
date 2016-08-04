package main

import (
	"fmt"
	"io"

	"github.com/golang/protobuf/proto"
	"github.com/jordanorelli/hyperstone/bit"
	"github.com/jordanorelli/hyperstone/dota"
)

// datagram represents the top-level envelope in the dota replay format. All
// data in the replay file is packed into datagram frames of at most 65kb.
type dataGram struct {
	cmd  dota.EDemoCommands
	tick int64
	body []byte
}

func (g dataGram) String() string {
	if len(g.body) > 30 {
		return fmt.Sprintf("{dataGram cmd: %v tick: %v size: %d body: %x...}", g.cmd, g.tick, len(g.body), g.body[:27])
	}
	return fmt.Sprintf("{dataGram cmd: %v tick: %v size: %d body: %x}", g.cmd, g.tick, len(g.body), g.body)
}

func (g *dataGram) check(dump bool) error {
	if g.cmd != dota.EDemoCommands_DEM_Packet {
		return fmt.Errorf("wrong command type in openPacket: %v", g.cmd)
	}

	packet := new(dota.CDemoPacket)
	if err := proto.Unmarshal(g.body, packet); err != nil {
		return wrap(err, "onPacket unable to unmarshal message body")
	}

	br := bit.NewBytesReader(packet.GetData())
	for {
		t := br.ReadUBitVar()
		s := br.ReadVarInt()
		b := make([]byte, s)
		br.Read(b)
		switch err := br.Err(); err {
		case nil:
			break
		case io.EOF:
			return nil
		default:
			return err
		}
		if dump {
			fmt.Printf("\t%v\n", entity{t: uint32(t), size: uint32(s), body: b})
		}
		e := entFactory.BuildMessage(int(t))
		if e == nil {
			fmt.Printf("\tno known entity for type id %d size: %d\n", int(t), len(b))
			continue
		}
		err := proto.Unmarshal(b, e)
		if err != nil {
			fmt.Printf("entity unmarshal error: %v\n", err)
		}
	}
	return nil
}
