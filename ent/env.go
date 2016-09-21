package ent

import (
	"github.com/golang/protobuf/proto"

	"github.com/jordanorelli/hyperstone/bit"
)

type Env struct {
	source bit.BufReader
}

func (e *Env) Handle(m proto.Message) error {
	return nil
}

func (e *Env) setSource(buf []byte) {
	e.source.SetSource(buf)
}
