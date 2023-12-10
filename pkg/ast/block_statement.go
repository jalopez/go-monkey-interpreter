package ast

import "github.com/jalopez/go-monkey-interpreter/pkg/token"

// BlockStatement is a block statement
type BlockStatement struct {
	Token      token.Token // the { token
	Statements []Statement
}

func (*BlockStatement) statementNode() {} //nolint:golint,unused

// TokenLiteral token literal
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }

// String string representation
func (bs *BlockStatement) String() string {
	var out string

	for _, s := range bs.Statements {
		out += s.String()
	}

	return out
}

// ToJSON to json
func (bs *BlockStatement) ToJSON() string {
	var out string

	for _, s := range bs.Statements {
		out += s.ToJSON()
	}

	return `{"type":"block","value":[` + out + `]}`
}
