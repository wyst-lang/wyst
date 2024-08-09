package token

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

func NewToken(name string, pattern regexp.Regexp) Rule {
	return Rule{Name: name, Pattern: pattern}
}

var (
	NUMBER     = NewToken("NUMBER", *regexp.MustCompile("^[0-9]*"))
	IDENTIFIER = NewToken("IDENTIFIER", *regexp.MustCompile("^[A-Za-z_][A-Za-z_0-9]*"))
)

var RULES_TOP = []Rule{IDENTIFIER, NUMBER}
