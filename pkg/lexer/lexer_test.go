package lexer

import (
	"testing"

	"github.com/jalopez/go-monkey-interpreter/pkg/token"
)

type TokenResult struct {
	expectedType    token.Type
	expectedLiteral string
	expectedLine    int
	expectedColumn  int
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
	input := `= + ( ) { } , ; - / * ! < > == !=`
	expected := []TokenResult{
		{token.ASSIGN, "=", 1, 1},
		{token.PLUS, "+", 1, 3},
		{token.LPAREN, "(", 1, 5},
		{token.RPAREN, ")", 1, 7},
		{token.LBRACE, "{", 1, 9},
		{token.RBRACE, "}", 1, 11},
		{token.COMMA, ",", 1, 13},
		{token.SEMICOLON, ";", 1, 15},
		{token.MINUS, "-", 1, 17},
		{token.SLASH, "/", 1, 19},
		{token.ASTERISK, "*", 1, 21},
		{token.BANG, "!", 1, 23},
		{token.LT, "<", 1, 25},
		{token.GT, ">", 1, 27},
		{token.EQ, "==", 1, 29},
		{token.NOTEQ, "!=", 1, 32},
		{token.EOF, "", 1, 35},
	}

	runTest(t, input, expected)
}

func TestNextToken_validIdentifiers(t *testing.T) {
	input := `variable with_snake_case andCamelCase and_1_2_3`
	expected := []TokenResult{
		{token.IDENT, "variable", 1, 1},
		{token.IDENT, "with_snake_case", 1, 10},
		{token.IDENT, "andCamelCase", 1, 26},
		{token.IDENT, "and_1_2_3", 1, 39},
		{token.EOF, "", 1, 48},
	}

	runTest(t, input, expected)
}

func TestNextToken_illegalIdentifiers(t *testing.T) {
	input := `1variable 1_variable 1variable 1_2_3 123`
	expected := []TokenResult{
		{token.ILLEGAL, "1variable", 1, 1},
		{token.ILLEGAL, "1_variable", 1, 10},
		{token.ILLEGAL, "1variable", 1, 21},
		{token.ILLEGAL, "1_2_3", 1, 30},
		{token.INT, "123", 1, 37},
		{token.EOF, "", 1, 40},
	}

	runTest(t, input, expected)
}

func TestNextToken_strings(t *testing.T) {
	input := `"hello" "hello world"`
	expected := []TokenResult{
		{token.STRING, "hello", 1, 1},
		{token.STRING, "hello world", 1, 9},
		{token.EOF, "", 1, 35},
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

if (5 < 10) {
	return true;
} else {
	return false;
}
`
	expected := []TokenResult{
		{token.LET, "let", 1, 1},
		{token.IDENT, "five", 1, 5},
		{token.ASSIGN, "=", 1, 10},
		{token.INT, "5", 1, 12},
		{token.SEMICOLON, ";", 1, 13},
		{token.LET, "let", 2, 1},
		{token.IDENT, "ten", 2, 5},
		{token.ASSIGN, "=", 2, 9},
		{token.INT, "10", 2, 11},
		{token.SEMICOLON, ";", 2, 13},
		{token.LET, "let", 4, 1},
		{token.IDENT, "add", 4, 5},
		{token.ASSIGN, "=", 4, 9},
		{token.FUNCTION, "fn", 4, 11},
		{token.LPAREN, "(", 4, 13},
		{token.IDENT, "x", 4, 14},
		{token.COMMA, ",", 4, 15},
		{token.IDENT, "y", 4, 17},
		{token.RPAREN, ")", 4, 18},
		{token.LBRACE, "{", 4, 20},
		{token.IDENT, "x", 5, 2},
		{token.PLUS, "+", 5, 4},
		{token.IDENT, "y", 5, 6},
		{token.SEMICOLON, ";", 5, 7},
		{token.RBRACE, "}", 6, 1},
		{token.SEMICOLON, ";", 6, 2},
		{token.LET, "let", 8, 1},
		{token.IDENT, "result", 8, 5},
		{token.ASSIGN, "=", 8, 12},
		{token.IDENT, "add", 8, 14},
		{token.LPAREN, "(", 8, 17},
		{token.IDENT, "five", 8, 18},
		{token.COMMA, ",", 8, 22},
		{token.IDENT, "ten", 8, 24},
		{token.RPAREN, ")", 8, 27},
		{token.SEMICOLON, ";", 8, 28},
		{token.IF, "if", 10, 1},
		{token.LPAREN, "(", 10, 4},
		{token.INT, "5", 10, 5},
		{token.LT, "<", 10, 7},
		{token.INT, "10", 10, 9},
		{token.RPAREN, ")", 10, 11},
		{token.LBRACE, "{", 10, 13},
		{token.RETURN, "return", 11, 2},
		{token.TRUE, "true", 11, 9},
		{token.SEMICOLON, ";", 11, 13},
		{token.RBRACE, "}", 12, 1},
		{token.ELSE, "else", 12, 3},
		{token.LBRACE, "{", 12, 8},
		{token.RETURN, "return", 13, 2},
		{token.FALSE, "false", 13, 9},
		{token.SEMICOLON, ";", 13, 14},
		{token.RBRACE, "}", 14, 1},
		{token.EOF, "", 14, 2},
	}

	runTest(t, input, expected)
}
