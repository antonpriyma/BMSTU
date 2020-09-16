package lexer

import (
	"bytes"
	"io"
	"io/ioutil"
	"regexp"
	"strings"
	"unicode/utf8"
)

var eolBytes = []byte("\n")

type DomainTag int

const (
	EOFDomainTag = EOF
	Whitespace = -2
	NewLine = -3
	Dog = -4
)

type regexpDefinition struct {
	re      *regexp.Regexp
	symbols map[string]rune
}



func Regexp(pattern string) (Definition, error) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	symbols := map[string]rune{
		"EOF": EOF,
	}
	for i, sym := range re.SubexpNames()[1:] {
		if sym != "" {
			symbols[sym] = EOF - 1 - rune(i)
		}
	}
	return &regexpDefinition{re: re, symbols: symbols}, nil
}

func (d *regexpDefinition) Lex(r io.Reader) (Lexer, error) {
	b, err := ioutil.ReadAll(r)
	b = []byte(strings.TrimRight(string(b), " "))
	b = []byte(strings.TrimRight(string(b), "\n"))
	b = []byte(strings.TrimLeft(string(b), " "))
	b = []byte(strings.TrimRight(string(b), "\n"))
	if err != nil {
		return nil, err
	}
	return &regexpLexer{
		pos: Position{
			Filename: NameOfReader(r),
			Line:     1,
			Column:   1,
		},
		b:     b,
		re:    d.re,
		names: d.re.SubexpNames(),
	}, nil
}

func (d *regexpDefinition) Symbols() map[string]rune {
	return d.symbols
}

type regexpLexer struct {
	pos   Position
	b     []byte
	re    *regexp.Regexp
	names []string
}

func (r *regexpLexer) Next() (Token, error) {
nextToken:
	for len(r.b) != 0 {
		matches := r.re.FindSubmatchIndex(r.b)
		if matches == nil || matches[0] != 0 {
			rn, _ := utf8.DecodeRune(r.b)
			return Token{}, Errorf(r.pos, "invalid token %q", rn)
		}
		match := r.b[:matches[1]]
		token := Token{
			Pos:   r.pos,
			Value: string(match),
		}

		
		r.pos.Offset += matches[1]
		lines := bytes.Count(match, eolBytes)
		r.pos.Line += lines
		
		if lines == 0 {
			r.pos.Column += utf8.RuneCount(match)
		} else {
			r.pos.Column = utf8.RuneCount(match[bytes.LastIndex(match, eolBytes):])
		}
		
		r.b = r.b[matches[1]:]

		
		for i := 2; i < len(matches); i += 2 {
			if matches[i] != -1 {
				if r.names[i/2] == "" {
					continue nextToken
				}
				token.Type = EOF - rune(i/2)
				break
			}
		}

		return token, nil
	}

	return EOFToken(r.pos), nil
}
