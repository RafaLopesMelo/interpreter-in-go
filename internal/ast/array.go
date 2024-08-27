package ast

import (
	"strings"

	"github.com/RafaLopesMelo/monkey-lang/internal/token"
)

type ArrayLiteral struct {
	Token    token.Token
	Elements []Expression
}

func (a *ArrayLiteral) expressionNode() {}

func (a *ArrayLiteral) TokenLiteral() string {
	return a.Token.Literal
}

func (a *ArrayLiteral) String() string {
	var elements []string

	for _, e := range a.Elements {
		elements = append(elements, e.String())
	}

	return "[" + strings.Join(elements, ", ") + "]"
}
