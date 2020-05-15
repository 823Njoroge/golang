package main

import "fmt"

func main() {

	//Variable declaration model 1
	// var (
	// 	name   string = "Paulo"
	// 	age    int    = 20
	// 	height int    = 180
	// )

	//Variable declaration model 2
	// name := "Paulo"
	// age := 20
	// height := 180

	//Variable declaration model 3
	name, age, height := "Paulo", 20, 180

	fmt.Println("My name is", name, "and I have", age, "years")
	fmt.Println("My height is", height)
}
