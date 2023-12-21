package compiler

// SymbolScope is the scope of the symbol.
type SymbolScope string

const (
	// GlobalScope is the global scope.
	GlobalScope SymbolScope = "GLOBAL"
)

// Symbol is a symbol in the symbol table.
type Symbol struct {
	Name  string
	Scope SymbolScope
	Index int
}

// SymbolTable is a symbol table.
type SymbolTable struct {
	store          map[string]Symbol
	numDefinitions int
}

// NewSymbolTable creates a new symbol table.
func NewSymbolTable() *SymbolTable {
	s := make(map[string]Symbol)
	return &SymbolTable{store: s}
}

// Define defines a symbol in the symbol table.
func (s *SymbolTable) Define(name string) Symbol {
	symbol := Symbol{Name: name, Index: s.numDefinitions, Scope: GlobalScope}
	s.store[name] = symbol
	s.numDefinitions++
	return symbol
}

// Resolve resolves a symbol in the symbol table.
func (s *SymbolTable) Resolve(name string) (Symbol, bool) {
	obj, ok := s.store[name]
	return obj, ok
}
