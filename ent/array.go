package ent

import (
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
	return array_t{elemType, count}
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

func (t array_t) read(r bit.Reader) (value, error) {
	var err error
	v := make(array, t.count)
	for i := range v {
		v[i], err = t.elem.read(r)
		if err != nil {
			return nil, wrap(err, "array read error at index %d", i)
		}
	}
	return v, r.Err()
}

type array []value
