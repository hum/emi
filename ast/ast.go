package ast

import (
	"github.com/hum/emi/token"
)

//	Every node in the AST has to implement the node interface.
type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	StatementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Identifier struct {
	Token token.Token // the IDENT token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

type LetStatement struct {
	Token token.Token // the LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) tokenLiteral() string { return ls.Token.Literal }

//	The parent Node of the AST
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}
