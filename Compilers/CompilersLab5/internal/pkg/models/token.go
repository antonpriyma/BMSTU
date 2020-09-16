package models

import (
	"fmt"
)

type TokenType int

const (
	Unknown   TokenType = 0
	Word      TokenType = 1
	Ident     TokenType = 2
	Number    TokenType = 3
	EndOfFile TokenType = 4
)

var TokenNames = map[TokenType]string{
	Word:      "string",
	Ident:     "ident",
	Number:    "number",
	EndOfFile: "EOF",
}

type Token struct {
	Coords Fragment

	DomainTag string
	Value     string
}

func (t Token) String() string {
	return fmt.Sprintf("({%v} %s %s)", t.Coords, t.DomainTag, t.Value)
}
