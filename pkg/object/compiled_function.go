package object

import (
	"fmt"

	"github.com/jalopez/go-monkey-interpreter/pkg/code"
)

// CompiledFunction compiled function
type CompiledFunction struct {
	Instructions code.Instructions
}

// Type type
func (*CompiledFunction) Type() Type { return COMPILED_FUNCTION_OBJ }

// Inspect inspect
func (cf *CompiledFunction) Inspect() string { return fmt.Sprintf("CompiledFunction[%p]", cf) }
