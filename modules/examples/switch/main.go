// Switch in the Go Language

package main

import "fmt"

func main() {

	name := "Paul"
	grade := 10

	switch name {
	case "Ana":
		fmt.Println("It's Ana")
	case "John":
		fmt.Println("It's John")
	default:
		fmt.Println("I don't know the person")
	}

	// Switch True in the Go Language

	switch true {
	case grade > 6:
		fmt.Println("Approved")
	case grade <= 5:
		fmt.Println("Disapproved")
	}

}
