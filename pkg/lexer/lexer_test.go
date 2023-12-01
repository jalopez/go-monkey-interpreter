package lexer

import (
	"testing"

	"github.com/jalopez/go-monkey-interpreter/pkg/token"
)

type TokenResult struct {
	expectedType    token.Type
	expectedLiteral string
}

func runTest(t *testing.T, input string, tokens []TokenResult) {
	l := New(input)

	for i, tt := range tokens {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextToken_simpleTokens(t *testing.T) {
	input := `=+(){},;`
	expected := []TokenResult{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	runTest(t, input, expected)
}

func TestNextToken_validIdentifiers(t *testing.T) {
	input := `variable with_snake_case andCamelCase and_1_2_3`
	expected := []TokenResult{
		{token.IDENT, "variable"},
		{token.IDENT, "with_snake_case"},
		{token.IDENT, "andCamelCase"},
		{token.IDENT, "and_1_2_3"},
		{token.EOF, ""},
	}

	runTest(t, input, expected)
}

func TestNextToken_illegalIdentifiers(t *testing.T) {
	input := `1variable 1_variable 1variable 1_2_3 123`
	expected := []TokenResult{
		{token.ILLEGAL, "1variable"},
		{token.ILLEGAL, "1_variable"},
		{token.ILLEGAL, "1variable"},
		{token.ILLEGAL, "1_2_3"},
		{token.INT, "123"},
		{token.EOF, ""},
	}

	runTest(t, input, expected)
}

func TestNextToken_fullProgram(t *testing.T) {
	input := `let five = 5;
let ten = 10;

let add = fn(x, y) {
x + y;
};

let result = add(five, ten);
`
	expected := []TokenResult{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	runTest(t, input, expected)
}
