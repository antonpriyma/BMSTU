package main

import (
	"bytes"
	"fmt"
	"github.com/AntonPriyma/compilers/7/1/lexer"
	"log"
)

func main() {
	program, err := readProgram()
	if err != nil {
		log.Fatal(err)
	}

	def, err := lexer.Regexp(`(?P<whitespace>[ ]+)|(?P<new_line>\n+)|(?P<dog>@)|(?P<or>:)|(?P<next>,)|(?P<b_bracket>\[)|(?P<b_c_bracket>\])|(?P<literal>\"([a-zA-Z1-9]|\|| |-|!|#|%|\$|\(|\)|\*|\+)+\")|(?P<rule_name>[a-zA-Z]+[\d\']*)|(?P<axiom>\{[ ]*[a-zA-Z]+[\d\']*[ ]*\})`)
	if err != nil {
		log.Fatal(err)
	}
	lex, err := def.Lex(bytes.NewReader(program))
	if err != nil {
		log.Fatal(err)
	}
	tokens, err := lexer.ConsumeAll(lex)
	if err != nil {
		log.Fatal(err)
	}

	rules := table{
		[]ruleR{{-1}, {}, {-3}, {}, {}, {}, {int(Syntax), int(Rule)}, {}, {}, {}, {int(Syntax), int(DeclarationList)}, {int(Syntax), int(DeclarationList)}},
		[]ruleR{{-1}, {int(DeclarationList), -2}, {-3}, {}, {}, {int(DeclarationList), -2, -6}, {}, {}, {}, {}, {int(DeclarationList), int(Declaration)}, {int(DeclarationList), int(Declaration)}},
		[]ruleR{{}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {-11}, {-12}},
		[]ruleR{{-1}, {-1}, {-3}, {-1}, {-1}, {-1}, {-8, int(OptLineEnd), int(Expression), -2, -5, -2, -11, -2, int(OptLineEnd), -7}},
		[]ruleR{{-1}, {}, {}, {int(ExpressionInner), int(List)}, {}, {}, {}, {}, {int(ExpressionInner), int(List)}, {}, {int(ExpressionInner), int(List)}, {}},
		[]ruleR{{-1}, {int(Expression), -2, -5, int(OptLineEnd)}, {int(Expression), -2, -5, int(OptLineEnd)}, {-1}, {int(Expression), -2, -5, int(OptLineEnd)}, {-1}, {-1}, {-1}, {-1}, {-1}, {-1}, {-1}},
		[]ruleR{{-1}, {-1}, {-1}, {int(OptLineEnd), int(ListInner), int(Term)}, {-1}, {-1}, {-1}, {-1}, {int(OptLineEnd), int(ListInner), int(Term)}, {-1}, {int(OptLineEnd), int(ListInner), int(Term)}, {-1}},
		[]ruleR{{-1}, {int(List), -2}, {-1}, {}, {}, {}, {}, {}, {}, {}, {}, {}},
		[]ruleR{{}, {}, {}, {-4}, {}, {}, {}, {}, {-9}, {}, {-11}, {}},
		[]ruleR{{}, {}, {-3}, {}, {}, {}, {}, {}, {}, {}, {}, {}},
	}

	parser := parser{tokens:tokens, rules:rules}

	err = parser.parse()
	if err != nil {
		fmt.Println(err)
	}
}