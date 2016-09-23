package ent

type tÿpe struct {
	name  string
	alloc func(...interface{}) value
}
