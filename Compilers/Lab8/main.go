package main

import (
	"errors"
	lexer2 "github.com/AntonPriyma/compilers/8/lexer"
	"github.com/AntonPriyma/compilers/8/parser"
	"io/ioutil"
	"log"
	"os"
)

func readProgram() ([]byte, error) {
	if len(os.Args) < 2 {
		return nil, errors.New("wrong args: Usage: parser <path sourse code>")
	}

	return ioutil.ReadFile(os.Args[1])
}

func main() {
	program, err := readProgram()
	if err != nil {
		log.Fatal(err)
	}

	lexer := lexer2.NewLexer(string(program))



	a := parser.NewAstTree(lexer.Tokens)

	products := a.Products

	f := parser.First{}
	f.FirstS = make(map[string][]string, 0)
	f.Products = make(map[string]parser.Product, 0)
	f.First(&products)

	f.PrintTable()
}