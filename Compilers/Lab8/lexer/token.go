package lexer

import "fmt"

type Token struct {
	Tag      Tag
	Text     string
	Fragment Fragment
}

func (t Token) String() string {
	return fmt.Sprintf("%s %v %s", TagToString[t.Tag], t.Fragment, t.Text)
}
