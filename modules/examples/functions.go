// Defining Functions in the Go / Golang Language

package main

import (
	"fmt"
)

func sum(x int, y int) int {
	return x + y
}

func main() {

	fmt.Println(sum(3, 4))
}
