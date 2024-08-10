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

func NewRule(name string, pattern regexp.Regexp) Rule {
	return Rule{Name: name, Pattern: pattern}
}

var (
	NUMBER     = NewRule("NUMBER", *regexp.MustCompile(`^\d+`))
	IDENTIFIER = NewRule("IDENTIFIER", *regexp.MustCompile("^[A-Za-z_][A-Za-z_0-9]*"))
)

var RULES = []Rule{IDENTIFIER, NUMBER}
