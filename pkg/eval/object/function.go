package object

import "github.com/jalopez/go-monkey-interpreter/pkg/ast"

// Function function
type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

// Type object type
func (*Function) Type() Type { return FUNCTION_OBJ }

// Inspect object
func (f *Function) Inspect() string {
	out := "fn("
	for _, p := range f.Parameters {
		out += p.String() + ", "
	}
	out += ") {\n"
	out += f.Body.String()
	out += "\n}"
	return out
}
