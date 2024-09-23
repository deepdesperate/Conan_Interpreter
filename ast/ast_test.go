package ast

import (
	"testing"

	"github.com/deepdesperate/Conan_Interpreter"
	"github.com/deepdesperate/Conan_Interpreter/token"
)

// let myVar = anotherVar

func TestString(t*testing.T) {
	program:=&Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != "let myVar=anotherVar;"{
		t.Errorf("prgram.String() wrong.got=%q",program.String())
	}
	
}