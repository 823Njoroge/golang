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

// null return
func calc2(a int) (square int, cube int) {
	square = a * a
	cube = a * a * a

	return
}

func main() {

	fmt.Println(sum(3, 4))

	fmt.Println(calc(2))

	fmt.Println(calc2(4))
}
