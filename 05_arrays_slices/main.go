package main

import "fmt"

func main() {
	// Arrays (Are arrays with a fixed number)
	var osArr [2]string

	// Assign values
	osArr[0] = "Mac OS"
	osArr[1] = "Ubuntu"

	// Declare and assign
	carArr := [3]string{"Rolls Roys", "Aston Martin", "Tesla"}

	// Slices (Are arrays without a fixed number)
	carSlice := []string{"Rolls Roys", "Aston Martin", "Ferrari", "Tesla"}

	fmt.Println(osArr)
	fmt.Println(osArr[0])
	fmt.Println(carArr)
	fmt.Println(carSlice)

	// Length of an array
	fmt.Println(len(carSlice))

	// Get range
	fmt.Println(carSlice[0:2])

}
