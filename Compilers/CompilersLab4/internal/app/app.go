package app

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/AntonPriyma/BMSTU/CompilersLab4/internal/pkg/models"
	"github.com/AntonPriyma/BMSTU/CompilersLab4/internal/pkg/scanner"
)

func Lexer() {
	program, err := readProgram()
	if err != nil {
		log.Fatal(err)
	}

	//add eof
	program = append(program, models.EOF)

	compiler := &models.Compiler{}
	scanner := scanner.NewScanner(string(program), compiler)

	fmt.Printf("PROGRAMM:\n%s \n", program)
	fmt.Printf("TOKENS: \n")
	for {
		token := scanner.NextToken()

		if token.DomainTag == models.EndOfFile {
			break
		}

		if token.DomainTag != models.Unknown {
			fmt.Println(token)
		}
	}

	fmt.Printf("MESSAGES:\n")
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
