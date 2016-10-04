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
	if elemName == "char" {
		return string_t(count)
	}
	return &array_t{elem: elemType, count: count}
}

func parseArrayName(s string) (string, uint) {
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
			return strings.TrimSpace(string(runes[:i])), uint(n)
		}
	}
	panic("invalid array type name: " + s)
}

type array_t struct {
	elem  tÿpe
	count uint
	bits  uint
}

func (t *array_t) sizeBits() uint {
	if t.bits == 0 {
		t.bits = bit.Length(t.count)
	}
	return t.bits
}

func (t *array_t) nü() value       { return array{t: t, slots: make([]value, t.count)} }
func (t array_t) typeName() string { return fmt.Sprintf("%s[%d]", t.elem.typeName(), t.count) }

type array struct {
	t     *array_t
	slots []value
}

func (a array) tÿpe() tÿpe { return a.t }

func (a array) read(r bit.Reader) error {
	n := r.ReadBits(a.t.bits)
	Debug.Printf("reading %d array elements", n)
	for i := uint64(0); i < n; i++ {
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

// ------------------------------------------------------------------------------
// strings are a special case of arrays
// ------------------------------------------------------------------------------

type string_t int

func (t string_t) nü() value {
	return &string_v{t: t, buf: make([]byte, int(t))}
}

func (t string_t) typeName() string { return "string" }

type string_v struct {
	t     string_t
	buf   []byte // the buffer of all possible bytes
	valid []byte // selection of current valid bytes within buf
}

func (s *string_v) tÿpe() tÿpe { return s.t }
func (s *string_v) read(r bit.Reader) error {
	for i := 0; i < int(s.t); i++ {
		b := r.ReadBits(8)
		if b == 0 {
			s.valid = s.buf[:i]
			return r.Err()
		}
		s.buf[i] = byte(b & 0xff)
	}
	s.valid = s.buf
	return r.Err()
}

func (s *string_v) String() string {
	return string(s.valid)
}

func (s *string_v) slotType(int) tÿpe     { return char_t }
func (s *string_v) slotName(n int) string { return strconv.Itoa(n) }
func (s *string_v) setSlotValue(slot int, v value) {
	s.buf[slot] = byte(*v.(*char_v))
	if slot >= len(s.valid) {
		s.valid = s.buf[:slot+1]
	}
}
func (s *string_v) getSlotValue(slot int) value {
	v := char_v(s.buf[slot])
	return &v
}
