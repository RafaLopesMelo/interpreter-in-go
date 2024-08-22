package ast

import (
	"bytes"
	"strings"

	"github.com/RafaLopesMelo/monkey-lang/internal/token"
)

type FunctionLiteral struct {
	Token      token.Token // The "fn" token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (f *FunctionLiteral) expressionNode() {}

func (f *FunctionLiteral) TokenLiteral() string {
	return f.Token.Literal
}

func (f *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}

	for _, param := range f.Parameters {
		params = append(params, param.String())
	}

	out.WriteString(f.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(")")
	out.WriteString(f.Body.String())

	return out.String()
}
