// main.go

package main

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"

	"github.com/jalopez/go-monkey-interpreter/pkg/repl"
)

func main() {
	parser := argparse.NewParser("monkey", "Monkey programming language interpreter")
	verbose := parser.Flag("v", "verbose", &argparse.Options{Required: false, Help: "Show verbose output"})
	jsonAST := parser.Flag("j", "json-ast", &argparse.Options{Required: false, Help: "Print AST as JSON output"})

	// Parse input
	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		_, err = fmt.Print(parser.Usage(err))
		if err != nil {
			panic(err)
		}
		os.Exit(1)
	}

	_, err = fmt.Printf("Hello! This is the Monkey programming language!\n")

	if err != nil {
		panic(err)
	}
	_, err = fmt.Printf("Feel free to type in commands\n")
	if err != nil {
		panic(err)
	}
	repl.Start(os.Stdin, os.Stdout, repl.Options{
		Verbose: *verbose,
		JSON:    *jsonAST,
	})
}
