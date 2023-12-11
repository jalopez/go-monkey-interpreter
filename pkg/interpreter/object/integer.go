package object

import "fmt"

// Integer integer
type Integer struct {
	Value int64
}

// Type type
func (*Integer) Type() Type { return INTEGER_OBJ }

// Inspect inspect
func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }
