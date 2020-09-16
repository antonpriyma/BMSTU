package models

import (
	"fmt"
	"unicode"
)

const EOF = 0
const EndByte byte = 0

type Position struct {
	Value string

	Line  int
	Pos   int
	Index int
}

func (p *Position) NextSymb() {
	if !p.IsEOF() {
		if p.IsNewLine() {
			if p.Value[p.Index] == '\r' {
				p.Index++
			}
			p.Line++
			p.Pos = 1
		} else {
			p.Pos++
		}
		p.Index++
	}
}

func (p Position) CurValue() byte {
	if p.IsEOF() {
		return EndByte
	}
	return p.Value[p.Index]
}

func (p Position) PrevValue() byte {
	if p.Index == 0 {
		return ' '
	}

	return p.Value[p.Index-1]
}

func (p Position) String() string {
	return fmt.Sprintf("(%d,%d)", p.Line, p.Pos)
}

func (p Position) IsWhiteSpace() bool {
	return !p.IsEOF() && p.Value[p.Index] == ' '
}

func (p Position) isLetter() bool {
	return !p.IsEOF() && unicode.IsLetter(rune(p.Value[p.Index]))
}

func (p Position) IsEOF() bool {
	return p.Index == len(p.Value)-1
}

func (p Position) IsNewLine() bool {
	if p.IsEOF() {
		return false
	}

	if p.Value[p.Index] == '\r' && p.Index+1 < len(p.Value) {
		return p.Value[p.Index+1] == '\n'
	}

	return p.Value[p.Index] == '\n'
}

func (p Position) IsDigit() bool {
	return unicode.IsDigit(rune(p.CurValue()))
}
