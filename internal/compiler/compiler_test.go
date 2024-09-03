package compiler

import "github.com/RafaLopesMelo/rmlang/internal/code"

type compilerTestCase struct {
	input                string
	expectedConstants    any
	expectedInstructions []code.Instructions
}
