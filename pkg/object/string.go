package object

// String string
type String struct {
	Value string
}

// Type type
func (*String) Type() Type { return STRING_OBJ }

// Inspect inspect
func (s *String) Inspect() string { return s.Value }
