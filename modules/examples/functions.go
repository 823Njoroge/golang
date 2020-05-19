// Defining Functions in the Go / Golang Language

package main

import (
	"fmt"
)

func sum(x int, y int) int {
	return x + y
}

func calc(a int) (int, int) {
	var square int = a * a
	var cube int = a * a * a

	return square, cube
}

func main() {

	fmt.Println(sum(3, 4))

	fmt.Println(calc(2))
}
