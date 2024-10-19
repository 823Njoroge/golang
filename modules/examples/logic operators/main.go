// Logic Operators (E Or and Non-Logical) (Go / Golang && ||!)

package main

import "fmt"

func main() {
	// && || !

	// 3 < x < 7
	// x > 3 && x <7

	var x int = 4
	fmt.Println("x > 3 && x < 7", x > 3 && x < 7)

	fmt.Println("x < 3 || x > 7", x < 3 || x > 7)

	fmt.Println(!false)
}
