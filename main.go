package main

import (
	"fmt"

	"github.com/wyst-lang/wyst/parser"
)

func main() {
	code := "int x;"
	ast, err := parser.ParseString(code)
	if err != nil {
		fmt.Printf("SyntaxErr: %s\n", err)
	}
	for i := 0; i < len(ast); i++ {
		fmt.Printf("%s\n", ast[i])
	}
}
