package main

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/wyst-lang/wyst/parser"
)

func Parse(code string) (antlr.ParseTree, *parser.WystParser) {
	chars := antlr.NewInputStream(code)
	lexer := parser.NewWystLexer(chars)
	stream := antlr.NewCommonTokenStream(lexer, 0)

	p := parser.NewWystParser(stream)
	// p.AddErrorListener(antlr.NewDiagnosticErrorListener(true).WithContext(p))

	tree := p.Expr()
	return tree, p
}

func iterateTree(node antlr.Tree, wparser *parser.WystParser, depth int) {
	if ruleContext, ok := node.(antlr.RuleContext); ok {
		ruleIndex := ruleContext.GetRuleIndex()
		ruleName := wparser.RuleNames[ruleIndex]
		fmt.Printf("%sRule: %s\n", indent(depth), ruleName)
	}

	if parseTree, ok := node.(antlr.ParseTree); ok {
		fmt.Printf("%sNode: %s\n", indent(depth), parseTree.GetText())
	}

	for i := 0; i < node.GetChildCount(); i++ {
		child := node.GetChild(i)
		iterateTree(child, wparser, depth+1)
	}
}

func indent(depth int) string {
	return fmt.Sprintf("%s", string(make([]byte, depth*2)))
}
