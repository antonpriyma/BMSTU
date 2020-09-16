package parser

import (
	"github.com/AntonPriyma/compilers/8/lexer"
)

type Symbol struct {
	IsTerm bool
	Name   string
}

type Grammar int

const (
	TerminalSymbol Grammar = iota
	NonTerminalSymbol
	ExpressionSymbol
)

type Fact struct {
	Type       Grammar
	Symbol     Symbol
	Expression Expression
	Reg        lexer.Tag
}

type Term struct {
	Factories []Fact
}

type Expression struct {
	Terms []Term
}

type Product struct {
	NTerm      Symbol
	Expression Expression
}
