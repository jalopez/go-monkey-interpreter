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
)

// Object types
type Object interface {
	Type() Type
	Inspect() string
}
