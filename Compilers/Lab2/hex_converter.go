package main

import (
	"fmt"
	"go/ast"
	"go/token"
	"math"
	"strconv"
)

func convertDecimalToHex(file *ast.File) {
	ast.Inspect(file, func(node ast.Node) bool {
		if decimal, ok := node.(*ast.BasicLit); ok {

			switch decimal.Kind {
			case token.INT:
				convertInt(decimal)
			case token.FLOAT:
				convertFloat(decimal)
			}

		}
		return true
	})
}

func convertFloat(node *ast.BasicLit) {
	decimal, _ := strconv.ParseFloat(node.Value, 64)

	node.Value = fmt.Sprintf("%x", math.Float64bits(decimal))
	node.Kind = token.STRING
}

func convertInt(node *ast.BasicLit) {
	decimal, _ := strconv.Atoi(node.Value)

	node.Value = fmt.Sprintf("0x%X", decimal)
	node.Kind = token.STRING
}
