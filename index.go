package main

import (
	"fmt"
)

// entry point of app
func main() {
	fmt.Println("Hello world!")
	secFunc()
	hello := "Hello world"
	fmt.Println(len("Hello World!"))
	fmt.Println("Hello World!")
	fmt.Println(hello)
	fmt.Println(true || false)

	var x string = "hello"
	var y string = "hello"

	fmt.Println(x == y)

	z := 10
	z += 1
	fmt.Println(z)

	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			fmt.Println("odd")
		} else {
			fmt.Println("even")
		}
	}
}

func secFunc() {
	fmt.Println("Go is cool")
}
