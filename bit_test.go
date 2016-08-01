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

	buf = []byte{0x4}
	bb = newBitBuffer(buf)
	u := bb.readVarUint()
	if u != 1 {
		t.Errorf("feck. wanted %v, got %v", 1, u)
	}
	if bb.readBits(2); bb.err != nil {
		t.Errorf("shouldn't have an error yet")
	}
	if bb.readBits(1); bb.err == nil {
		t.Errorf("we should be at EOF now")
	}

	buf = []byte{0x3c}
	bb = newBitBuffer(buf)
	u = bb.readVarUint()
	if u != 15 {
		t.Errorf("feck. wanted %v, got %v", 15, u)
	}
	if bb.readBits(2); bb.err != nil {
		t.Errorf("shouldn't have an error yet")
	}
	if bb.readBits(1); bb.err == nil {
		t.Errorf("we should be at EOF now")
	}

	buf = []byte{0x48, 0x10}
	// 0100 1000 0001 0000
	// 01                  - prefix bits. indicates length 12
	//   00 10             - least significant four
	//        00 0001 00   - most significant eight
	//                  00 - not read.
	//
	// 0000 0100 0010      - actual value (0x42, or 66)
	bb = newBitBuffer(buf)
	u = bb.readVarUint()
	if u != 66 {
		t.Errorf("feck. wanted %v, got %v", 66, u)
	}
	if bb.readBits(2); bb.err != nil {
		t.Errorf("shouldn't have an error yet")
	}
	if bb.readBits(1); bb.err == nil {
		t.Errorf("we should be at EOF now")
	}
}
