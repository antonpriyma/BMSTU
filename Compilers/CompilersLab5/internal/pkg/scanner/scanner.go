package scanner

import (
	"github.com/AntonPriyma/BMSTU/CompilersLab4/internal/pkg/models"
	"strings"
	"unicode"
)

type Scanner interface {
	NextToken() models.Token
}

type stateScanner struct {
	sourse   string
	cur      *models.Position
	messages []models.Message
	compiler *models.Compiler
}

func NewStateScanner(sourse string, compiler *models.Compiler) Scanner {

	sourse = strings.TrimRight(sourse, " ")
	sourse = strings.TrimRight(sourse, "\n")
	sourse = strings.TrimLeft(sourse, " ")
	sourse = strings.TrimLeft(sourse, "\n")
	sourse = string(append([]byte(sourse), models.EOF))

	cur := models.Position{
		Value: sourse,
		Line:  0,
		Pos:   1,
		Index: 0,
	}
	res := stateScanner{sourse: sourse, cur: &cur, compiler: compiler}
	return res
}

func (s stateScanner) NextToken() models.Token {
	if !s.cur.IsEOF() {
		token := s.processNextToken()

		if token.DomainTag == keyToString[SpaceState] {
			return s.processNextToken()
		}
		return token
	}

	return models.Token{
		Coords:    models.Fragment{},
		DomainTag: EOFDomainTag,
		Value:     "",
	}
}

func (s *stateScanner) processNextToken() models.Token {
	state := InitialState
	prevState := InitialState
	start := *s.cur

	for state != FinishState {
		symbState := getCode(s.cur.CurValue())
		prevState = state

		if symbState == EOF {
			state = FinishState
		} else {
			state = table[state][symbState]
		}

		if state != FinishState {
			s.cur.NextSymb()
		}
	}

	if prevState != SpaceState && prevState != InitialState {
		if prevState == OP11State {
			s.compiler.AddError(*s.cur, "error: expected '>' after '-'")
		} else if prevState == STRState {
			s.compiler.AddError(*s.cur, "error: expected 'h' as string and")
		} else {
			token := models.Token{
				Coords: models.Fragment{
					Start:     start,
					Following: *s.cur,
				},
				DomainTag: keyToString[prevState],
				Value:     s.cur.Value[start.Index:s.cur.Index],
			}
			return token
		}
	} else if prevState == SpaceState {
		return models.Token{
			Coords:    models.Fragment{},
			DomainTag: keyToString[prevState],
			Value:     "",
		}
	} else {
		for !s.cur.IsEOF() && getCode(s.cur.CurValue()) == OtherSymbolType {
			s.cur.NextSymb()
		}

		s.compiler.AddError(*s.cur, s.cur.Value[start.Index:s.cur.Index])
	}

	return models.Token{}
}

const EOFDomainTag = "EOF"

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
