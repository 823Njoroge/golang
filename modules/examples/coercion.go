// Type Conversion (Coercion) in the Go / Golang Language

package main

import "fmt"

func main() {
	var a int = 64
	var b float64 = 6.4

	var c int =  int (b)
	var d float64 = float64(a)

	fmt.Println("The value of c is", c)

	fmt.Println("The value of d is", d)
}