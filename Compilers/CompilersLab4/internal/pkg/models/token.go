package models

import "fmt"

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

	DomainTag TokenType
	Value     string
	Attr      interface{}
}

func (t Token) String() string {
	if t.Attr != nil {
		return fmt.Sprintf("({%v} %s %s) %v", t.Coords, TokenNames[t.DomainTag], t.Value, t.Attr)
	} else {
		return fmt.Sprintf("({%v} %s %s)", t.Coords, TokenNames[t.DomainTag], t.Value)
	}
}
