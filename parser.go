package main

import (
	"fmt"
	"io"
)

type parser struct {
	// the source of replay bytes. Must NOT be compressed.
	source io.Reader
}

func newParser(r io.Reader) *parser {
	return &parser{source: r}
}

// checks whether we have an acceptable header at the current reader position.
func (p *parser) checkHeader() (bool, error) {
	buf := make([]byte, 8)
	if _, err := p.source.Read(buf); err != nil {
		return false, fmt.Errorf("unable to read header bytes: %v", err)
	}
	return string(buf) == replay_header, nil
}
