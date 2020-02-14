package main

import "fmt"

func main() {
	a := 10

	b := &a

	fmt.Println(a, b)

	// check type
	fmt.Printf("%T\n", b)

	// Use * to read value from address
	fmt.Println(*b)
	fmt.Println(*&a)

	// Change value using *
	*b = 20
	fmt.Println(a)

}
