package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParsingErrors(t, p)
	if program == nil {
		t.Fatalf("Program is nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("Length of program is not 3, actual length: %d", len(program.Statements))
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
		if !testLetStatements(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatements(t *testing.T, statement ast.Statement, name string) bool {
	if statement.TokenLiteral() != "let" {
		t.Errorf("Statement is not a Let Statement")
		return false
	}
	letstmt, ok := statement.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", statement)
		return false
	}
	if letstmt.Name.IdentValue != name {
		t.Errorf("Name.value is not %s, got this instead: %s", name, letstmt.Name.IdentValue)
		return false
	}
	if letstmt.Name.TokenLiteral() != name {
		t.Errorf("Name.TokenLiteral is not %s, got this instead: %s", name, letstmt.Name.TokenLiteral())
		return false
	}
	return true
}

func checkParsingErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}
