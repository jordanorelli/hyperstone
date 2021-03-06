package main

import (
	"bufio"
	"fmt"
	"io"
	"sync"

	"github.com/golang/protobuf/proto"
	"github.com/jordanorelli/hyperstone/bit"
	"github.com/jordanorelli/hyperstone/dota"
)

type parser struct {
	// the source of replay bytes. Must NOT be compressed.
	source *bufio.Reader

	scratch []byte
	pbuf    *proto.Buffer
	err     error

	ewl entityWhitelist
	pwl packetWhitelist

	packets *sync.Pool

	br *bit.BufReader
}

func newParser(r io.Reader) *parser {
	return &parser{
		source:  bufio.NewReaderSize(r, 1<<16),
		scratch: make([]byte, 1<<17),
		pbuf:    new(proto.Buffer),
		ewl:     allEntities,
		pwl:     allPackets,
		packets: &sync.Pool{
			New: func() interface{} {
				return &packet{body: make([]byte, 1<<16)}
			},
		},
		br: new(bit.BufReader),
	}
}

func (p *parser) run(out chan maybe) {
	defer close(out)
	ok, err := p.checkHeader()
	if err != nil {
		p.err = wrap(err, "parser start error")
	}
	if !ok {
		p.err = fmt.Errorf("parser start error: invalid header")
	}
	if _, err := p.source.Discard(8); err != nil {
		p.err = wrap(err, "parser start error")
	}
	for {
		pkt, err := p.readPacket()
		if err != nil {
			p.err = wrap(err, "read packet error in run loop")
			return
		}

		msg, err := pkt.Open(&messages, p.pbuf)
		if err != nil {
			p.err = wrap(err, "open packet error in run loop")
			return
		}

		p.packets.Put(pkt)

		switch m := msg.(type) {
		case *dota.CDemoPacket:
			p.emitChildren(m, out)
		}

		out <- maybe{Message: msg}
	}
}

func (p *parser) emitChildren(pkt *dota.CDemoPacket, c chan maybe) {
	p.br.SetSource(pkt.GetData())
	for {
		t := entityType(bit.ReadUBitVar(p.br))
		s := bit.ReadVarInt(p.br)

		if p.ewl[t] {
			p.br.Read(p.scratch[:s])
		} else {
			p.br.DiscardBytes(int(s))
			continue
		}

		switch err := p.br.Err(); err {
		case nil:
			break
		case io.EOF:
			return
		default:
			c <- maybe{error: err}
			return
		}

		p.pbuf.SetBuf(p.scratch[:s])
		e, err := messages.BuildEntity(t)
		if err != nil {
			c <- maybe{error: wrap(err, "skipping entity of size %d, type %s", s, t)}
			continue
		}
		if err := p.pbuf.Unmarshal(e); err != nil {
			c <- maybe{error: wrap(err, "entity unmarshal error")}
			continue
		}
		c <- maybe{Message: e}
	}
}

// DecodeVarint reads a varint-encoded integer from the source reader.
// It returns the value as a uin64 and any errors encountered. The reader will
// be advanced by the number of bytes needed to consume this value. On error,
// the reader will not be advanced.
//
// This is the format for the
// int32, int64, uint32, uint64, bool, and enum
func (p *parser) decodeVarint() (uint64, error) {
	// protobuf defines values that are up to 64 bits wide. The largest value
	// stored in a protobuf varint is 64 data bits, which in varint encoding,
	// would require 10 bytes.
	buf, err := p.source.Peek(10)
	if err != nil {
		return 0, wrap(err, "decode varint couldn't peek 10 bytes")
	}

	var x uint64
	var s uint
	for i, b := range buf {
		// when msb is 0, we're at the last byte of the value
		if b < 0x80 {
			if _, err := p.source.Discard(i + 1); err != nil {
				return 0, wrap(err, "decode varint couldn't discard %d bytes", i)
			}
			return x | uint64(b)<<s, nil
		}

		// otherwise, include the 7 least significant bits in our value
		x |= uint64(b&0x7f) << s
		s += 7
	}

	return 0, fmt.Errorf("decode varint never saw the end of varint")
}

// reads n bytes from the source into the scratch buffer. the returned slice is
// the beginning of the scratch buffer. it will be corrupted on the next call
// to readn or the next operation that utilizes the scratch buffer.
func (p *parser) readn(n int) ([]byte, error) {
	buf := make([]byte, n)
	if _, err := io.ReadFull(p.source, buf); err != nil {
		return nil, wrap(err, "error reading %d bytes", n)
	}
	return buf, nil
}

// checks whether we have an acceptable header at the current reader position.
func (p *parser) checkHeader() (bool, error) {
	buf, err := p.readn(8)
	if err != nil {
		return false, wrap(err, "unable to read header bytes")
	}
	return string(buf) == replayHeader, nil
}

// reads the next packet type indicator off the wire. Also looks for a
// compression flag, returning true if the contents to follow are snappy
// compressed, false otherwise.
func (p *parser) readPacketType() (packetType, bool, error) {
	n, err := p.decodeVarint()
	if err != nil {
		return EDemoCommands_DEM_Error, false, wrap(err, "readPacketType couldn't read varint")
	}

	compressed := false
	if n&0x40 > 0 {
		compressed = true
		n &^= 0x40
	}
	return packetType(n), compressed, nil
}

func (p *parser) readPacket() (*packet, error) {
	cmd, compressed, err := p.readPacketType()
	if err != nil {
		return nil, wrap(err, "readPacket couldn't get a command")
	}

	tick, err := p.decodeVarint()
	if err != nil {
		return nil, wrap(err, "readPacket couldn't read the tick value")
	}

	size, err := p.decodeVarint()
	if err != nil {
		return nil, wrap(err, "readPacket couldn't read the size value")
	}

	if !p.pwl[cmd] {
		if size > 0 {
			_, err := p.source.Discard(int(size))
			if err != nil {
				return nil, wrap(err, "readPacket couldn't skip packet length %d", size)
			}
		}
		return p.readPacket()
	}

	pkt := p.packets.Get().(*packet)
	pkt.cmd = cmd
	pkt.tick = int64(tick)
	pkt.size = int(size)
	pkt.compressed = compressed

	if size > 0 {
		if _, err := io.ReadFull(p.source, pkt.body[:pkt.size]); err != nil {
			return nil, wrap(err, "readPacket couldn't read packet body")
		}
	}

	return pkt, nil
}
