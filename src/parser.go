package main

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/wyst-lang/wyst/parser"
)

type ASTNode struct {
	Rule  string
	Text  string
	Inner []ASTNode
}

func getTokenName(tokenType int, lexer *parser.WystLexer) string {
	if tokenType >= 0 && tokenType < len(lexer.SymbolicNames) {
		if lexer.SymbolicNames[tokenType] != "" {
			return lexer.SymbolicNames[tokenType]
		}
		return lexer.LiteralNames[tokenType]
	}
	return "UNKNOWN"
}

func ConvertAST(node antlr.Tree, wparser *parser.WystParser, lexer *parser.WystLexer) ASTNode {
	var ast = ASTNode{}
	switch node := node.(type) {
	case antlr.TerminalNode:
		token := node.GetSymbol()
		ast.Text = token.GetText()
		ast.Rule = getTokenName(token.GetTokenType(), lexer)
	default:
		if ruleContext, ok := node.(antlr.RuleContext); ok {
			ast.Rule = wparser.RuleNames[ruleContext.GetRuleIndex()]
		}
		if parseTree, ok := node.(antlr.ParseTree); ok {
			ast.Text = parseTree.GetText()
		}
		for i := 0; i < node.GetChildCount(); i++ {
			l := ConvertAST(node.GetChild(i), wparser, lexer)
			if l.Rule != ("'"+l.Text+"'") || l.Text == ";" {
				ast.Inner = append(ast.Inner, l)
			}
		}
	}
	return ast
}

func Parse(code string) ASTNode {
	chars := antlr.NewInputStream(code)
	lexer := parser.NewWystLexer(chars)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewWystParser(stream)
	// p.AddErrorListener(antlr.NewDiagnosticErrorListener(true).WithContext(p))
	tree := p.Top()
	return ConvertAST(tree, p, lexer)
}
