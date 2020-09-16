package lexer

import "fmt"


type Error struct {
	Msg string
	Tok Token
}


func Errorf(pos Position, format string, args ...interface{}) *Error {
	return &Error{Msg: fmt.Sprintf(format, args...), Tok: Token{Pos: pos}}
}


func ErrorWithTokenf(tok Token, format string, args ...interface{}) *Error {
	return &Error{Msg: fmt.Sprintf(format, args...), Tok: tok}
}

func (e *Error) Message() string { return e.Msg } 
func (e *Error) Token() Token    { return e.Tok } 


func (e *Error) Error() string {
	return FormatError(e.Tok.Pos, e.Msg)
}


func FormatError(pos Position, message string) string {
	msg := ""
	if pos.Filename != "" {
		msg += pos.Filename + ":"
	}
	if pos.Line != 0 || pos.Column != 0 {
		msg += fmt.Sprintf("%d:%d:", pos.Line, pos.Column)
	}
	if msg != "" {
		msg += " " + message
	} else {
		msg = message
	}
	return msg
}
