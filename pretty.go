package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"

	"github.com/golang/protobuf/proto"
)

func prettySlice(v reflect.Value) string {
	if v.Type().Elem().Kind() == reflect.Uint8 {
		l := v.Len()
		if l > 16 {
			l = 16
		}
		b := make([]byte, l)
		for i := 0; i < l; i++ {
			b[i] = byte(v.Index(i).Uint())
		}
		return fmt.Sprintf("%x", b)
	}

	width := 0
	parts := make([]string, 0, v.Len())
	for i := 0; i < v.Len() && width <= 32; i++ {
		parts = append(parts, pretty(v.Index(i)))
		width += len(parts[i]) // obligatory byte count is not rune count rabble
	}
	return fmt.Sprintf("[%s]", strings.Join(parts, ", "))
}
func prettyPrint(m proto.Message) {
	v := reflect.ValueOf(m)
	fmt.Println(pretty(v))
}
func prettyStruct(v reflect.Value) string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "{%s ", v.Type())
	for fn := 0; fn < v.NumField(); fn++ {
		field := v.Type().Field(fn)
		if field.Name == "XXX_unrecognized" {
			continue
		}
		fv := v.Field(fn)
		fmt.Fprintf(&buf, "%s: %s ", field.Name, pretty(fv))
	}
	fmt.Fprint(&buf, "}")
	return buf.String()
}

func pretty(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Ptr:
		if v.IsNil() {
			return "nil"
		}
		return pretty(v.Elem())
	case reflect.Struct:
		return prettyStruct(v)
	case reflect.Slice:
		return prettySlice(v)
	case reflect.String:
		return fmt.Sprintf("%q", v.String())
	case reflect.Int32:
		return fmt.Sprintf("%d", v.Int())
	case reflect.Uint8, reflect.Uint32:
		return fmt.Sprintf("%d", v.Uint())
	case reflect.Bool:
		return fmt.Sprintf("%t", v.Bool())
	default:
		return v.Type().Name()
	}
}
