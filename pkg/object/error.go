package object

import "fmt"

// Error object
type Error struct {
	Message string
	Line    int
	Column  int
}

// Type object type
func (*Error) Type() Type { return ERROR_OBJ }

// Inspect object
func (e *Error) Inspect() string {
	if e.Line == 0 {
		return fmt.Sprintf("Error: %s", e.Message)
	}
	return fmt.Sprintf("Error: %s at line %d column %d", e.Message, e.Line, e.Column)
}
