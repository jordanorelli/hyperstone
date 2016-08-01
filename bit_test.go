package main

import (
	"io"
	"testing"
)

func TestBits(t *testing.T) {
	buf := []byte{0x00}
	bb := newBitBuffer(buf)
	for i := 0; i < 8; i++ {
		if bb.readBits(1) != 0x00 {
			t.Error("hahha what")
		}
		if bb.err != nil {
			t.Errorf("oh weird error: %v", bb.err)
		}
	}
	if bb.readBits(1) != 0x00 {
		t.Error("hahha what")
	}
	if bb.err != io.ErrUnexpectedEOF {
		t.Errorf("oh weird error: %v", bb.err)
	}

	buf = []byte{0x10}
	bb = newBitBuffer(buf)
	if n := bb.readBits(4); n != 0x01 {
		t.Errorf("shit. wanted %v, got %v", 0x01, n)
	}
	if n := bb.readBits(4); n != 0x00 {
		t.Errorf("poop. wanted %v, got %v", 0x00, n)
	}
	if bb.err != nil {
		t.Errorf("fuck")
	}
}
