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
		res = strings.ReplaceAll(res, "::", "::_0x")
	case "HEX":
		res += node.Text
	case "MATH":
		res += node.Text
	case "STRING":
		res += node.Text
	case "';'":
		res += "\n  "
	case "','":
		res += ", "
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
		res += fmt.Sprintf("{\n  %s}", strings.TrimSuffix(TranspileNodes(node.Inner), " "))
	case "var_def":
		res += fmt.Sprintf("var %s %s", TranspileNode(node.Inner[1]), TranspileNode(node.Inner[0]))
	case "round_def":
		fmt.Sprintf("(%s)", strings.TrimSuffix(strings.ReplaceAll(TranspileNodes(node.Inner), "var ", ""), " "))
	case "power_keyword_call":
		if node.Inner[0].Text == "map" && len(node.Inner) == 4 {
			wyst_map = append(wyst_map, wmap{node.Inner[1].Text, node.Inner[2].Text})
		}
	case "call_tree":
		for i, c := range node.Inner {
			res += TranspileNode(c)
			if i+1 < len(node.Inner) {
				res += "."
			}
		}

	case "fn_call":
		res += TranspileNode(node.Inner[0])
		res += TranspileNode(node.Inner[1])

	case "round_call":
		res += fmt.Sprintf("(%s)", TranspileNodes(node.Inner))

	case "struct_def":
		res += fmt.Sprintf("type %s struct %s\n", TranspileNode(node.Inner[0]), TranspileNode(node.Inner[1]))

	case "enum_curly":
		res += fmt.Sprintf("{\n  %s}", strings.TrimSuffix(TranspileNodes(node.Inner), " "))

	case "namespace":
		namespace_name := "n" + TranspileNode(node.Inner[0])
		res += fmt.Sprintf("type %s struct {\n  ", namespace_name)
		vals := ""
		for _, c := range node.Inner {
			if c.Rule == "var_def" {
				res += strings.TrimPrefix(TranspileNode(c), "var ")
				res += "\n  "
			} else if c.Rule == "var_def_set" {
				res += strings.TrimPrefix(TranspileNode(c.Inner[0]), "var ")
				res += "\n  "
				vals += fmt.Sprintf("%s: %s, ", TranspileNode(c.Inner[0].Inner[1]), TranspileNode(c.Inner[1]))
			}
		}
		res += "}\n"
		for _, c := range node.Inner {
			if c.Rule == "func_def" {
				res += fmt.Sprintf("func (%s) %s%s %s %s", namespace_name, TranspileNode(c.Inner[1]), TranspileNode(c.Inner[2]), TranspileNode(c.Inner[0]), TranspileNode(c.Inner[3]))
				res += "\n  "
			}
		}
		res += fmt.Sprintf("var %s %s = %s{%s}", TranspileNode(node.Inner[0]), namespace_name, namespace_name, vals)
	
	case "var_def_set":
		res += fmt.Sprintf("%s = %s", TranspileNode(node.Inner[0]), TranspileNode(node.Inner[1]))

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
	return res
}
