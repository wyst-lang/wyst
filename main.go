package main

import (
	"fmt"

	"github.com/wyst-lang/wyst/parser"
)

func main() {
	code := "int x"
	ast, state, err := parser.Parse(code, parser.RULE_SET)
	if err != nil {
		fmt.Printf("SyntaxErr at %d:%d: %s\n", state.Line, state.Column, err)
		return
	}
	for i := 0; i < len(ast); i++ {
		fmt.Printf("%s\n", ast[i])
	}
}
