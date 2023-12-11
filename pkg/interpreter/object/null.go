package object

// Null null
type Null struct{}

// Type type
func (*Null) Type() Type { return NULL_OBJ }

// Inspect inspect
func (*Null) Inspect() string { return "null" }
