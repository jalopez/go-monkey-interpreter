// main.go

package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/jalopez/go-monkey-interpreter/pkg/repl"
)

func main() {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	_, err = fmt.Printf("Hello %s! This is the Monkey programming language!\n",
		u.Username)

	if err != nil {
		panic(err)
	}
	_, err = fmt.Printf("Feel free to type in commands\n")
	if err != nil {
		panic(err)
	}
	repl.Start(os.Stdin, os.Stdout)
}
