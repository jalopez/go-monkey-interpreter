package compiler

// SymbolScope is the scope of the symbol.
type SymbolScope string

const (
	// GlobalScope is the global scope.
	GlobalScope SymbolScope = "GLOBAL"
	// LocalScope is the local scope.
	LocalScope SymbolScope = "LOCAL"
)

// Symbol is a symbol in the symbol table.
type Symbol struct {
	Name  string
	Scope SymbolScope
	Index int
}

// SymbolTable is a symbol table.
type SymbolTable struct {
	outer          *SymbolTable
	store          map[string]Symbol
	numDefinitions int
}

// NewSymbolTable creates a new symbol table.
func NewSymbolTable() *SymbolTable {
	s := make(map[string]Symbol)
	return &SymbolTable{store: s}
}

// NewEnclosedSymbolTable creates a new enclosed symbol table.
func NewEnclosedSymbolTable(outer *SymbolTable) *SymbolTable {
	s := NewSymbolTable()
	s.outer = outer
	return s
}

// Define defines a symbol in the symbol table.
func (s *SymbolTable) Define(name string) Symbol {
	symbol := Symbol{Name: name, Index: s.numDefinitions}

	if s.outer == nil {
		symbol.Scope = GlobalScope
	} else {
		symbol.Scope = LocalScope
	}

	s.store[name] = symbol
	s.numDefinitions++
	return symbol
}

// Resolve resolves a symbol in the symbol table.
func (s *SymbolTable) Resolve(name string) (Symbol, bool) {
	obj, ok := s.store[name]

	if !ok && s.outer != nil {
		obj, ok = s.outer.Resolve(name)
		return obj, ok
	}

	return obj, ok
}
