package ent

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	t_element  = iota // a known element class
	t_object          // c++ object type that is not a known element class
	t_array           // c++ array type
	t_template        // c++ template type
)

// constant identifiers
var cIdents = map[string]int{"MAX_ABILITY_DRAFT_ABILITIES": 48}

type typeSpec struct {
	name     string
	kind     int
	size     int
	template string
	member   *typeSpec
}

func (t typeSpec) String() string {
	if t.member != nil {
		return fmt.Sprintf("{%v %v %v %v %v}", t.name, t.kind, t.size, t.template, *t.member)
	}
	return fmt.Sprintf("{%v %v %v %v %v}", t.name, t.kind, t.size, t.template, t.member)
}

func parseTypeName(n *Namespace, s string) typeSpec {
	s = strings.TrimSpace(s)
	t := typeSpec{name: s}

	if n.HasClass(s) {
		t.kind = t_element
		return t
	}

	// presumably this is some sort of array type
	if strings.ContainsRune(s, '[') {
		memName, count := parseArrayName(s)
		t.kind = t_array
		t.size = count
		t.member = new(typeSpec)
		*t.member = parseTypeName(n, memName)
		return t
	}

	if strings.ContainsRune(s, '<') {
		t.kind = t_template
		template, member := parseTemplateName(s)
		t.template = template
		t.member = new(typeSpec)
		*t.member = parseTypeName(n, member)
		return t
	}

	t.kind = t_object
	return t
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
				n = cIdents[ns]
				if n == 0 {
					panic("invalid array type name: " + err.Error())
				}
			}
			return strings.TrimSpace(string(runes[:i])), n
		}
	}
	panic("invalid array type name: " + s)
}

func parseTemplateName(s string) (string, string) {
	if strings.ContainsRune(s, ',') {
		panic("can't handle templates with multiple parameters")
	}

	runes := []rune(s)
	template := ""
	depth := 0
	bracket_start := -1
	for i, r := range runes {
		if r == '<' {
			if depth == 0 {
				bracket_start = i
				template = strings.TrimSpace(string(runes[:i]))
			}
			depth++
		}
		if r == '>' {
			depth--
			if depth == 0 {
				return template, strings.TrimSpace(string(runes[bracket_start+1 : i]))
			}
		}
	}
	panic("invalid template type definition: " + s)
}
