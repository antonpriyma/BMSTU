package lexer

import "fmt"

type Lexer struct {
	Text    string
	Current *Position
	Tokens  []Token
}

func NewLexer(s string) Lexer {
	lex := Lexer{
		Text:    s + string(0xFF),
		Current: &Position{Text: s + string(0xFF)},
	}

	lex.createTokens()

	return lex
}

func (l *Lexer) GetTokenByPos(start *Position) Token {
	for !start.isEnd() {
		start.skipSpaces()
		pos := *start
		c := pos.getChar()
		var t Token

		if c == '{' || c == '}' || c == '[' || c == ']' || c == ':' || c == '@' || c == ',' || c == '*' || c == '?' || c == '+' {
			pos.next()
			tag := EPS

			if c == '{' {
				tag = LEFT_FIGURE_PAREN
			} else if c == '}' {
				tag = RIGHT_FIGURE_PAREN
			} else if c == '[' {
				tag = LEFT_SQUARE_PAREN
			} else if c == ']' {
				tag = RIGHT_SQUARE_PAREN
			} else if c == ':' {
				tag = COLON
			} else if c == ',' {
				tag = COMMA
			} else if c == '+' {
				tag = PLUS
			} else if c == '*' {
				tag = STAR
			} else if c == '?' {
				tag = QUESTION
			}

			t = Token{
				Tag:  tag,
				Text: string(c),
				Fragment: Fragment{
					First:  *start,
					Second: pos,
				},
			}
		} else if pos.isLetter() {
			t = l.skipIdent(&pos)
		} else if pos.getChar() == '"' {
			t = l.skipString(&pos)
		} else if c == byte(0xFF) {
			break
		} else {
			t = Token{
				Tag:  ERROR,
				Text: "",
				Fragment: Fragment{
					First:  *start,
					Second: *start,
				},
			}
		}

		if t.Tag != ERROR {
			*start = pos
			return t
		}

		fmt.Println("Error in " + start.String())
		start.next()
	}
	return Token{
		Tag:  END_OF_PROGRAM,
		Text: "EOF",
		Fragment: Fragment{
			First:  *start,
			Second: *start,
		},
	}

}

func (l *Lexer) skipString(cur *Position) Token {
	if cur.getChar() != '"' {
		return Token{}
	}
	start := *cur
	cur.next()
	word := ""

	for cur.getChar() != '"' && cur.getChar() != byte(0xFF) {
		if cur.getChar() == '\\' {
			cur.next()
		}
		word += string(cur.getChar())
		cur.next()
	}

	if cur.getChar() == byte(0xFF) {
		return Token{
			Tag:  ERROR,
			Text: word,
			Fragment: Fragment{
				First:  start,
				Second: *cur,
			},
		}
	}

	word = `"` + word + `"`
	cur.next()
	return Token{
		Tag:  TERM,
		Text: word,
		Fragment: Fragment{
			First:  start,
			Second: *cur,
		},
	}
}

func (l *Lexer) skipIdent(cur *Position) Token {
	start := *cur
	word := ""

	for cur.isLetter() {
		word += string(cur.getChar())
		cur.next()
	}

	if cur.getChar() == '\'' {
		word = string(cur.getChar())
		cur.next()
	}

	return Token{
		Tag:  NTERM,
		Text: word,
		Fragment: Fragment{
			First:  start,
			Second: *cur,
		},
	}
}

func (l *Lexer) createTokens() {
	token := Token{}

	for token.Tag != END_OF_PROGRAM {
		token = l.GetTokenByPos(l.Current)
		l.Tokens = append(l.Tokens, token)
		if token.Fragment.First.Lint == 5 {
			fmt.Println("")
		}
		fmt.Println(token)
	}
}
