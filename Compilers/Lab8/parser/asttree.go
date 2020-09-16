package parser

import (
	"github.com/AntonPriyma/compilers/8/lexer"
)

type AstTree struct {
	Products map[string]Product
	Tokens   []lexer.Token
	Index    int
}

func NewAstTree(tokens []lexer.Token)  AstTree{
	a := AstTree{}
	a.Tokens = tokens
	a.Products = make(map[string]Product, 0)
	a.ParseS()

	return a
}

func (t *AstTree) ParseS() {
	for t.Tokens[t.Index].Tag == lexer.LEFT_SQUARE_PAREN {
		var product Product
		t.ParseP(&product)
		t.Products[product.NTerm.Name] = product
	}
}

func (t *AstTree) ParseP(product *Product) {
	t.Index++
	if t.Tokens[t.Index].Tag != lexer.NTERM {
		return
	}

	product.NTerm = Symbol{
		IsTerm: false,
		Name:   t.Tokens[t.Index].Text,
	}

	t.Index++
	if t.Tokens[t.Index].Tag != lexer.COLON {
		return
	}
	t.Index++
	var expression Expression
	t.ParseE(&expression)
	product.Expression = expression
	if t.Tokens[t.Index].Tag != lexer.RIGHT_SQUARE_PAREN {
		return
	}
	t.Index++

}

func (t *AstTree) ParseE(expression *Expression) {
	var term Term
	t.ParseT(&term)
	expression.Terms = append(expression.Terms, term)

	for t.Tokens[t.Index].Tag == lexer.COLON {
		var term Term
		t.Index++
		t.ParseT(&term)
		expression.Terms = append(expression.Terms, term)
	}
}

func (t *AstTree) ParseT(term *Term) {
	var factory Fact
	t.ParseF(&factory)
	term.Factories = append(term.Factories, factory)

	for t.Tokens[t.Index].Tag == lexer.NTERM || t.Tokens[t.Index].Tag == lexer.TERM || t.Tokens[t.Index].Tag == lexer.LEFT_SQUARE_PAREN {
		factory := Fact{}
		t.ParseF(&factory)
		term.Factories = append(term.Factories, factory)
	}
}

func (t *AstTree) ParseF(f *Fact) {
	tag := t.Tokens[t.Index].Tag

	if tag == lexer.NTERM || tag == lexer.TERM {
		symbol := Symbol{
			IsTerm: tag == lexer.TERM,
			Name:   t.Tokens[t.Index].Text,
		}

		factType := Grammar(0)
		if symbol.IsTerm {
			factType = TerminalSymbol
		} else {
			factType = NonTerminalSymbol
		}
		*f = Fact{
			Type:   factType,
			Symbol: symbol,
		}
	} else if tag == lexer.LEFT_SQUARE_PAREN {
		t.Index++
		expression := Expression{}
		t.ParseE(&expression)
		if t.Tokens[t.Index].Tag != lexer.RIGHT_SQUARE_PAREN {
			return
		}
		t.Index++
		tag := t.Tokens[t.Index].Tag
		if tag == lexer.PLUS || tag == lexer.STAR || tag == lexer.QUESTION {
			*f = Fact{
				Type:       ExpressionSymbol,
				Expression: expression,
				Reg:        tag,
			}
		} else {
			*f = Fact{
				Type:       ExpressionSymbol,
				Expression: expression,
				Reg:        lexer.EPSILON,
			}
		}
	} else {
		return
	}

	t.Index++
}
