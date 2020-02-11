package main

import "fmt"

func main() {
	// Define a map
	emails := make(map[string]string)

	// Assign key value pairs
	emails["deno"] = "deno@gmail.com"
	emails["sonya"] = "sonya@gmail.com"
	emails["jakes"] = "jakes@gmail.com"

	// Declare map and assign
	locations := map[string]string{"Kawasaki": "London", "Mitsubishi": "Germany"}

	fmt.Println(locations)

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["deno"])

	// Delete a pair
	delete(emails, "jakes")
	fmt.Println(emails)

}
