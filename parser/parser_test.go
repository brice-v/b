package parser

import (
	"b/ast"
	"b/lexer"
	"testing"
)

func TestVarStatements(t *testing.T) {
	input := `
	var x = 5;
	var y = 10;
	var foobar = 5353;
	`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testVarStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func TestValStatements(t *testing.T) {
	input := `
	val x = 5;
	val y = 10;
	val foobar = 5353;
	`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testValStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testVarStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "var" {
		t.Errorf("s.TokenLiteral not `let`. got=%q", s.TokenLiteral())
		return false
	}
	varStmt, ok := s.(*ast.VarStatement)
	if !ok {
		t.Errorf("s not *ast.VarStatement. got=%T", s)
		return false
	}
	if varStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not `%s`. got=%s", name, varStmt.Name)
	}
	return true
}

func testValStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "val" {
		t.Errorf("s.TokenLiteral not `let`. got=%q", s.TokenLiteral())
		return false
	}
	varStmt, ok := s.(*ast.ValStatement)
	if !ok {
		t.Errorf("s not *ast.ValStatement. got=%T", s)
		return false
	}
	if varStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not `%s`. got=%s", name, varStmt.Name)
	}
	return true
}