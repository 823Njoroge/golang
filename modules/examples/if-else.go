// The If Else Declaration in the Go Language

package main

import "fmt"

func main() {
	// if else and else if

	var a int = 8

	// a > 10: "The a is greater than 10"

	// a <=10 && a >5: "a is between 6 and 10"

	// a <=5: "a is less than or equal to 10"

	if a > 10 {
		fmt.Println("The a is greater than 10")
	} else if a > 5 {
		fmt.Println("a is between 6 and 10")
	} else {
		fmt.Println("a is less than or equal to 10")
	}
}
