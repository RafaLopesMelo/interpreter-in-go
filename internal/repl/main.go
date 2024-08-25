package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/RafaLopesMelo/monkey-lang/internal/evaluator"
	"github.com/RafaLopesMelo/monkey-lang/internal/lexer"
	"github.com/RafaLopesMelo/monkey-lang/internal/object"
	"github.com/RafaLopesMelo/monkey-lang/internal/parser"
)

const PROMPT = ">> "

func StartRepl(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)

		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}

	}
}
