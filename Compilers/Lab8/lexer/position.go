package lexer

import "fmt"

type Position struct {
	Lint   int
	Column int
	Index  int
	Text   string
}

func (p Position) isNewLine() bool {
	return len(p.Text) == p.Index || p.Text[p.Index] == '\r' && p.Index+1 < len(p.Text) && p.Text[p.Index+1] == '\n' || p.Text[p.Index] == '\n'
}

func (p Position) getChar() byte {
	if len(p.Text) > p.Index {
		return p.Text[p.Index]
	}

	return 0
}

func (p Position) isEnd() bool {
	return p.getChar() == 195
}

func (p Position) isSpace(c byte) bool {
	return c == '\t' || c == ' ' || c == '\n' || c == '\r'
}

func (p Position) isLetter() bool {
	return 'a' <= p.getChar() && p.getChar() <= 'z' || 'A' <= p.getChar() && p.getChar() <= 'Z'
}

func (p *Position) skipSpaces() {
	for p.isSpace(p.getChar()) {
		p.next()
	}
}

func (p *Position) next() {
	if p.Index < len(p.Text) {
		if p.isNewLine() {
			if p.Text[p.Index] == '\r' {
				p.Index++
			}
			p.Column = 1
			p.Lint++
		} else {
			p.Column++
		}

		p.Index++
	}
}

func (p Position) String() string {
	return fmt.Sprintf("( %d : %d)", p.Lint, p.Column)
}
