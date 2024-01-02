package compiler

// SymbolScope is the scope of the symbol.
type SymbolScope string

const (
	// GlobalScope is the global scope.
	GlobalScope SymbolScope = "GLOBAL"
	// LocalScope is the local scope.
	LocalScope SymbolScope = "LOCAL"
	// BuiltinScope is the builtin scope.
	BuiltinScope SymbolScope = "BUILTIN"
	// FreeScope is the free scope.
	FreeScope SymbolScope = "FREE"
	// FunctionScope is the function scope.
	FunctionScope SymbolScope = "FUNCTION"
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

	FreeSymbols []Symbol
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

// DefineBuiltin defines a builtin in the symbol table.
func (s *SymbolTable) DefineBuiltin(index int, name string) Symbol {
	symbol := Symbol{Name: name, Index: index, Scope: BuiltinScope}
	s.store[name] = symbol
	return symbol
}

// DefineFunctionName defines a function name in the symbol table.
func (s *SymbolTable) DefineFunctionName(name string) Symbol {
	s.store[name] = Symbol{Name: name, Index: 0, Scope: FunctionScope}
	return s.store[name]
}

// Resolve resolves a symbol in the symbol table.
func (s *SymbolTable) Resolve(name string) (Symbol, bool) {
	obj, ok := s.store[name]

	if !ok && s.outer != nil {
		obj, ok = s.outer.Resolve(name)
		if !ok {
			return obj, ok
		}

		if obj.Scope == GlobalScope || obj.Scope == BuiltinScope {
			return obj, ok
		}

		free := s.defineFree(obj)
		return free, true
	}

	return obj, ok
}

func (s *SymbolTable) defineFree(original Symbol) Symbol {
	s.FreeSymbols = append(s.FreeSymbols, original)

	symbol := Symbol{
		Name:  original.Name,
		Index: len(s.FreeSymbols) - 1,
		Scope: FreeScope,
	}

	s.store[original.Name] = symbol
	return symbol
}
