package lexer

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

/* FIRST
term = {literal, rule_name, dog, b_bracket}
list = {literal, rule_name, dog}
opt_line_end = {new_line, e}
opt_whitespace = {whitespace, e}
expression =  {literal, rule_name, dog}
rule = {e, b_bracket}
declaration = { axiom, rule_name}
declaration_list = { axiom, rule_name }
syntax = { axiom, rule_name, b_bracket }
*/

/* FOLLOW
syntax = {$}
declaration_list = {$,\n}
declaration = { whitespace, next, new_line}
rule = { $, new_line }
expression = {new_line, whitespace, b_c_bracket}
list = { whitespace, line_end, or }
term = { whitespace, line_end, or }
 */

func TestRegexp(t *testing.T) {
	def, err := Regexp(`(?P<whitespace>[ ]+)|(?P<new_line>\n+)|(?P<dog>@)|(?P<or>:)|(?P<next>,)|(?P<b_bracket>\[)|(?P<b_c_bracket>\])|(?P<literal>\"([a-zA-Z1-9]|\|| |-|!|#|%|\$|\(|\)|\*)+\")|(?P<rule_name>[a-zA-Z]+[\d\']*)|(?P<axiom>\{[ ]*[a-zA-Z]+[\d\']*[ ]*\})`)
	require.NoError(t, err)
	require.Equal(t, map[string]rune{
		"EOF":    -1,
		"whitespace": -2,
		"new_line": -3,
		"dog": -4,
		"or": -5,
		"next": -6,
		"b_bracket": -7,
		"b_c_bracket": -8,
		"literal": -9,
		"rule_name": -11,
		"axiom": -12,
	}, def.Symbols())
	lexer, err := def.Lex(strings.NewReader("{    a123'   } a \"a\" a123'"))
	require.NoError(t, err)
	tokens, err := ConsumeAll(lexer)
	require.NoError(t, err)
	require.Equal(t, []Token{
		{Type: -12, Value: "{    a123'   }", Pos: Position{Filename: "", Offset: 0, Line: 1, Column: 1}},
		{Type: -2, Value: " ", Pos: Position{Filename: "", Offset: 14, Line: 1, Column: 15}},
		{Type: -11, Value: "a", Pos: Position{Filename: "", Offset: 15, Line: 1, Column: 16}},
		{Type: -2, Value: " ", Pos: Position{Filename: "", Offset: 16, Line: 1, Column: 17}},
		{Type: -9, Value: `"a"`, Pos: Position{Filename: "", Offset: 17, Line: 1, Column: 18}},
		{Type: -2, Value: " ", Pos: Position{Filename: "", Offset: 20, Line: 1, Column: 21}},
		{Type: -11, Value: "a123'", Pos: Position{Filename: "", Offset: 21, Line: 1, Column: 22}},
		{Type: EOF, Value: "", Pos: Position{Filename: "", Offset: 26, Line: 1, Column: 27}},
	}, tokens)
	lexer, err = def.Lex(strings.NewReader("hello ?"))
	require.NoError(t, err)
	_, err = ConsumeAll(lexer)
	require.Error(t, err)
}


