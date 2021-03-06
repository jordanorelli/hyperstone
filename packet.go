package main

import (
	"fmt"
	"sync"

	"github.com/golang/protobuf/proto"
	"github.com/golang/snappy"
)

var (
	p_decode_bufs = &sync.Pool{New: func() interface{} { return make([]byte, 1<<18) }}
)

// packet represents the top-level envelope in the dota replay format. All
// data in the replay file is packed into packets of at most 65kb.
type packet struct {
	cmd        packetType
	tick       int64
	compressed bool
	size       int
	body       []byte
}

func (p packet) String() string {
	if len(p.body) > 30 {
		return fmt.Sprintf("{packet cmd: %v tick: %v size: %d body: %x...}", p.cmd, p.tick, p.size, p.body[:27])
	}
	return fmt.Sprintf("{packet cmd: %v tick: %v size: %d body: %x}", p.cmd, p.tick, p.size, p.body)
}

func (p *packet) Open(m *messageFactory, pbuf *proto.Buffer) (proto.Message, error) {
	msg, err := m.BuildPacket(p.cmd)
	if err != nil {
		return nil, err
	}

	if p.compressed {
		buf := p_decode_bufs.Get().([]byte)
		defer p_decode_bufs.Put(buf)
		buf, err := snappy.Decode(buf, p.body[:p.size])
		if err != nil {
			return nil, wrap(err, "packet open failed snappy decode")
		}
		pbuf.SetBuf(buf)
	} else {
		pbuf.SetBuf(p.body[:p.size])
	}

	if err := pbuf.Unmarshal(msg); err != nil {
		return nil, err
	}
	return msg, nil
}
