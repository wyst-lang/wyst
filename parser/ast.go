package parser

import (
	"fmt"

	"github.com/wyst-lang/wyst/tokenizer"
)

type inner interface {
	isInner()
}

type Rule struct {
	Name  string
	Inner inner
}

func (t Rule) String() string {
	return fmt.Sprintf("AstRule(%s)", t.Name)
}

type Rules struct {
	Value []Rule
}

func (v Rules) isInner() {}

type Token struct {
	Value tokenizer.Rule
}

func (v Token) isInner() {}

type String struct {
	Value string
}

func (v String) isInner() {}

type AstNode struct {
	Rule  Rule
	Value string
	Inner []AstNode
}

func (t AstNode) String() string {
	return fmt.Sprintf("AstNode(\n  rule=%s,\n  value=%s,\n  inner=%v\n)", t.Rule, t.Value, t.Inner)
}

// Ported rules from the tokenizer
var (
	IDENTIFIER = Rule{"IDENTIFIER", Token{tokenizer.IDENTIFIER}}
)

var (
	TOP = Rule{"TOP", Rules{
		[]Rule{
			IDENTIFIER,
		}}}
)
