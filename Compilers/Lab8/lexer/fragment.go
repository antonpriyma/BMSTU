package lexer

import "fmt"

type Fragment struct {
	First  Position
	Second Position
}

func (f Fragment) String() string {
	return fmt.Sprintf("%s - %s", f.First, f.Second)
}
