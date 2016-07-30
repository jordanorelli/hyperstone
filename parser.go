package main

import (
	"fmt"
	"io"
)

type parser struct {
	// the source of replay bytes. Must NOT be compressed.
	source io.Reader

	// re-useable scratch buffer. Contents never guaranteed to be clean.
	scratch []byte
}

func newParser(r io.Reader) *parser {
	return &parser{source: r, scratch: make([]byte, 1<<10)}
}

func (p *parser) start() error {
	ok, err := p.checkHeader()
	if err != nil {
		return fmt.Errorf("parser start error: %v", err)
	}
	if !ok {
		return fmt.Errorf("parser start error: invalid header")
	}
	return nil
}

// checks whether we have an acceptable header at the current reader position.
func (p *parser) checkHeader() (bool, error) {
	buf := p.scratch[:8]
	if _, err := p.source.Read(buf); err != nil {
		return false, fmt.Errorf("unable to read header bytes: %v", err)
	}
	return string(buf) == replay_header, nil
}

// skips n bytes in the underlying source
func (p *parser) skip(n int) error {
	if _, err := p.source.Read(p.scratch[:n]); err != nil {
		return fmt.Errorf("unable to skip %d bytes: %v", n, err)
	}
	return nil
}
