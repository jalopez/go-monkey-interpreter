// main.go

package main

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"

	"github.com/jalopez/go-monkey-interpreter/pkg/repl"
)

func main() {
	argparser := argparse.NewParser("monkey", "Monkey programming language interpreter")
	verbose := argparser.Flag("v", "verbose", &argparse.Options{Required: false, Help: "Show verbose output (lexer tokens and AST)"})
	interpreter := argparser.Flag("e", "eval", &argparse.Options{Required: false, Help: "Use interpreter to evaluate expression instead of compiler"})
	file := argparser.StringPositional(&argparse.Options{Required: false, Help: "File to execute"})
	// Parse input
	err := argparser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		_, err = fmt.Print(argparser.Usage(err))
		if err != nil {
			panic(err)
		}
		os.Exit(1)
	}

	if *file != "" {
		repl.StartFile(*file, os.Stdout, repl.Options{
			Verbose: *verbose,
		})
		return
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
		Verbose:        *verbose,
		UseInterpreter: *interpreter,
	})
}
