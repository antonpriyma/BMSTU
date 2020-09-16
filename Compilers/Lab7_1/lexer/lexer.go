package lexer

import (
	"fmt"
	"io"
)

const (
	
	EOF rune = -(iota + 1)
)


func EOFToken(pos Position) Token {
	return Token{Type: EOF, Pos: pos}
}


type Definition interface {
	
	Lex(io.Reader) (Lexer, error)
	
	
	
	Symbols() map[string]rune
}


type Lexer interface {
	
	Next() (Token, error)
}


func SymbolsByRune(def Definition) map[rune]string {
	out := map[rune]string{}
	for s, r := range def.Symbols() {
		out[r] = s
	}
	return out
}


func NameOfReader(r interface{}) string {
	if nr, ok := r.(interface{ Name() string }); ok {
		return nr.Name()
	}
	return ""
}



//

//

func Must(def Definition, err error) Definition {
	if err != nil {
		panic(err)
	}
	return def
}


func ConsumeAll(lexer Lexer) ([]Token, error) {
	tokens := []Token{}
	for {
		token, err := lexer.Next()
		if err != nil {
			return nil, err
		}
		tokens = append(tokens, token)
		if token.Type == EOF {
			return tokens, nil
		}
	}
}


type Position struct {
	Filename string
	Offset   int
	Line     int
	Column   int
}

func (p Position) GoString() string {
	return fmt.Sprintf("Position{Filename: %q, Offset: %d, Line: %d, Column: %d}",
		p.Filename, p.Offset, p.Line, p.Column)
}

func (p Position) String() string {
	filename := p.Filename
	if filename == "" {
		return fmt.Sprintf("%d:%d", p.Line, p.Column)
	}
	return fmt.Sprintf("%s:%d:%d", filename, p.Line, p.Column)
}


type Token struct {
	
	Type  rune
	Value string
	Pos   Position
}


func RuneToken(r rune) Token {
	return Token{Type: r, Value: string(r)}
}


func (t Token) EOF() bool {
	return t.Type == EOF
}

func (t Token) String() string {
	if t.EOF() {
		return "<EOF>"
	}
	return t.Value
}

func (t Token) GoString() string {
	if t.Pos == (Position{}) {
		return fmt.Sprintf("Token{%d, %q}", t.Type, t.Value)
	}
	return fmt.Sprintf("Token@%s{%d, %q}", t.Pos.String(), t.Type, t.Value)
}


//

func MakeSymbolTable(def Definition, types ...string) (map[rune]bool, error) {
	symbols := def.Symbols()
	table := map[rune]bool{}
	for _, symbol := range types {
		rn, ok := symbols[symbol]
		if !ok {
			return nil, fmt.Errorf("lexer does not support symbol %q", symbol)
		}
		table[rn] = true
	}
	return table, nil
}
