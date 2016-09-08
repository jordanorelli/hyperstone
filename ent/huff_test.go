package ent

import (
	"bytes"
	"testing"
)

func TestDump(t *testing.T) {
	t.Log(hlist)

	var buf bytes.Buffer
	dump(htree, "", &buf)
	t.Logf("%s", buf.String())
}
