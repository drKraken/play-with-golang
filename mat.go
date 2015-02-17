package main

import (
	"fmt"
)

func main() {
	var radius float32

	fmt.Scanf("%f", &radius)

	length := 2 * 3.14 * radius

	fmt.Println("Length: ", length)
}
