package object

type ReturnValue struct {
	Value Object
}

// Type object type
func (*ReturnValue) Type() Type { return RETURN_VALUE_OBJ }

// Inspect object
func (rv *ReturnValue) Inspect() string { return rv.Value.Inspect() }
