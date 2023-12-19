package object

// Array array
type Array struct {
	Elements []Object
}

// Type type
func (*Array) Type() Type { return ARRAY_OBJ }

// Inspect inspect
func (ao *Array) Inspect() string {
	out := "["
	for _, e := range ao.Elements {
		out += e.Inspect()
		out += ","
	}
	out = out[:len(out)-1]
	out += "]"
	return out
}
