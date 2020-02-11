package main

import "fmt"

func main() {
	x := 100
	y := 36

	// if else
	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)

	} else {
		fmt.Printf("%d is less than %d\n", y, x)

	}

	// else if
	car := "Ferrari"

	if car == "Mercedes" {
		fmt.Println("The model is Mercedes")
	} else if car == "Tesla" {
		fmt.Println("The model is Tesla")
	} else {
		fmt.Println("The model is NOT Mercedes or Tesla")

	}

	// Switch statement
	switch car {
	case "Ferrari":
		fmt.Println("The model is Ferrari")

	case "Mercedes":
		fmt.Println("The model is Mercedes")

	default:
		fmt.Println("The model is NOT Mercedes or Ferrari")

	}

}
