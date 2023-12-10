package ast

import (
	"strings"

	"github.com/jalopez/go-monkey-interpreter/pkg/token"
)

type CallExpression struct {
	Token     token.Token // The '(' token
	Function  Expression  // Identifier or FunctionLiteral
	Arguments []Expression
}

func (*CallExpression) expressionNode() {} //nolint:golint,unused

// TokenLiteral token literal
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }

// String string representation
func (ce *CallExpression) String() string {
	var out string

	arguments := []string{}
	for _, a := range ce.Arguments {
		arguments = append(arguments, a.String())
	}

	out += ce.Function.String()
	out += "("
	out += strings.Join(arguments, ", ")
	out += ")"

	return out
}
