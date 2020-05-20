// Relationship / Comparison Operators in Golang

package main

import (
	"fmt"
)

func main() {
	// > < >= <=
	// == !=

	var a int = 24
	var b int = 8
	var c float64 = 6

	fmt.Println("a > b", a > b)

	fmt.Println("a < b", a < b)

	fmt.Println("a >= b", a >= b)

	fmt.Println("a <= b", a <= b)

	fmt.Println("3 ==4", 3 == 4)

	fmt.Println("3 ==3", 3 == 3)

	fmt.Println("b ==c", b == int(c))

	fmt.Println("a != b", a != b)

}
