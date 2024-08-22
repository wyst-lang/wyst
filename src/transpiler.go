package main

import (
	"fmt"
	"strings"
)

type wmap [2]string

var wyst_map []wmap = []wmap{{"int", "int"}, {"void", "void"}, {"main", "main"}, {"string", "string"}, {"true", "true"}, {"false", "false"}}

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
		res += "I0x"
		for _, c := range node.Text {
			if c == ':' || c == '.' {
				res += string(c)
			} else {
				res += fmt.Sprintf("%x", c)
			}
		}
		res = strings.ReplaceAll(res, "::", ".I0x")
	case "HEX":
		res += node.Text
	case "MATH":
		res += node.Text
	case "STRING":
		if strings.HasPrefix(node.Text, "'") {
			res += "\"" + strings.ReplaceAll(strings.TrimSuffix(strings.TrimPrefix(node.Text, "'"), "'"), "\\'", "'") + "\""
		} else {
			res += node.Text
		}
	case "';'":
		res += "\n  "
	case "','":
		res += ", "
	case "IDENT":
		res += fmt.Sprintf("_%#x", node.Text)
	case "GOIDENTIFIER":
		res += strings.TrimPrefix(node.Text, "^")
	default:
		fmt.Printf("\x1b[33mDEV NOTICE:\x1b[0m Missing token case for '%s'\n", node.Rule)

	}

	return res
}

func (m *Module) TranspileNode(node ASTNode) string {
	res := ""
	if node.Rule == strings.ToUpper(node.Rule) {
		return TranspileToken(node)
	}
	switch node.Rule {

	case "func_def":
		return_type := m.TranspileNode(node.Inner[0])
		if return_type == "void" {
			return_type = ""
		}
		res += fmt.Sprintf("func %s%s %s %s", m.TranspileNode(node.Inner[1]), m.TranspileNode(node.Inner[2]), return_type, m.TranspileNode(node.Inner[3]))
	case "expr":
		res += m.TranspileNodes(node.Inner)
	case "code_block":
		res += fmt.Sprintf("{\n  %s}", strings.TrimRight(m.TranspileNodes(node.Inner), " "))
	case "var_def":
		res += fmt.Sprintf("var %s %s", m.TranspileNode(node.Inner[1]), m.TranspileNode(node.Inner[0]))
	case "round_def":
		res += fmt.Sprintf("(%s)", strings.TrimRight(strings.ReplaceAll(m.TranspileNodes(node.Inner), "var ", ""), " "))
	case "call_tree":
		for i, c := range node.Inner {
			res += m.TranspileNode(c)
			if i+1 < len(node.Inner) {
				res += "."
			}
		}

	case "fn_call":
		res += m.TranspileNode(node.Inner[0])
		res += m.TranspileNode(node.Inner[1])

	case "struct_def":
		res += fmt.Sprintf("type %s struct %s\n", m.TranspileNode(node.Inner[0]), m.TranspileNode(node.Inner[1]))

	case "enum_curly":
		res += fmt.Sprintf("{\n  %s}", strings.TrimRight(m.TranspileNodes(node.Inner), " "))

	case "namespace":
		namespace_name := "n" + m.TranspileNode(node.Inner[0])
		res += fmt.Sprintf("type %s struct {\n  ", namespace_name)
		vals := ""
		for _, c := range node.Inner {
			if c.Rule == "var_def" {
				res += strings.TrimPrefix(m.TranspileNode(c), "var ")
				res += "\n  "
			} else if c.Rule == "var_def_set" {
				res += strings.TrimPrefix(m.TranspileNode(c.Inner[0]), "var ")
				res += "\n  "
				vals += fmt.Sprintf("%s: %s, ", m.TranspileNode(c.Inner[0].Inner[1]), m.TranspileNode(c.Inner[1]))
			}
		}
		res += "}\n"
		for _, c := range node.Inner {
			if c.Rule == "func_def" {
				return_type := m.TranspileNode(c.Inner[0])
				if return_type == "void" {
					return_type = ""
				}
				res += fmt.Sprintf("func (%s) %s%s %s %s", namespace_name, m.TranspileNode(c.Inner[1]), m.TranspileNode(c.Inner[2]), return_type, m.TranspileNode(c.Inner[3]))
				res += "\n  "
			}
		}
		res += fmt.Sprintf("var %s %s = %s{%s}", m.TranspileNode(node.Inner[0]), namespace_name, namespace_name, vals)

	case "var_def_set":
		res += fmt.Sprintf("%s = %s", m.TranspileNode(node.Inner[0]), m.TranspileNode(node.Inner[1]))

	case "import_statement":
		res += fmt.Sprintf("import . \"%s\"\n", m.ImportModule(node.Inner[0].Text))

	case "use_statement":
		ident := strings.Split(node.Inner[0].Text, "::")
		ident2 := strings.Split(m.TranspileNode(node.Inner[0]), ".")
		modname := m.ImportModule(ident[0])
		md := strings.Split(modname, "/")
		res += fmt.Sprintf("import %s \"%s\"\n", md[len(md)-1], modname)
		if len(ident) > 1 {
			res += fmt.Sprintf("var %s = %s\n", ident2[len(ident2)-1], m.TranspileNode(node.Inner[0]))
		}

	case "go_import":
		res += "import " + m.TranspileNode(node.Inner[0]) + "\n"

	case "map":
		wyst_map = append(wyst_map, wmap{node.Inner[0].Text, node.Inner[2].Text})

	case "go_call":
		res += m.TranspileNode(node.Inner[0])
		res += m.TranspileNode(node.Inner[1])

	case "round_call":
		res += "(" + m.TranspileNodes(node.Inner) + ")"

	case "if_tree":
		res += m.TranspileNodes(node.Inner)

	case "if_statement":
		res += "if " + m.TranspileNodes(node.Inner)

	case "if_expr":
		res += m.TranspileNodes(node.Inner)

	case "else_statement":
		res += "else "
		res += m.TranspileNodes(node.Inner)

	case "elseif_statement":
		res += "else "
		res += m.TranspileNodes(node.Inner)

	default:
		fmt.Printf("\x1b[33mDEV NOTICE:\x1b[0m Missing case for '%s'\n", node.Rule)

	}
	return res
}

func (m *Module) TranspileNodes(nodes []ASTNode) string {
	res := ""
	for i := 0; i < len(nodes); i++ {
		res += m.TranspileNode(nodes[i])
	}
	return res
}

func (m *Module) Transpile(code string) string {
	ast := Parse(code)
	res := m.TranspileNodes(ast.Inner)
	return res
}
