package repl

import (
	"bufio"
	"fmt"
	"io"

	interpreter "github.com/jalopez/go-monkey-interpreter/pkg/eval"
	"github.com/jalopez/go-monkey-interpreter/pkg/eval/object"
	"github.com/jalopez/go-monkey-interpreter/pkg/lexer"
	"github.com/jalopez/go-monkey-interpreter/pkg/parser"
	"github.com/jalopez/go-monkey-interpreter/pkg/token"
)

// PROMPT prompt
const PROMPT = "> "

// Options options
type Options struct {
	Verbose bool
}

// Start starts the REPL
func Start(in io.Reader, out io.Writer, options Options) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		_, err := fmt.Fprintf(out, PROMPT)

		if err != nil {
			panic(err)
		}

		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			if options.Verbose {
				printLexerTokens(out, line)
			}
			printParserErrors(out, p.Errors())
			continue
		}

		result := interpreter.Eval(program, env)
		if result != nil {
			io.WriteString(out, result.Inspect())
			io.WriteString(out, "\n")
		} else {
			io.WriteString(out, "nil\n")
		}

		if options.Verbose {
			io.WriteString(out, "----DEBUG\n")

			printLexerTokens(out, line)

			io.WriteString(out, " AST:\n")

			_, err := fmt.Fprintf(out, "%s\n", program.ToJSON())
			if err != nil {
				panic(err)
			}

		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "Error: "+msg+"\n")
	}
}

func printLexerTokens(out io.Writer, line string) {
	io.WriteString(out, "\n")

	l := lexer.New(line)

	io.WriteString(out, " TOKENS:\n")
	for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		_, err := fmt.Fprintf(out, "%+v\n", tok)
		if err != nil {
			panic(err)
		}
	}
}
