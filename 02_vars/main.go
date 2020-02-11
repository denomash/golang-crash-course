package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int  int8  int16  int32  int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// var name = "dennis"
	var age = 24
	const isFit = true

	// Shorthand
	name := "Dennis Macharia"
	size := 2.33784

	movie, genre := "6Underground", "Action"

	// print line
	fmt.Println(name, age, movie, genre)

	// Get variable type
	fmt.Printf("%T\n", size)

}
