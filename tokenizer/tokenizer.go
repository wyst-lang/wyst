package tokenizer

import (
	"fmt"
	"strings"
)

type State struct {
	Line   int
	Column int
}

type Token struct {
	Rule     Rule
	Value    string
	Position State
}

func (t Token) String() string {
	return fmt.Sprintf("Token(\n  rule=%s,\n  value=%s\n)", t.Rule, t.Value)
}

func Tokenize(code string) ([]Token, State, error) {
	var tokens = []Token{}
	var state = State{1, 0}
	for code != "" {
		matched := false
		for i := 0; i < len(RULES); i++ {
			if RULES[i].Pattern.MatchString(code) {
				match := RULES[i].Pattern.FindString(code)
				matched = true
				code = code[len(match):]
				tokens = append(tokens, Token{Rule: RULES[i], Value: match, Position: state})
				state.Column += len(match)
				break
			} else if strings.HasPrefix(code, "\n") {
				state.Line += 1
				state.Column = 0
				code = code[1:]
				matched = true
			} else if strings.HasPrefix(code, "\t") || strings.HasPrefix(code, " ") {
				state.Column += 1
				code = code[1:]
				matched = true
			}
		}
		if !matched {
			return tokens, state, fmt.Errorf("invalid character or token")
		}
	}
	return tokens, state, nil
}
