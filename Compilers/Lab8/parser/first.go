package parser

import (
	"fmt"
	"github.com/AntonPriyma/compilers/8/lexer"
)

type First struct {
	FirstS   map[string][]string
	Products map[string]Product
}

func (f *First) First(products *map[string]Product) {
	f.Products = *products
	for v, p := range *products {
		s := make([]string, 0)
		f.findFirstS(&p.NTerm, &s)
		fmt.Println("")
		f.FirstS[v] = s
	}
}

func (f *First) findFirst(expression *Expression, firsts *[]string){
	for _, t := range expression.Terms {
		f.findFirstT(&t, firsts)
	}
}

func (f *First) findFirstT(term *Term, firsts *[]string){
	for _, fact:= range term.Factories {
		f.findFirstF(&fact, firsts)
		if !f.hasEpsilonF(fact) {
			break
		}
		*firsts = append(*firsts, "EPS")
	}
}

func (f *First) findFirstF(fact *Fact, firsts *[]string){
	if fact.Type == ExpressionSymbol {
		f.findFirst(&fact.Expression, firsts);
	} else {
		f.findFirstS(&fact.Symbol, firsts);
	}
}

func (f *First) findFirstS(s *Symbol, firsts *[]string){
	if s.IsTerm {
		*firsts = append(*firsts, s.Name)
		return
	}
	expr := f.Products[s.Name].Expression
	f.findFirst(&expr, firsts);
}

func (f *First) hasEpsilon(s Symbol) bool{
	if s.Name == "EPS" {
		return true
	}
	if s.IsTerm {
		return false
	}

	for _, t :=range f.Products[s.Name].Expression.Terms{
		if len(t.Factories) == 2 {
			fmt.Println(" ")
		}
		if f.hasEpsilonT(t) {
			return true
		}
	}

	return false
}



func (f *First) hasEpsilonT(t Term) bool {
	return f.hasEpsilonF(t.Factories[0])
}

func (f *First) hasEpsilonF(fact Fact) bool {
	if fact.Type == ExpressionSymbol {
		if fact.Reg == lexer.QUESTION || fact.Reg == lexer.STAR {
			return true
		}
		return f.hasEpsilonE(fact.Expression)
	}
	return f.hasEpsilon(fact.Symbol)
}

func (f *First) hasEpsilonE(expression Expression) bool {
	for _, t := range expression.Terms {
		if f.hasEpsilonT(t) {
			return true
		}
	}
	return false
}

func (f *First) PrintTable() {
	for v, a := range f.FirstS {
		fmt.Print(v + ":")

		for _, b := range a{
			fmt.Print(b)
		}

		fmt.Println("")
	}
}



