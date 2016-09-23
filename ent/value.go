package ent

import (
	"github.com/jordanorelli/hyperstone/bit"
)

type value interface {
	read(bit.Reader) error
}
