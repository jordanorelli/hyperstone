package dt

// the internal representation of table data refers to all labels as
// interned strings (symbols). This array of string contains the mapping of
// symbol ids to symbol display representations. The sample replay I have
// at the time of writing this contains 2215 symbols in its symbol table.
// The dota replay format uses an ordered list of symbols.
type SymbolTable []string

func (t *SymbolTable) Symbol(id int) Symbol { return Symbol{id: id, table: t} }

type Symbol struct {
	id    int
	table *SymbolTable
}

func (s Symbol) String() string { return (*s.table)[s.id] }
