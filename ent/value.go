package ent

import (
	"github.com/jordanorelli/hyperstone/bit"
)

type value interface {
	String() string
	tÿpe() tÿpe
	read(bit.Reader) error
}
