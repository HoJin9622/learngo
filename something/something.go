package something

import "fmt"

// private fuction
func sayBye() {
	fmt.Println("Bye")
}

// export 된 function
func SayHello() {
	fmt.Println("Hello")
}