// Repetition Structure While (While Loop) in the Go Language

package main

import (
	"fmt"
)

func main() {

	// We want sum: 2 + 4 + 8 + 16 + ...
	// 2^1 + 2^2 + 2^3 + 2^4 + ...

	var sum int = 2

	for sum < 600 {
		sum = sum + sum
	}

	// Curiosity: how to make an infinite loop
	// for true { //while
	// 	sum += sum
	// 	fmt.Println(sum)
	// 	time.Sleep(100 * time.Millisecond)

	// }

	fmt.Println(sum)
}
