package ent

import (
	"fmt"
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

	// genericName, elemName := parts[0], parts[1]
	// elemType := func() tÿpe {
	// 	var elemSpec typeSpec
	// 	elemSpec = *spec
	// 	elemSpec.name = elemName
	// 	return parseTypeSpec(&elemSpec, env)
	// }

	switch parts[0] {
	case "CStrongHandle":
		return typeFn(func(r bit.Reader) (value, error) {
			return handle(bit.ReadVarInt(r)), r.Err()
		})
	default:
		return typeError("unknown generic name: %v", genericName)
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
				out[0] = strings.TrimSpace(string(runes[0:i]))
			}
			depth++
			b_start = i
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
	return out, nil
}
