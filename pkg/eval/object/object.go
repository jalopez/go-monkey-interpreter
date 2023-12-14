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
	ERROR_OBJ = "ERROR"
)

// Object types
type Object interface {
	Type() Type
	Inspect() string
}
