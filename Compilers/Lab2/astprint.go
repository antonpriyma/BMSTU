package main

import (
	"fmt"
	"go/format"
	"go/parser"
	"go/token"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	fset := token.NewFileSet()
	if file,err := parser.ParseFile(fset,os.Args[1],nil,parser.ParseComments); err==nil {
		convertDecimalToHex(file)

		if format.Node(os.Stdout, fset, file) != nil {
			fmt.Printf("Formatter error: %v\n", err)
		}
		//ast.Fprint(os.Stdout, fset, file, nil)
	} else {
		fmt.Printf("Errors in %s\n", os.Args[1])
	}


}