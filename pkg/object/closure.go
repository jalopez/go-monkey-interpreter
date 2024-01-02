package object

import "fmt"

// Closure represents a function object.
type Closure struct {
	Fn   *CompiledFunction
	Free []Object
}

// Type returns the type of the object.
func (*Closure) Type() Type { return CLOSURE_OBJ }

// Inspect returns a stringified version of the object.
func (c *Closure) Inspect() string {
	return fmt.Sprintf("Closure[%p]", c)
}
