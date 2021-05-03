package main

import (
	"fmt"
	"strings"
)

// naked return
func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func main() {
	totalLength, up := lenAndUpper("hojin")
	fmt.Println(totalLength, up)
}