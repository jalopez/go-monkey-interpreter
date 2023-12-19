package object

// Type object type
type Type string

const (
	// nolint:revive
	INTEGER_OBJ = "INTEGER"
	// nolint:revive
	BOOLEAN_OBJ = "BOOLEAN"
	// nolint:revive
	NULL_OBJ = "NULL"
	// nolint:revive
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	// nolint:revive
	FUNCTION_OBJ = "FUNCTION"
	// nolint:revive
	STRING_OBJ = "STRING"
	// nolint:revive
	ERROR_OBJ = "ERROR"
	// nolint:revive
	BUILTIN_OBJ = "BUILTIN"
	// nolint:revive
	ARRAY_OBJ = "ARRAY"
)

// Object types
type Object interface {
	Type() Type
	Inspect() string
}
