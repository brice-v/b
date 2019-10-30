package ast

import (
	"b/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&VarStatement{
				Token: token.Token{Type: token.VAR, Literal: "var"},
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

	if program.String() != "var myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}

func TestVal(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&VarStatement{
				Token: token.Token{Type: token.VAL, Literal: "val"},
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

	if program.String() != "val myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}

func TestReturnStatement(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&ReturnStatement{
				Token: token.Token{Type: token.RETURN, Literal: "return"},
				ReturnValue: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
			},
		},
	}

	if program.String() != "return myVar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
