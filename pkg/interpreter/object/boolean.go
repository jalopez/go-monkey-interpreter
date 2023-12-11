package object

import "fmt"

// Boolean boolean
type Boolean struct {
	Value bool
}

// Type type
func (*Boolean) Type() Type { return BOOLEAN_OBJ }

// Inspect inspect
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }
