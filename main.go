package main

import "fmt"

func main() {
	ast := Parse("abc123::hello")
	fmt.Printf("%v", ast)
}
