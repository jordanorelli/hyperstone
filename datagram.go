package main

import (
	"fmt"

	"github.com/golang/protobuf/proto"
)

// datagram represents the top-level envelope in the dota replay format. All
// data in the replay file is packed into datagram frames of at most 65kb.
type dataGram struct {
	cmd  datagramType
	tick int64
	body []byte
}

func (g dataGram) String() string {
	if len(g.body) > 30 {
		return fmt.Sprintf("{dataGram cmd: %v tick: %v size: %d body: %x...}", g.cmd, g.tick, len(g.body), g.body[:27])
	}
	return fmt.Sprintf("{dataGram cmd: %v tick: %v size: %d body: %x}", g.cmd, g.tick, len(g.body), g.body)
}

func (g *dataGram) Open(m *messageFactory) (proto.Message, error) {
	msg, err := m.BuildDatagram(g.cmd)
	if err != nil {
		return nil, err
	}
	if err := proto.Unmarshal(g.body, msg); err != nil {
		return nil, err
	}
	return msg, nil
}
