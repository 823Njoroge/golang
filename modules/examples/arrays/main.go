// Learn Arrays in the Go Language

package main

import "fmt"

func main() {

	// Arrays
	// 1, 1, 2, 3, 5, 8, 14

	var numbers [7]int
	numbers[0] = 1
	numbers[1] = 1
	numbers[2] = 2
	numbers[3] = 3
	numbers[4] = 5
	numbers[5] = 8
	numbers[6] = 14

	fmt.Println(numbers)

	// Other way
	var x = [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(x)

	// Other way 2
	y := []int{1, 5, 10, 15, 20, 25, 30}
	fmt.Println(y)

	// How to Access Array Elements and Use For Loop to Show Each Value
	for index := 0; index < len(numbers); index++ {
		fmt.Println(numbers[index])
	}

	// How to Sum All Integers in an Array
	var sum = 0

	for index := 0; index < len(numbers); index++ {
		sum = sum + numbers[index]
	}

	fmt.Println(sum)

}
