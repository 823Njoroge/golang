//Integer types
//int int8 int16 in32 int64

//Integer types without negative numbers
//uint uint8 uint16 uin32 uint64

//Text types
//string

//Float Types
//float64 float 32

//Boolean type
//bool

//Complex number types
//complex64 complex128 a + bi

//Byte types
//byte rune

// Print the type of the variable and its value
package main

import "fmt"

func main() {
height := 1.88
american := true

fmt.Printf("Type: %T Value: %v\n", height, height )

fmt.Printf("Type: %T Value: %v\n", american, american )

}