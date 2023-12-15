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
			default:
				return nil, fmt.Errorf("argument to `len` not supported, got %s", args[0].Type())
			}
		},
	},
}
