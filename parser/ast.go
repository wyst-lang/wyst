package parser

import (
	"fmt"
)

type RuleKind int

const (
	RK_Required RuleKind = iota
	RK_Optional
	RK_Repeat
	RK_Port
)

type Rule struct {
	Name  string
	Kind  RuleKind
	Inner []Rule
}

type AstNode struct {
	Rule  Rule
	Value string
	Inner []AstNode
}

func (t AstNode) String() string {
	return fmt.Sprintf("AstNode(\n  value=%s\n  rule=%v,\n  inner=%v,\n)", t.Value, t.Rule, t.Inner)
}
func (t Rule) String() string {
	return t.Name
}

func NewRule(name string, inner []Rule) Rule {
	return Rule{name, RK_Required, inner}
}

func NewPort(name string, inner []Rule) Rule {
	return Rule{name, RK_Port, inner}
}

func Optional(rule Rule) Rule {
	rule.Kind = RK_Optional
	return rule
}

func Repeat(rule Rule) Rule {
	rule.Kind = RK_Repeat
	return rule
}

var (
	IDENTIFIER = NewPort("IDENTIFIER", []Rule{})
	NUMBER     = NewPort("NUMBER", []Rule{})
	SEMICOLON  = NewPort("SEMICOLON", []Rule{})
	EXPR       = NewRule("EXPR", []Rule{IDENTIFIER})
	CODE_BLOCK = NewRule("CODE_BLOCK", []Rule{EXPR, SEMICOLON})
	VAR_DEF    = NewRule("VAR_DEF", []Rule{IDENTIFIER, IDENTIFIER})
	FUNC_DEF   = NewRule("FUNC_DEF", []Rule{IDENTIFIER, IDENTIFIER, CODE_BLOCK})
)

var (
	TOP_RULE = []Rule{FUNC_DEF}
)
