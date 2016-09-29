package ent

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jordanorelli/hyperstone/bit"
)

var constants = map[string]int{
	"MAX_ABILITY_DRAFT_ABILITIES": 48,
}

func arrayType(spec *typeSpec, env *Env) tÿpe {
	if !strings.Contains(spec.typeName, "[") {
		return nil
	}
	elemName, count := parseArrayName(spec.typeName)
	elemSpec := *spec
	elemSpec.typeName = elemName
	elemType := parseTypeSpec(&elemSpec, env)
	return &array_t{elemType, count}
}

func parseArrayName(s string) (string, int) {
	runes := []rune(s)
	if runes[len(runes)-1] != ']' {
		panic("invalid array type name: " + s)
	}
	for i := len(runes) - 2; i >= 0; i-- {
		if runes[i] == '[' {
			ns := strings.TrimSpace(string(runes[i+1 : len(runes)-1]))
			n, err := strconv.Atoi(ns)
			if err != nil {
				n = constants[ns]
				if n <= 0 {
					panic("invalid array type name: " + err.Error())
				}
			}
			return strings.TrimSpace(string(runes[:i])), n
		}
	}
	panic("invalid array type name: " + s)
}

type array_t struct {
	elem  tÿpe
	count int
}

func (t *array_t) nü() value       { return array{t: t, slots: make([]value, t.count)} }
func (t array_t) typeName() string { return fmt.Sprintf("%s[%d]", t.elem.typeName(), t.count) }

type array struct {
	t     *array_t
	slots []value
}

func (a array) tÿpe() tÿpe { return a.t }

func (a array) read(r bit.Reader) error {
	for i := range a.slots {
		if a.slots[i] == nil {
			a.slots[i] = a.t.elem.nü()
		}
		if err := a.slots[i].read(r); err != nil {
			return wrap(err, "array read error at index %d", i)
		}
	}
	return r.Err()
}

func (a array) String() string {
	if len(a.slots) > 8 {
		return fmt.Sprintf("%s(%d)%v...", a.t.typeName(), len(a.slots), a.slots[:8])
	}
	return fmt.Sprintf("%s(%d)%v", a.t.typeName(), len(a.slots), a.slots)
}

func (a array) slotType(int) tÿpe     { return a.t.elem }
func (a array) slotName(n int) string { return strconv.Itoa(n) }
func (a array) setSlotValue(slot int, v value) {
	// TODO: type check here?
	a.slots[slot] = v
}
func (a array) getSlotValue(slot int) value { return a.slots[slot] }
