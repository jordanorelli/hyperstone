package main

import (
	"compress/bzip2"
	"io"
)

// bzipCloser wraps an io.ReadCloser in a bzip reader, while allowing a caller
// to close the underlying source
type bzipCloser struct {
	reader io.Reader
	source io.ReadCloser
}

func newBzipCloser(r io.ReadCloser) io.ReadCloser {
	return &bzipCloser{reader: bzip2.NewReader(r), source: r}
}

func (b *bzipCloser) Read(p []byte) (int, error) { return b.reader.Read(p) }

func (b *bzipCloser) Close() error { return b.source.Close() }
