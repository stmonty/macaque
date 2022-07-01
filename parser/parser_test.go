package parser

mport (
	"simian/ast"
	"simian/lexer"
	"testing"
)

func TestLetStatement(t *testing.T) {
	input := `let x = 5; let y = 10; let variable = 42;`

	l := lexer.New(input)
	p := New(l)

	program := p.parseProgram()
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
