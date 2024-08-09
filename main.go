package main

import (
	"fmt"

	"github.com/wyst-lang/wyst/token"
)

func main() {
	tokens, err := Tokenize("test 123", token.RULES_TOP)
	if err != nil {
		fmt.Printf("Syntax error at %s \n", err)
		return
	}
	for i := 0; i < len(tokens); i++ {
		fmt.Printf("%s\n", tokens[i])
	}
}
