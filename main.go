package main

import (
	"fmt"
	"strings"
)

// 상수 - 변경 불가능
// const name string = "nico"
// 변수 - 변경 가능
// var name string = "nico"
// 축약형
// name := "nico" - 한줄 위와 같은 뜻, Go가 타입을 자동으로 찾아준다. 그리고 func 안에서만 사용 가능하다.

// 여러개의 변수를 받을 수 있는 함수
func repeatMe(words ...string) {
	fmt.Println(words)
}

// naked return
func lenAndUpper(name string) (length int, uppercase string) {
	// defer는 function이 끝난 후 작동한다
	defer fmt.Println("I'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	// for i:= 0; i < len(numbers); i++ {
	// 	fmt.Println(numbers[i])
	// }
	return total
}

func canIDrink(age int) bool {
	// variable expression
	// if koreanAge := age + 2; koreanAge < 18 {
	// 	return false
	// }
	// return true

	// switch 문
	switch koreanAge:= age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

func main() {
	repeatMe("hi", "bye", "hello", "yo")
	totalLength, up := lenAndUpper("hojin")
	fmt.Println(totalLength, up)
	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)
	fmt.Println(canIDrink((15)))

	a := 2
	b := &a
	a = 5
	*b = 20
	fmt.Println(a, &a, b, &b, *b)

	// Data Structure Array
	names:= [5]string{"nico", "lynn", "dal"}
	names[3] = "lalala"
	names[4] = "llalalala"
	fmt.Println(names)

	// Data Structure Slice
	newNames := []string{"nico", "lynn", "dal"}
	// Append 는 새로운 slice를 return 한다.
	newNames = append(newNames, "flynn")

	// Map
	nico := map[string]string{"name": "nico", "age": "12"}
	fmt.Println(nico)
	for key, value := range nico {
		fmt.Println(key, value)
	}
}