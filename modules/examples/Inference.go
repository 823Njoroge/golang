// Type Inference in the Go Language

package main

import "fmt"

func main() {

	var i int
	var j = i // int

	var a = 36

	fmt.Printf("Type: %T Value: %v\n", j, j)

	fmt.Printf("Type: %T Value: %v\n", a, a)

}


