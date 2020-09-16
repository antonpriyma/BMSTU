package scanner

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/AntonPriyma/BMSTU/CompilersLab4/internal/pkg/models"
)

type Scanner interface {
	NextToken() models.Token
}

type scannerImpl struct {
	compiler *models.Compiler
	cur      models.Position
	program  string
}

func NewScanner(program string, compiler *models.Compiler) Scanner {
	return &scannerImpl{program: program, compiler: compiler, cur: models.Position{
		Value: program,
		Line:  1,
		Pos:   1,
		Index: 0,
	}}
}

func (s *scannerImpl) NextToken() models.Token {
	for !s.cur.IsEOF() {
		return s.ProcessNextToken()
	}
	return models.Token{
		Coords:    models.Fragment{},
		DomainTag: models.EndOfFile,
		Value:     "",
	}
}

func (s *scannerImpl) ProcessNextToken() models.Token {
	for s.cur.IsWhiteSpace() || s.cur.IsNewLine() {
		s.cur.NextSymb()
	}

	token := models.Token{}
	curValue := s.cur.CurValue()

	switch curValue {
	case '"':
		token = s.readString()
		//default:
		//	if unicode.IsDigit(rune(curValue)) {
		//		token = s.readDigit()
		//	}
		//	if (curValue > 'a' && curValue < 'z') || (curValue > 'A' && curValue < 'Z') || curValue == '_' {
		//		token = s.readIdent()
		//	}

	}

	if unicode.IsDigit(rune(curValue)) {
		token = s.readDigit()
	}

	if (curValue >= 'a' && curValue <= 'z') || (curValue >= 'A' && curValue <= 'Z') || curValue == '_' {
		token = s.readIdent()
	}

	if curValue == models.EOF {
		token = models.Token{
			Coords:    models.Fragment{},
			DomainTag: models.EndOfFile,
			Value:     "",
		}
	}

	if token.DomainTag == models.Unknown {
		s.compiler.AddError(s.cur, "unrecognized token")
		s.cur.NextSymb()
	}

	return token
}

func (s *scannerImpl) readString() models.Token {
	start := s.cur
	var end models.Position

	res := strings.Builder{}
	res.WriteByte(s.cur.CurValue())
	s.cur.NextSymb()
	escapes := make([]string, 0)

	for !(s.cur.CurValue() == '"' && s.cur.PrevValue() != '\\') && s.cur.CurValue() != '\n' && !s.cur.IsEOF() {
		res.WriteByte(s.cur.CurValue())
		if s.cur.PrevValue() == '\\' && s.cur.CurValue() != ' ' {
			escapes = append(escapes, fmt.Sprintf("%s%s", string(s.cur.PrevValue()), string(s.cur.CurValue())))
		}
		s.cur.NextSymb()
	}

	if s.cur.CurValue() != '"' {
		s.compiler.AddError(s.cur, "expected end of string")
		return models.Token{}
	} else {
		res.WriteByte(s.cur.CurValue())
		s.cur.NextSymb()

		end = models.Position{
			Line:  s.cur.Line,
			Pos:   s.cur.Pos - 1,
			Index: s.cur.Index - 1,
		}

		return models.Token{
			Coords: models.Fragment{
				Start:     start,
				Following: end,
			},
			DomainTag: models.Word,
			Value:     res.String(),
			Attr:      escapes,
		}
	}
}

func (s *scannerImpl) readIdent() models.Token {
	start := s.cur
	var end models.Position

	res := strings.Builder{}

	for isIdentByte(s.cur.CurValue()) {
		res.WriteByte(s.cur.CurValue())
		s.cur.NextSymb()
	}

	end = models.Position{
		Line:  s.cur.Line,
		Pos:   s.cur.Pos - 1,
		Index: s.cur.Index - 1,
	}

	return models.Token{
		Coords: models.Fragment{
			Start:     start,
			Following: end,
		},
		DomainTag: models.Ident,
		Value:     res.String(),
	}
}

func (s *scannerImpl) readDigit() models.Token {
	start := s.cur
	var end models.Position

	res := strings.Builder{}

	for {
		batch, err := s.readDigits()
		if err != nil {
			break
		}

		if !unicode.IsDigit(rune(s.cur.CurValue())) {
			res.Write(batch)
		} else {
			break
		}

		if s.cur.CurValue() != '.' {
			break
		}

		res.WriteByte(s.cur.CurValue())
		s.cur.NextSymb()
	}

	if unicode.IsDigit(rune(s.cur.CurValue())) {
		s.compiler.AddError(s.cur, "expected only 3 digits in batch")
		for unicode.IsDigit(rune(s.cur.CurValue())) {
			s.cur.NextSymb()
		}
		return models.Token{}
	}

	resWithoutDots := strings.ReplaceAll(res.String(), ".", "")
	attr, _ := strconv.Atoi(resWithoutDots)

	end = models.Position{
		Line:  s.cur.Line,
		Pos:   s.cur.Pos - 1,
		Index: s.cur.Index - 1,
	}

	return models.Token{
		Coords: models.Fragment{
			Start:     start,
			Following: end,
		},
		DomainTag: models.Number,
		Value:     res.String(),
		Attr:      attr,
	}
}

func (s *scannerImpl) readDigits() ([]byte, error) {
	var res []byte
	i := 0
	for i < 3 {
		i++

		symb := s.cur.CurValue()

		if !s.cur.IsDigit() {
			return []byte{}, errors.New("failed to read batch")
		}
		if symb == models.EndByte {
			return []byte{}, errors.New("failed to read batch")
		}

		res = append(res, s.cur.CurValue())
		s.cur.NextSymb()

	}

	return res, nil
}

func isIdentByte(b byte) bool {
	return unicode.IsDigit(rune(b)) || ((b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || b == '_')
}
