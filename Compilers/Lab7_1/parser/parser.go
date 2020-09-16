package main

import (
	"errors"
	"fmt"
	"github.com/AntonPriyma/compilers/7/1/lexer"
	"io/ioutil"
	"os"
)

//type RuleTag int
//
//type Rules []RuleTag
//
//type Table []Rules
//
//var table = Table{
//
//}

type Stack []int

func NewStack(n int) Stack {
	return Stack(make([]int, 0, n))
}

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str int) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

func (s Stack) Cap() int {
	return s[len(s)-1]
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func (s Stack) isTerm() bool {
	return s.Cap() < 0
}

func isTerm(next int) bool {
	return next < 0
}

type NTag int

type ruleR []int

func (r ruleR) isEmpty() bool {
	return len(r) == 0
}

const (
	Syntax NTag = iota
	DeclarationList
	Declaration
	Rule
	Expression
	ExpressionInner
	List
	ListInner
	Term
	OptLineEnd
	END
)

var NTermToString = map[NTag]string {
	Syntax: "Syntax",
	DeclarationList: "Declaration list",
	Declaration: "Decalration",
	Rule: "Rule",
	Expression: "Expression",
	ExpressionInner: "Expression Inner",
	List: "List",
	ListInner: "ListInner",
	Term: "Term",
	OptLineEnd: "OptLineEnd",
	END: "End",
}

var DomainToString  = map[int]string{
	-1: "EOF",
	-2: "whitespace",
	-3: "new line",
	-4: "dog",
	-5: "or",
	-6: "next",
	-7: "b_bracket",
	-8: "b_c_bracker",
	-9: "literal",
	-11: "rule name",
	-12: "rule name",
}

type table [][]ruleR

func delete(a []int, i int) {
	copy(a[i:], a[i+1:]) // Shift a[i+1:] left one index.
	a[len(a)-1] = 0      // Erase last element (write zero value).
	a = a[:len(a)-1]     // Truncate slice. slice.
}

func readProgram() ([]byte, error) {
	if len(os.Args) < 2 {
		return nil, errors.New("wrong args: Usage: parser <path sourse code>")
	}

	return ioutil.ReadFile(os.Args[1])
}

/*
<syntax>         ::= <rule> <syntax> | <declaration_list> <syntax>  | ""
 <declaration_list> ::= <declaration> <declaration_list> | <next> <declaration_lint> | <opt_whitespace> <declaration_list> | ""
 <declaration>     ::= <axiom> | <rule_name>
 <rule>           ::= <b_bracket> <opt_line_end> <opt_whitespace> <rule_name> <opt_whitespace> <or> <opt_whitespace> <expression> <opt_line_end> <b_c_bracket> <line_end>
 <opt_whitespace> ::= <whitespace>*
 <expression>     ::= <list><expression_1>
 <expression_1> ::= <opt_whitespace> <opt_line_end> <or> <opt_whitespace> <expression> | ""
 <opt_line_end>   ::= <line_end>*
 <line_end>       ::= <new_line> | <line_end> <line_end>
 <list>           ::= <term><list_1><line_end> | ""
 <list_1> ::=   <opt_whitespace> <list>
 <term>           ::= <literal> | <rule_name> | <dog>

 <axiom>      ::= <a_bracket> <opt_whitespace> <rule_name> <opt_whitespace> <a_c_bracket>
 <literal>        ::= <quote> <text> <quote>
 <text>          ::= <character> <text> | <character>
 <character>      ::= <letter> | <digit> | <symbol>
 <rule_name>      ::= <letter> | <rule_name> <rule_char>
 <next> ::= ","
 <a_bracket>  ::= "{"
 <a_c_bracket> ::= "}"
 <b_bracket> ::= "["
 <b_c_bracket> ::= "]"
 <or> ::= ":"
 <new_line> ::= "\n"
 <whitespace> ::= " "
 <dog> ::= "@"
 <quote> ::= "`"
 <rule_char>      ::= <letter> | <digit> | "'"
 <symbol>         ::=  "|" | " " | "-" | "!" | "#" | "$" | "%" | "&" | "(" | ")" | "*" | "+" | "," | "-" | "." | "/" | ":" | ";" | "<" | "="
 <letter>         ::= "A" | "B" | "C" | "D" | "E" | "F" | "G" | "H" | "I" | "J" | "K" | "L" | "M" | "N" | "O" | "P" | "Q" | "R" | "S" | "T" | "U" | "V" | "W" | "X" | "Y" | "Z" | "a" | "b" | "c" | "d" | "e" | "f" | "g" | "h" | "i" | "j" | "k" | "l" | "m" | "n" | "o" | "p" | "q" | "r" | "s" | "t" | "u" | "v" | "w" | "x" | "y" | "z"
 <digit>          ::= "0" | "1" | "2" | "3" | "4" | "5" | "6" | "7" | "8" | "9"
*/

type ParseNode struct {
	Val    lexer.Token
	Childs []*ParseNode
}

type ParseTree struct {
	ruleStack [][]int
	nodeStack []ParseNode
	Root      *ParseNode
}

func NewTree() ParseTree {
	return ParseTree{
		ruleStack: make([][]int, 0),
		nodeStack: make([]ParseNode, 0),
		Root: &ParseNode{
			Val:    lexer.Token{},
			Childs: make([]*ParseNode, 0),
		},
	}
}

type parseError struct {
	tree ParseTree
	error
}

type parser struct {
	tokens []lexer.Token
	rules  table
}

func (p *parser) parse() error {
	tree := NewTree()
	stack := NewStack(0)
	stack.Push(int(END))
	stack.Push(int(Syntax))

	var res []ruleR

	var in []int

	for _, token := range p.tokens {
		in = append(in, int(token.Type))
	}

	i := 0
	j := 0
	for {
		node := tree.Root
		next := in[i]

		x := stack.Cap()
		if x == int(END) {
			break
		}

		if next == 0 {
			next = -1
		}

		if isTerm(x) {
			if x == next {
				stack.Pop()
				delete(in, i)
				j++
			} else {
				if x == -1 {
					stack.Pop()
					continue
				}
				return parseError{
					tree:  tree,
					error: errors.New(fmt.Sprintf("Expected term %q, got %q", DomainToString[x], DomainToString[next])),
				}

			}
		} else if !p.rules[x][-next-1].isEmpty() {
			stack.Pop()
			for _, r := range p.rules[x][-next-1] {
				stack.Push(r)
			}
			res = append(res, p.rules[x][-next-1])
			newNode := &ParseNode{
				Val:    p.tokens[j],
				Childs: make([]*ParseNode, 0)}
			node.Childs = append(node.Childs, newNode)
			tree.nodeStack = append(tree.nodeStack, *newNode)
			tree.ruleStack = append(tree.ruleStack, p.rules[x][-next-1])
		} else {
			if x == int(OptLineEnd) {
				stack.Pop()
				continue
			}
			return parseError{
				tree:  tree,
				error: errors.New(fmt.Sprintf("No rule for pair: (Noterm: %v, Term %v)", NTermToString[NTag(x)], DomainToString[next])),
			}
		}

	}

	fmt.Printf("Success parse.\n Rule stack: %#v.\n Node Stack %#v.\n Parse tree: %#v", tree.ruleStack, tree.nodeStack, tree.Root)
	return nil
}
