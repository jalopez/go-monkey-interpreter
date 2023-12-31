package object

import (
	"fmt"
)

// Builtins builtins
var Builtins = []struct {
	Name    string
	Builtin *Builtin
}{
	{
		"len",
		&Builtin{
			Fn: func(args ...Object) (Object, error) {
				if len(args) != 1 {
					return nil, fmt.Errorf("wrong number of arguments. got=%d, want=1", len(args))
				}

				switch arg := args[0].(type) {
				case *String:
					return &Integer{Value: int64(len(arg.Value))}, nil
				case *Array:
					return &Integer{Value: int64(len(arg.Elements))}, nil
				default:
					return nil, fmt.Errorf("argument to `len` not supported, got %s", args[0].Type())
				}
			},
		},
	},
	{
		"first",
		&Builtin{
			Fn: func(args ...Object) (Object, error) {
				if len(args) != 1 {
					return nil, fmt.Errorf("wrong number of arguments. got=%d, want=1", len(args))
				}

				if args[0].Type() != ARRAY_OBJ {
					return nil, fmt.Errorf("argument to `first` must be ARRAY, got %s", args[0].Type())
				}

				arr, ok := args[0].(*Array)

				if !ok {
					return nil, fmt.Errorf("unexpected error")
				}

				if len(arr.Elements) > 0 {
					return arr.Elements[0], nil
				}

				return nil, nil
			},
		},
	},
	{
		"last",
		&Builtin{
			Fn: func(args ...Object) (Object, error) {
				if len(args) != 1 {
					return nil, fmt.Errorf("wrong number of arguments. got=%d, want=1", len(args))
				}

				if args[0].Type() != ARRAY_OBJ {
					return nil, fmt.Errorf("argument to `last` must be ARRAY, got %s", args[0].Type())
				}

				arr, ok := args[0].(*Array)

				if !ok {
					return nil, fmt.Errorf("unexpected error")
				}

				length := len(arr.Elements)

				if length > 0 {
					return arr.Elements[length-1], nil
				}

				return nil, nil
			},
		},
	},
	{
		"rest",
		&Builtin{
			Fn: func(args ...Object) (Object, error) {
				if len(args) != 1 {
					return nil, fmt.Errorf("wrong number of arguments. got=%d, want=1", len(args))
				}

				if args[0].Type() != ARRAY_OBJ {
					return nil, fmt.Errorf("argument to `rest` must be ARRAY, got %s", args[0].Type())
				}

				arr, ok := args[0].(*Array)

				if !ok {
					return nil, fmt.Errorf("unexpected error")
				}

				length := len(arr.Elements)

				if length > 0 {
					newElements := make([]Object, length-1)
					copy(newElements, arr.Elements[1:length])
					return &Array{Elements: newElements}, nil
				}

				return nil, nil
			},
		},
	},
	{
		"push",
		&Builtin{
			Fn: func(args ...Object) (Object, error) {
				if len(args) != 2 {
					return nil, fmt.Errorf("wrong number of arguments. got=%d, want=2", len(args))
				}

				if args[0].Type() != ARRAY_OBJ {
					return nil, fmt.Errorf("argument to `push` must be ARRAY, got %s", args[0].Type())
				}

				arr, ok := args[0].(*Array)

				if !ok {
					return nil, fmt.Errorf("unexpected error")
				}

				length := len(arr.Elements)

				newElements := make([]Object, length+1)
				copy(newElements, arr.Elements)
				newElements[length] = args[1]

				return &Array{Elements: newElements}, nil
			},
		},
	},
	{
		"puts",
		&Builtin{
			Fn: func(args ...Object) (Object, error) {
				for _, arg := range args {
					_, err := fmt.Println(arg.Inspect())
					if err != nil {
						return nil, err
					}
				}
				return nil, nil
			},
		},
	},
}

// GetBuiltinByName get builtin by name
func GetBuiltinByName(name string) *Builtin {
	for _, def := range Builtins {
		if def.Name == name {
			return def.Builtin
		}
	}
	return nil
}
