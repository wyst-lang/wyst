package parser

import (
	"fmt"

	"github.com/wyst-lang/wyst/tokenizer"
)

func ParseString(code string) ([]AstNode, error) {
	tokens, state, err := tokenizer.Tokenize(code)
	ast, err1 := Parse(tokens, TOP_RULE)
	if err != nil {
		return ast, fmt.Errorf("LexingError at %d:%d: %s", state.Line, state.Column, err)
	} else if err1 != nil {
		return ast, fmt.Errorf("ParserError")
	}
	return ast, nil
}

func Parse(tokens []tokenizer.Token, rule_set []Rule) ([]AstNode, error) {
	var ast = []AstNode{}
	for i := 0; i < len(tokens); i++ {
		for r := 0; r < len(rule_set); r++ {
			matching := false
			var match_nodes = []AstNode{}
			if rule_set[r].Kind == RK_Port {
				if tokens[i].Rule.Name == rule_set[r].Name {
					matching = true
					match_nodes = append(match_nodes, AstNode{Rule: rule_set[r], Value: tokens[i].Value})
				}
			} else {
				for v := 0; v < len(rule_set[r].Inner); v++ {
					if rule_set[r].Inner[v].Kind == RK_Required {
						if len(tokens)-i > v && tokens[i].Rule.Name == rule_set[r].Inner[v].Name {
							matching = true
							match_nodes = append(match_nodes, AstNode{Rule: rule_set[r].Inner[v], Value: tokens[i].Value})
						} else {
							matching = false
							break
						}
					} else if rule_set[r].Inner[v].Kind == RK_Repeat {
						for l := 0; len(tokens)-i > l && tokens[i+l].Rule.Name == rule_set[r].Inner[v].Name; l++ {
							match_nodes = append(match_nodes, AstNode{Rule: rule_set[r].Inner[v], Value: tokens[i+l].Value})
							matching = true
						}
					} else if rule_set[r].Inner[v].Kind == RK_Optional {
						match_nodes = append(match_nodes, AstNode{Rule: rule_set[r].Inner[v], Value: tokens[i].Value})
						matching = true
					}
				}
			}
			if matching {
				ast = append(ast, AstNode{Rule: rule_set[r], Inner: match_nodes, Value: merge_nodes(match_nodes)})
				i += len(match_nodes) - 1
				break
			}
		}
	}
	return ast, nil
}

func merge_nodes(match_nodes []AstNode) string {
	str := ""
	for i := 0; i < len(match_nodes); i++ {
		str += match_nodes[i].Value
		if len(match_nodes) > i {
			str += " "
		}
	}
	return str
}
