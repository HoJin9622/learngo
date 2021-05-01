package main

import "fmt"

func main() {
	const name string = "hojin"
	fmt.Println(name)

	var nickname string = "nico"
	nickname = "coco"
	fmt.Println(nickname)

	score := 20 // 축약형은 오로지 func 안에서만 사용 가능
	score = 40
	fmt.Println(score)
}