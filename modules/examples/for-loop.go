// Traditional Repetition Structure (For Loop) in the Go Language

package main

import "fmt"

func main() {

	// add integer numbers from 1 to 10

	var sum int = 0

	// for startup; condition; step

	for i := 1; i <= 10; i++ {
		fmt.Println("The value of i is:", i)
		// sum = sum + i
		sum += i
	}
	fmt.Println("The sum of the integers is:", sum)
}
