package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/RafaLopesMelo/monkey-lang/internal/repl"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	args := os.Args
	kind := args[1]

	fmt.Printf("Hello, %s! This is the Monkey programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")

	if kind == "lexer" {
		repl.StartLexerRepl(os.Stdin, os.Stdout)
	} else if kind == "parser" {
		repl.StartParserRepl(os.Stdin, os.Stdout)
	}
}
