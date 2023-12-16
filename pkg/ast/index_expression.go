package ast

import "github.com/jalopez/go-monkey-interpreter/pkg/token"

// IndexExpression literal with index value
type IndexExpression struct {
	Token token.Token // The '[' token
	Left  Expression
	Index Expression
}

func (*IndexExpression) expressionNode() {} //nolint:golint,unused

// TokenLiteral token literal
func (ie *IndexExpression) TokenLiteral() string { return ie.Token.Literal }

// String string representation
func (ie *IndexExpression) String() string {
	return "(" + ie.Left.String() + "[" + ie.Index.String() + "]" + ")"
}

// ToJSON to json
func (ie *IndexExpression) ToJSON() string {
	return `{"type":"index", "left":` + ie.Left.ToJSON() + `, "index":` + ie.Index.ToJSON() + `}`
}
