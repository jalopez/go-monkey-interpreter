package lexer

import "github.com/jalopez/go-monkey-interpreter/pkg/token"

// Lexer lexer
type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

// New creates a new lexer
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// NextToken returns the next token of the input
func (l *Lexer) NextToken() token.Token {
	var nextToken token.Token

	switch l.ch {
	case '=':
		nextToken = newToken(token.ASSIGN, l.ch)
	case ';':
		nextToken = newToken(token.SEMICOLON, l.ch)
	case '(':
		nextToken = newToken(token.LPAREN, l.ch)
	case ')':
		nextToken = newToken(token.RPAREN, l.ch)
	case ',':
		nextToken = newToken(token.COMMA, l.ch)
	case '+':
		nextToken = newToken(token.PLUS, l.ch)
	case '{':
		nextToken = newToken(token.LBRACE, l.ch)
	case '}':
		nextToken = newToken(token.RBRACE, l.ch)
	case 0:
		nextToken.Literal = ""
		nextToken.Type = token.EOF
	}

	l.readChar()
	return nextToken
}

// readChar reads the next character in the input
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ASCII code for "NUL"
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func newToken(tokenType token.Type, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
