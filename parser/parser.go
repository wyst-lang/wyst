package parser

import (
	"github.com/wyst-lang/wyst/tokenizer"
)

func Parse(code string, rule Rule) ([]AstNode, tokenizer.State, error) {
	tokens, state, _ := tokenizer.Tokenize(code)
	var ast = []AstNode{}
	for i := 0; i < len(tokens); i++ {
		// token := tokens[i]
		if rule.Inner == (String{""}) {
			
		}
	}
	return ast, state, nil
}
