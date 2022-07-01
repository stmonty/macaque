package ast

import "simian/token"

// All Expressions and Statements start with Nodes
type Node interface {
	TokenLiteral() string
}

// Statement: let a = 5;
type Statement interface {
	Node
	statementNode()
}

// Expression: add(5, 10);
type Expression interface {
	Node
	statementNode()
}

type Program struct {
	Statements []Statement
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}
