package main

import (
	"github.com/golang/protobuf/proto"
)

// either a value or an error
type maybe struct {
	proto.Message
	error
}
