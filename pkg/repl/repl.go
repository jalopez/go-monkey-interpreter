package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/jalopez/go-monkey-interpreter/pkg/compiler"
	interpreter "github.com/jalopez/go-monkey-interpreter/pkg/eval"
	"github.com/jalopez/go-monkey-interpreter/pkg/lexer"
	"github.com/jalopez/go-monkey-interpreter/pkg/object"
	"github.com/jalopez/go-monkey-interpreter/pkg/parser"
	"github.com/jalopez/go-monkey-interpreter/pkg/token"
	"github.com/jalopez/go-monkey-interpreter/pkg/vm"
)

// PROMPT prompt
const PROMPT = "> "

// Options options
type Options struct {
	Verbose        bool
	CompileEnabled bool
}

// Start starts the REPL
func Start(in io.Reader, out io.Writer, options Options) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	constants := []object.Object{}
	globals := make([]object.Object, vm.GlobalsSize)
	symbolTable := compiler.NewSymbolTable()

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

		if options.CompileEnabled {
			comp := compiler.NewWithState(symbolTable, constants)
			err := comp.Compile(program)
			if err != nil {
				fmt.Fprintf(out, "Compilation failed:\n %s\n", err)
				continue
			}

			code := comp.Bytecode()
			constants = code.Constants

			machine := vm.NewWithGlobalsStore(code, globals)
			err = machine.Run()
			if err != nil {
				fmt.Fprintf(out, "Executing bytecode failed:\n %s\n", err)
				continue
			}

			stackTop := machine.LastPoppedStackElem()
			io.WriteString(out, stackTop.Inspect())
			io.WriteString(out, "\n")

			if options.Verbose {
				io.WriteString(out, "----DEBUG\n")
				io.WriteString(out, "Constants:\n")
				for i, constant := range constants {
					io.WriteString(out, fmt.Sprintf("%d: %s\n", i, constant.Inspect()))
				}
				io.WriteString(out, "Instructions:\n")
				io.WriteString(out, comp.Bytecode().Instructions.String())
			}
		} else {
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
}

// StartFile reads a file and executes it
func StartFile(filename string, out io.Writer, options Options) {
	env := object.NewEnvironment()

	f, err := os.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	fileContent := string(f)

	l := lexer.New(fileContent)
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		if options.Verbose {
			printLexerTokens(out, fileContent)
		}
		printParserErrors(out, p.Errors())
		return
	}

	if options.CompileEnabled {
		comp := compiler.New()
		err := comp.Compile(program)
		if err != nil {
			fmt.Fprintf(out, "Compilation failed:\n %s\n", err)
		}

		machine := vm.New(comp.Bytecode())
		err = machine.Run()
		if err != nil {
			fmt.Fprintf(out, "Executing bytecode failed:\n %s\n", err)
		}

		stackTop := machine.LastPoppedStackElem()
		io.WriteString(out, stackTop.Inspect())
		io.WriteString(out, "\n")

		if options.Verbose {
			io.WriteString(out, "----DEBUG\n")
			io.WriteString(out, comp.Bytecode().Instructions.String())
		}
	} else {
		result := interpreter.Eval(program, env)
		if result != nil {
			io.WriteString(out, result.Inspect())
			io.WriteString(out, "\n")
		} else {
			io.WriteString(out, "nil\n")
		}

		if options.Verbose {
			io.WriteString(out, "----DEBUG\n")

			printLexerTokens(out, fileContent)

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
