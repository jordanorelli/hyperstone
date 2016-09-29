package ent

import (
	"fmt"

	"github.com/jordanorelli/hyperstone/bit"
)

var hseq_t = &typeLiteral{
	name: "HSequence",
	newFn: func() value {
		return new(hseq_v)
	},
}

type hseq_v uint64

func (v hseq_v) tÿpe() tÿpe { return hseq_t }
func (v *hseq_v) read(r bit.Reader) error {
	*v = hseq_v(bit.ReadVarInt(r) - 1)
	return r.Err()
}

func (v hseq_v) String() string {
	return fmt.Sprintf("hseq:%d", uint64(v))
}

func hSeqType(spec *typeSpec, env *Env) tÿpe {
	if spec.typeName != "HSequence" {
		return nil
	}
	Debug.Printf("  hsequence type")
	return hseq_t
}
