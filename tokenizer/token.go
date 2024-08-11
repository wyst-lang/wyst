package tokenizer

import (
	"fmt"
	"regexp"
)

type Rule struct {
	Pattern regexp.Regexp
	Name    string
}

func (t Rule) String() string {
	return fmt.Sprintf("TokenRule(%s)", t.Name)
}

var RULES = []Rule{}

func NewRule(name string, pattern regexp.Regexp) Rule {
	var rule = Rule{Name: name, Pattern: pattern}
	RULES = append(RULES, rule)
	return rule
}

var (
	NUMBER     = NewRule("NUMBER", *regexp.MustCompile(`^\d+`))
	IDENTIFIER = NewRule("IDENTIFIER", *regexp.MustCompile("^[A-Za-z_][A-Za-z_0-9]*"))
	SEMICOLON  = NewRule("SEMICOLON", *regexp.MustCompile("^;"))
)
