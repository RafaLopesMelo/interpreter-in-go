package compiler

import (
	"github.com/RafaLopesMelo/rmlang/internal/ast"
	"github.com/RafaLopesMelo/rmlang/internal/code"
	"github.com/RafaLopesMelo/rmlang/internal/object"
)

type ByteCode struct {
	Instructions code.Instructions
	Constants    []object.Object
}

type Compiler struct {
	instructions code.Instructions
	constants    []object.Object
}

func (c *Compiler) Compile(node ast.Node) error {
	return nil
}

func (c *Compiler) ByteCode() *ByteCode {
	return &ByteCode{
		Instructions: c.instructions,
		Constants:    c.constants,
	}
}

func New() *Compiler {
	return &Compiler{
		instructions: code.Instructions{},
		constants:    []object.Object{},
	}
}
