package ast

import (
	"simian/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "another"},
					Value: "another",
				},
			},
		},
	}
	if program.String() != "let myVar = another;" {
		t.Errorf("program.String() incorrect. Got=%q", program.String())
	}
}
