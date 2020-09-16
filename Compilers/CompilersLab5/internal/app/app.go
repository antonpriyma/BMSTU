package app

import (
	"errors"
	"fmt"
	scanner2 "github.com/AntonPriyma/BMSTU/CompilersLab4/internal/pkg/scanner"
	"io/ioutil"
	"log"
	"os"
	"unicode"

	"github.com/AntonPriyma/BMSTU/CompilersLab4/internal/pkg/models"
)

type StateTable [][]State

type SymbolType int

const (
	EOF             SymbolType = -1
	SpaceSymbolType SymbolType = iota - 1
	NewLineSymbolType
	Operation1SymbolType
	Operation21SymbolType
	Operation22SymbolType
	DigitSymbolType
	KeywMSymbType
	KeywJSymbType
	KeywOSymbType
	KeywPSymbType
	KeywVSymbType
	StringBegSymbType
	StringEndSymbType
	LetterSymbType
	OtherSymbolType
)

type State int

const (
	FinishState  State = -1
	InitialState State = iota - 1
	SpaceState
	NumberState
	OP1State
	OP2State
	STRState
	Key1State
	IdentState
	Key2State
	OP11State
	MState
	OState
	JState
	M1State
	EndSTRState
)

var keyToString = map[State]string{
	FinishState:  "empty",
	InitialState: "init",
	SpaceState:   "space",
	NumberState:  "number",
	OP1State:     "operator ->",
	OP2State:     "operator $",
	EndSTRState:  "string",
	Key1State:    "key_first",
	IdentState:   "ident",
	Key2State:    "key_second",
}

// 0 - " "|"\t"
// 1 - "\n"
// 2 - "$"
// 3 - '-'
// 4 - '>'
// 5 - [0-9]{1}
// 6 - m
// 7 - j
// 8 - o
// 9 - p
// 10 - v
// 11 - #
// 12 - h
// 13 - [a-zA-Z]{1}
// 14 - other letters

var table = StateTable{
	{1, 1, 4, 9, -1, 2, 10, 12, 7, 7, 7, 5, 7, 7, -1},                // 0 - initial
	{1, 1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},       // 1 - space
	{-1, -1, -1, -1, -1, 2, -1, -1, -1, -1, -1, -1, -1, -1, -1},      // 2 - number
	{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},     // 3 - ->op
	{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},     // 4 - $op
	{5, -1, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 14, 5, 5},                  // 5 - STR
	{-1, -1, -1, -1, -1, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, -1},           // 6 - KEY_MOV
	{-1, -1, -1, -1, -1, 7, 7, 7, 7, 7, 7, -1, 7, 7, -1},             // 7 - IDENT
	{-1, -1, -1, -1, -1, 7, 7, 7, 7, 7, 7, 7, -1, 7, 7, -1},          // 8 - KEY_JMP
	{-1, -1, -1, -1, 3, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},      // 9 - "-"
	{-1, -1, -1, -1, -1, 7, 7, 7, 11, 7, -1, 11, 13, -1, -1},         // 10 - m
	{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 6, -1, 7, 7, -1},        // 11 - o
	{-1, -1, -1, -1, -1, -1, 13, 7, 7, 7, 7, -1, 7, 7, -1, -1},       // 12 - j
	{-1, -1, -1, -1, -1, 7, 7, 7, 7, 8, -1, 7, 7, -1, -1},            // 13 - m
	{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}, // 14 - end string
}

func getCode(char byte) SymbolType {
	switch char {
	case models.EOF:
		return EOF
	case ' ', '\t':
		return SpaceSymbolType
	case '\n':
		return NewLineSymbolType
	case '$':
		return Operation1SymbolType
	case '-':
		return Operation21SymbolType
	case '>':
		return Operation22SymbolType
	case 'm':
		return KeywMSymbType
	case 'j':
		return KeywJSymbType
	case 'o':
		return KeywOSymbType
	case 'p':
		return KeywPSymbType
	case 'v':
		return KeywVSymbType
	case '#':
		return StringBegSymbType
	case 'h':
		return StringEndSymbType
	default:
		if unicode.IsDigit(rune(char)) {
			return DigitSymbolType
		}
		if unicode.IsLetter(rune(char)) {
			return LetterSymbType
		}
	}

	return OtherSymbolType
}

func Lexer() {
	program, err := readProgram()
	if err != nil {
		log.Fatal(err)
	}

	compiler := models.Compiler{}
	scanner := scanner2.NewStateScanner(string(program), &compiler)

	for {
		token := scanner.NextToken()

		if token.DomainTag == scanner2.EOFDomainTag {
			break
		}

		fmt.Println(token)
	}

	for _, message := range compiler.MessageList {
		fmt.Println(message)
	}
}

func readProgram() ([]byte, error) {
	if len(os.Args) < 2 {
		return nil, errors.New("wrong args: Usage: lexer <path sourse code>")
	}

	return ioutil.ReadFile(os.Args[1])
}


//syntax         -> declaration_list | rule | rule syntax | declaration_list syntax.
//declaration_list -> declaration opt_whitespace "," opt_whitespace declaration_list line_end* | declaration line_end.
//declaration 	  -> axiom | rule_name.
//axiom		  -> "{" opt_whitespace rule_name opt_whitespace "}".
//rule           -> opt_whitespace "[" opt_line_end opt_whitespace rule_name opt_whitespace ":" opt_whitespace expression opt_whitespace opt_line_end "]" line_end.
//opt_whitespace -> " "|"".
//expression     -> list | list opt_whitespace opt_line_end ":" opt_whitespace expression
//opt_line_end   -> line_end*
//line_end       -> opt_whitespace "\n" | line_end line_end
//list           -> term | term opt_whitespace list
//term           -> literal | rule_name | "@"
//literal        -> "`" text "`"
//text          -> character text | character
//character      -> letter | digit | symbol
//symbol         ->  "|" | " " | "-" | "!" | "#" | "$" | "%" | "&" | "(" | ")" | "*" | "+" | "," | "-" | "." | "/" | ":" | ";" | "<" | "="
//letter         -> "A" | "B" | "C" | "D" | "E" | "F" | "G" | "H" | "I" | "J" | "K" | "L" | "M" | "N" | "O" | "P" | "Q" | "R" | "S" | "T" | "U" | "V" | "W" | "X" | "Y" | "Z" | "a" | "b" | "c" | "d" | "e" | "f" | "g" | "h" | "i" | "j" | "k" | "l" | "m" | "n" | "o" | "p" | "q" | "r" | "s" | "t" | "u" | "v" | "w" | "x" | "y" | "z"
//digit          -> "0" | "1" | "2" | "3" | "4" | "5" | "6" | "7" | "8" | "9"
//rule_name      -> letter | rule_name rule_char
//rule_char      -> letter | digit | "'" | "_"