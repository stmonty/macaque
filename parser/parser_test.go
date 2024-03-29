package parser

import (
	"simian/ast"
	"simian/lexer"
	"testing"
)

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.errors

	if len(errors) == 0 {
		return
	}

	t.Errorf("Parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("Parser error: %q", msg)
	}
	t.FailNow()
}

func TestLetStatement(t *testing.T) {
	input := `let x = 5; let y = 10; let variable = 42;`

	l := lexer.New(input)
	p := New(l)

	program := p.parseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("Parsing Program returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("Expected=3 : got=%d statements", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifer string
	}{
		{"x"},
		{"y"},
		{"variable"},
	}

	for i, tt := range tests {
		statement := program.Statements[i]
		if !testLetStatement(t, statement, tt.expectedIdentifer) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("Expected='let' : got=%q", s.TokenLiteral())
		return false
	}
	letStatement, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("Expected *ast.LetStatement: got=%T", s)
		return false
	}

	if letStatement.Name.Value != name {
		t.Errorf("Expected=%s : got=%s", name, letStatement.Name.Value)
		return false
	}

	if letStatement.Name.TokenLiteral() != name {
		t.Errorf("Expected=%s : got=%s", name, letStatement.Name)
		return false
	}

	return true
}

func TestReturnStatements(t *testing.T) {
	input := "return 5; return 10; return 42;"
	l := lexer.New(input)
	p := New(l)

	program := p.parseProgram()

	checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("Parsing Program returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("Expected=3 : got=%d statements", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.ReturnStatement. Found=%T", stmt)
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return'. Found=%q", returnStmt.TokenLiteral())
		}
	}
}

func TestIdentifierExpression(t *testing.T) {
	input := "hello;"
	l := lexer.New(input)
	p := New(l)

	program := p.parseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("Program has too few statements. Got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("Program.Statements[0] is not an ast.ExpressionStatement. Got=%T", program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("Exp not *ast.Identifier. Got=%T", stmt.Expression)
	}

	if ident.Value != "hello" {
		t.Errorf("Incorrect ident.Value. Not=%s, got=%s", "hello", ident.Value)
	}

	if ident.TokenLiteral() != "hello" {
		t.Errorf("ident.TokenLiteral() does not equal %s. Got=%s", "hello", ident.TokenLiteral())
	}
	
}

func TestIntegerLiteralExpression(t *testing.T) {
	input := "7;"

	l := lexer.New(input)
	p := New(l)
	program := p.parseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("Program has too few statements. Got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("Program.Statements[0] is not an ast.ExpressionStatement. Got=%T", program.Statements[0])
	}

	literal, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("Exp not *ast.Identifier. Got=%T", stmt.Expression)
	}

	if literal.Value != 7 {
		t.Errorf("Incorrect literal.Value. Not=%d, got=%d", 7, literal.Value)
	}

	if literal.TokenLiteral() != "7" {
		t.Errorf("literal.TokenLiteral() does not equal %s. Got=%s", "hello", literal.TokenLiteral())
	}
}
