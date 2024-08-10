package parser

import (
	"fmt"
)

type Rule struct {
	Name  string
	Inner []Rule
}

type AstNode struct {
	Rule  Rule
	Inner []AstNode
}

func (t AstNode) String() string {
	return fmt.Sprintf("AstNode(\n  rule=%v,\n  inner=%v,\n)", t.Rule, t.Inner)
}
func (t Rule) String() string {
	return t.Name
}

func NewRule(name string, inner []Rule) Rule {
	return Rule{name, inner}
}

var (
	IDENTIFIER = NewRule("IDENTIFIER", []Rule{})
	VAR_DEF    = NewRule("VAR_DEF", []Rule{IDENTIFIER, IDENTIFIER})
)

var (
	RULE_SET = []Rule{IDENTIFIER, VAR_DEF}
)
