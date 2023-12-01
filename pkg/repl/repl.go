package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/jalopez/go-monkey-interpreter/pkg/lexer"
	"github.com/jalopez/go-monkey-interpreter/pkg/token"
)

// PROMPT prompt
const PROMPT = "> "

// Start starts the REPL
func Start(in io.Reader, out io.Writer) {
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

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			_, err := fmt.Fprintf(out, "%+v\n", tok)
			if err != nil {
				panic(err)
			}
		}
	}
}
