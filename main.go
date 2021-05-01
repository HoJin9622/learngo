package main

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	fmt.Println(multiply(2, 2))
	totalLength, _ := lenAndUpper("hojin")
	fmt.Println(totalLength)

	repeatMe("hojin", "sanghyun", "go", "java")
}