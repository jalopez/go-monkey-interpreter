package object

// BuiltinFunction builtin function
type BuiltinFunction func(args ...Object) (Object, error)

// Builtin builtin
type Builtin struct {
	Fn BuiltinFunction
}

// Type type
func (*Builtin) Type() Type { return BUILTIN_OBJ }

// Inspect inspect
func (*Builtin) Inspect() string { return "builtin function" }
