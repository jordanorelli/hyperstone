package main

import (
	"bufio"
	"fmt"
	"io"
	// "github.com/golang/protobuf/proto"
)

type parser struct {
	// the source of replay bytes. Must NOT be compressed.
	source *bufio.Reader

	// re-useable scratch buffer. Contents never guaranteed to be clean.
	scratch []byte
}

func newParser(r io.Reader) *parser {
	br := bufio.NewReader(r)
	return &parser{source: br, scratch: make([]byte, 1<<10)}
}

func (p *parser) start() error {
	ok, err := p.checkHeader()
	if err != nil {
		return fmt.Errorf("parser start error: %v", err)
	}
	if !ok {
		return fmt.Errorf("parser start error: invalid header")
	}
	if _, err := p.source.Discard(8); err != nil {
		return err
	}
	return nil
}

func (p *parser) run() {
	for {
		msg, err := p.readMessage()
		if err != nil {
			fmt.Printf("error: %v\n", err)
			return
		}
		fmt.Println(msg)
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
		return 0, fmt.Errorf("decode varint couldn't peek 9 bytes: %v", err)
	}

	var x uint64
	var s uint
	for i, b := range buf {
		// when msb is 0, we're at the last byte of the value
		if b < 0x80 {
			if _, err := p.source.Discard(i + 1); err != nil {
				return 0, fmt.Errorf("decode varint couldn't discard %d bytes: %v", i, err)
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
	if n > cap(p.scratch) {
		p.scratch = make([]byte, 2*cap(p.scratch))
		return p.readn(n)
	}
	buf := p.scratch[:n]
	if _, err := io.ReadFull(p.source, buf); err != nil {
		return nil, fmt.Errorf("error reading %d bytes: %v", n, err)
	}
	return buf, nil
}

// checks whether we have an acceptable header at the current reader position.
func (p *parser) checkHeader() (bool, error) {
	buf, err := p.readn(8)
	if err != nil {
		return false, fmt.Errorf("unable to read header bytes: %v", err)
	}
	return string(buf) == replayHeader, nil
}

func (p *parser) readCommand() (EDemoCommands, bool, error) {
	n, err := p.decodeVarint()
	if err != nil {
		return EDemoCommands_DEM_Error, false, fmt.Errorf("readCommand couldn't read varint: %v", err)
	}

	compressed := false
	if n&0x40 > 0 {
		compressed = true
		n &^= 0x40
	}
	return EDemoCommands(n), compressed, nil
}

type message struct {
	cmd        EDemoCommands
	tick       int64
	compressed bool
	body       []byte
}

func (m *message) String() string {
	if len(m.body) > 30 {
		return fmt.Sprintf("{cmd: %v tick: %v compressed: %t body(%d): %q...}", m.cmd, m.tick, m.compressed, len(m.body), m.body[:27])
	}
	return fmt.Sprintf("{cmd: %v tick: %v compressed: %t body(%d): %q}", m.cmd, m.tick, m.compressed, len(m.body), m.body)
}

func (p *parser) readMessage() (*message, error) {
	cmd, compressed, err := p.readCommand()
	if err != nil {
		return nil, fmt.Errorf("readMessage couldn't get a command: %v", err)
	}

	tick, err := p.decodeVarint()
	if err != nil {
		return nil, fmt.Errorf("readMessage couldn't read the tick value: %v", err)
	}

	size, err := p.decodeVarint()
	if err != nil {
		return nil, fmt.Errorf("readMessage couldn't read the size value: %v", err)
	}

	if size > 0 {
		buf, err := p.readn(int(size))
		if err != nil {
			return nil, fmt.Errorf("readMessage couldn't read message body: %v", err)
		}
		return &message{cmd, int64(tick), compressed, buf}, nil
	}

	return &message{cmd, int64(tick), compressed, nil}, nil
}
