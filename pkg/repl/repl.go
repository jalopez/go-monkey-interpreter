package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/jalopez/go-monkey-interpreter/pkg/interpreter"
	"github.com/jalopez/go-monkey-interpreter/pkg/lexer"
	"github.com/jalopez/go-monkey-interpreter/pkg/parser"
	"github.com/jalopez/go-monkey-interpreter/pkg/token"
)

// PROMPT prompt
const PROMPT = "> "

const monkeyFace = `            __,__
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ 0 | 0 /  /   / |
  \ '- ,\.-"""""""-./, -' /
   ''-' /_   ^ ^   _\ '-''
       |  \._   _./  |
       \   \ '~' /   /
        '._ '-=-' _.'
           '-----'
`

// Options options
type Options struct {
	Verbose bool
}

// Start starts the REPL
func Start(in io.Reader, out io.Writer, options Options) {
	scanner := bufio.NewScanner(in)

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

		result := interpreter.Eval(program)
		if result != nil {
			io.WriteString(out, result.Inspect())
			io.WriteString(out, "\n")
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
	io.WriteString(out, monkeyFace)
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
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
