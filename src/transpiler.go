package main

import (
	"fmt"
	"strings"
)

type wmap [2]string

var wyst_map []wmap

func TranspileToken(node ASTNode) string {
	res := ""
	switch node.Rule {

	case "NUMBER":
		res += node.Text
	case "IDENTIFIER":
		for _, i := range wyst_map {
			if node.Text == i[0] {
				return i[1]
			}
		}
		res += "_0x"
		for _, c := range node.Text {
			if c == ':' || c == '.' {
				res += string(c)
			} else {
				res += fmt.Sprintf("%x", c)
			}
		}
	case "HEX":
		res += node.Text
	case "MATH":
		res += node.Text
	case "STRING":
		res += node.Text
	case "';'":
		res += "\n  "
	default:
		fmt.Printf("\x1b[33mDEV NOTICE:\x1b[0m Missing token case for '%s'\n", node.Rule)

	}

	return res
}

func TranspileNode(node ASTNode) string {
	res := ""
	if node.Rule == strings.ToUpper(node.Rule) {
		return TranspileToken(node)
	}
	switch node.Rule {

	case "func_def":
		res += fmt.Sprintf("func %s%s %s %s", TranspileNode(node.Inner[1]), TranspileNode(node.Inner[2]), TranspileNode(node.Inner[0]), TranspileNode(node.Inner[3]))
	case "expr":
		res += TranspileNodes(node.Inner)
	case "code_block":
		res += "{\n  "
		res += strings.TrimSuffix(TranspileNodes(node.Inner), " ")
		res += "}\n"
	case "var_def":
		res += fmt.Sprintf("var %s %s", TranspileNode(node.Inner[1]), TranspileNode(node.Inner[0]))
	case "round_def":
		res += "("
		res += strings.TrimSuffix(strings.ReplaceAll(TranspileNodes(node.Inner), "var ", ""), " ")
		res += ")"
	case "asm":
		if node.Inner[0].Text == "map" && len(node.Inner) == 4 {
			wyst_map = append(wyst_map, wmap{node.Inner[1].Text, node.Inner[2].Text})
		}

	default:
		fmt.Printf("\x1b[33mDEV NOTICE:\x1b[0m Missing case for '%s'\n", node.Rule)

	}
	return res
}

func TranspileNodes(nodes []ASTNode) string {
	res := ""
	for i := 0; i < len(nodes); i++ {
		res += TranspileNode(nodes[i])
	}
	return res
}

func Transpile(code string) string {
	ast := Parse(code)
	res := TranspileNodes(ast.Inner)
	// fmt.Printf("%v", ast)
	return res
}