package ast

import (
	"bytes"
	"strings"

	"github.com/RafaLopesMelo/monkey-lang/internal/token"
)

type HashLiteral struct {
	Token token.Token // The '{' token
	Pairs map[Expression]Expression
}

func (hl *HashLiteral) expressionNode() {}

func (hl *HashLiteral) TokenLiteral() string {
	return hl.Token.Literal
}

func (hl *HashLiteral) String() string {
	var out bytes.Buffer

	pairs := []string{}
	for key, value := range hl.Pairs {
		pairs = append(pairs, key.String()+": "+value.String())
	}

	out.WriteString(hl.TokenLiteral())
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}
