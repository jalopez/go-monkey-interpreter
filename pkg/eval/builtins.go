package eval

import (
	"fmt"

	"github.com/jalopez/go-monkey-interpreter/pkg/eval/object"
)

var builtins = map[string]*object.Builtin{
	"len": {
		Fn: func(args ...object.Object) (object.Object, error) {
			if len(args) != 1 {
				return nil, fmt.Errorf("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}, nil
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}, nil
			default:
				return nil, fmt.Errorf("argument to `len` not supported, got %s", args[0].Type())
			}
		},
	},
	"first": {
		Fn: func(args ...object.Object) (object.Object, error) {
			if len(args) != 1 {
				return nil, fmt.Errorf("wrong number of arguments. got=%d, want=1", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return nil, fmt.Errorf("argument to `first` must be ARRAY, got %s", args[0].Type())
			}

			arr, ok := args[0].(*object.Array)

			if !ok {
				return nil, fmt.Errorf("unexpected error")
			}

			if len(arr.Elements) > 0 {
				return arr.Elements[0], nil
			}

			return NULL, nil
		},
	},
	"last": {
		Fn: func(args ...object.Object) (object.Object, error) {
			if len(args) != 1 {
				return nil, fmt.Errorf("wrong number of arguments. got=%d, want=1", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return nil, fmt.Errorf("argument to `last` must be ARRAY, got %s", args[0].Type())
			}

			arr, ok := args[0].(*object.Array)

			if !ok {
				return nil, fmt.Errorf("unexpected error")
			}

			length := len(arr.Elements)

			if length > 0 {
				return arr.Elements[length-1], nil
			}

			return NULL, nil
		},
	},
	"rest": {
		Fn: func(args ...object.Object) (object.Object, error) {
			if len(args) != 1 {
				return nil, fmt.Errorf("wrong number of arguments. got=%d, want=1", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return nil, fmt.Errorf("argument to `rest` must be ARRAY, got %s", args[0].Type())
			}

			arr, ok := args[0].(*object.Array)

			if !ok {
				return nil, fmt.Errorf("unexpected error")
			}

			length := len(arr.Elements)

			if length > 0 {
				newElements := make([]object.Object, length-1)
				copy(newElements, arr.Elements[1:length])
				return &object.Array{Elements: newElements}, nil
			}

			return NULL, nil
		},
	},
	"push": {
		Fn: func(args ...object.Object) (object.Object, error) {
			if len(args) != 2 {
				return nil, fmt.Errorf("wrong number of arguments. got=%d, want=2", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return nil, fmt.Errorf("argument to `push` must be ARRAY, got %s", args[0].Type())
			}

			arr, ok := args[0].(*object.Array)

			if !ok {
				return nil, fmt.Errorf("unexpected error")
			}

			length := len(arr.Elements)

			newElements := make([]object.Object, length+1)
			copy(newElements, arr.Elements)
			newElements[length] = args[1]

			return &object.Array{Elements: newElements}, nil
		},
	},
}
