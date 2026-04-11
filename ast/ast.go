package ast

import "monkey/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.statements) > 0 {
		return p.statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

type Identifier struct {
	Token token.Token
	Value string
}

func (l *LetStatement) statementNode()

func (l *LetStatement) TokenLiteral() string {
	return l.Token.Literal
}

func (i *Identifier) expressionNode()

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
