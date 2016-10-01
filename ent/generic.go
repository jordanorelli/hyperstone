package ent

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jordanorelli/hyperstone/bit"
)

func genericType(spec *typeSpec, env *Env) tÿpe {
	if !strings.Contains(spec.typeName, "<") {
		return nil
	}

	parts, err := genericName(spec.typeName)
	if err != nil {
		return typeError("bad generic name: %v", err)
	}

	genericName, elemName := parts[0], parts[1]
	elemSpec := *spec
	elemSpec.typeName = elemName
	elem := parseTypeSpec(&elemSpec, env)

	switch genericName {
	case "CHandle", "CStrongHandle":
		t := handle_t(fmt.Sprintf("%s<%s>", genericName, elem.typeName()))
		return &t
	case "CUtlVector":
		return &cutl_vector_t{elem}
	default:
		return typeError("unknown generic name: %v", parts[0])
	}
}

func genericName(name string) ([2]string, error) {
	var out [2]string
	runes := []rune(strings.TrimSpace(name))
	b_start := 0
	depth := 0
	for i, r := range runes {
		if r == '<' {
			if depth == 0 {
				b_start = i
				out[0] = strings.TrimSpace(string(runes[0:i]))
			}
			depth++
		}
		if r == '>' {
			depth--
			if depth == 0 {
				if i == len(runes)-1 {
					out[1] = strings.TrimSpace(string(runes[b_start+1 : i]))
				} else {
					return out, fmt.Errorf("extra runes in generic type name")
				}
			}
		}
	}
	if out[0] == "" {
		return out, fmt.Errorf("empty generic container name")
	}
	if out[1] == "" {
		return out, fmt.Errorf("empty generic element name")
	}
	Debug.Printf("  generic name in: %s out: %v", name, out)
	return out, nil
}

type cutl_vector_t struct {
	elem tÿpe
}

func (t *cutl_vector_t) nü() value {
	return &cutl_vector{t: t}
}

func (t cutl_vector_t) typeName() string {
	return fmt.Sprintf("CUtlVector<%s>", t.elem.typeName())
}

type cutl_vector struct {
	t     *cutl_vector_t
	slots []value
}

func (v *cutl_vector) tÿpe() tÿpe { return v.t }

func (v *cutl_vector) read(r bit.Reader) error {
	count := bit.ReadVarInt32(r)
	Debug.Printf("allocating cutl_vector of size %d with element type %s", count, v.t.elem.typeName())
	v.slots = make([]value, count)
	return r.Err()
}

func (v cutl_vector) String() string {
	if len(v.slots) > 8 {
		return fmt.Sprintf("%s(%d)%v...", v.t.typeName(), len(v.slots), v.slots[:8])
	}
	return fmt.Sprintf("%s(%d)%v", v.t.typeName(), len(v.slots), v.slots)
}

func (v *cutl_vector) slotType(int) tÿpe     { return v.t.elem }
func (v *cutl_vector) slotName(n int) string { return strconv.Itoa(n) }
func (v *cutl_vector) setSlotValue(slot int, val value) {
	if slot == len(v.slots) {
		v.slots = append(v.slots, val)
	} else {
		v.slots[slot] = val
	}
}
func (v *cutl_vector) getSlotValue(slot int) value { return v.slots[slot] }
