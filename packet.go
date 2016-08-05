package main

import (
	"fmt"

	"github.com/golang/protobuf/proto"
)

// packet represents the top-level envelope in the dota replay format. All
// data in the replay file is packed into packets of at most 65kb.
type packet struct {
	cmd  packetType
	tick int64
	body []byte
}

func (p packet) String() string {
	if len(p.body) > 30 {
		return fmt.Sprintf("{packet cmd: %v tick: %v size: %d body: %x...}", p.cmd, p.tick, len(p.body), p.body[:27])
	}
	return fmt.Sprintf("{packet cmd: %v tick: %v size: %d body: %x}", p.cmd, p.tick, len(p.body), p.body)
}

func (p *packet) Open(m *messageFactory, pbuf *proto.Buffer) (proto.Message, error) {
	msg, err := m.BuildPacket(p.cmd)
	if err != nil {
		return nil, err
	}
	pbuf.SetBuf(p.body)
	if err := pbuf.Unmarshal(msg); err != nil {
		return nil, err
	}
	return msg, nil
}
