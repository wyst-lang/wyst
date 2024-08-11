package main

import (
)

func main() {
	tree, parser := Parse("")
	iterateTree(tree, parser, 0)
}
