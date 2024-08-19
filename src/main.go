package main

import (
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("test/main.wst")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	code := Transpile(string(content))
	fmt.Printf("output:\n%s\n", code)
}
