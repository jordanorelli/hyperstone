package ent

type symbol struct {
	id    int
	table *symbolTable
}

func (s symbol) String() string { return (*s.table)[s.id] }

type symbolTable []string
