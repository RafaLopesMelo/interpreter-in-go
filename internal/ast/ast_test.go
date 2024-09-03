package ast

import (
	"testing"

	"github.com/RafaLopesMelo/rmlang/internal/token"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{
					Type:    token.LET,
					Literal: "let",
				},
				Name: &Identifier{
					Token: token.Token{
						Type:    token.IDENT,
						Literal: "myVar",
					},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{
						Type:    token.IDENT,
						Literal: "myValue",
					},
					Value: "myValue",
				},
			},
		},
	}

	str := program.String()

	if str != "let myVar = myValue;" {
		t.Errorf("program.String() wrong. got %q", str)
	}
}
