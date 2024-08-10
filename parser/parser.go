package parser

import (
	"github.com/wyst-lang/wyst/tokenizer"
)

func Parse(code string, rule_set []Rule) ([]AstNode, tokenizer.State, error) {
	var ast = []AstNode{}
	tokens, state, err := tokenizer.Tokenize(code)
	for i := 0; i < len(tokens); i++ {
		for r := 0; r < len(rule_set); r++ {
			matching := false
			var match_nodes = []AstNode{}
			for v := 0; v < len(rule_set[r].Inner); v++ {
				if len(tokens)-i > v && tokens[i].Rule.Name == rule_set[r].Inner[v].Name {
					matching = true
					match_nodes = append(match_nodes, AstNode{Rule: rule_set[r].Inner[v]})
				} else {
					matching = false
					break
				}
			}
			if matching {
				ast = append(ast, AstNode{Rule: rule_set[r], Inner: match_nodes})
				i += len(match_nodes) - 1
				break
			}
		}
	}
	if err != nil {
		return ast, state, err
	}
	return ast, state, nil
}
