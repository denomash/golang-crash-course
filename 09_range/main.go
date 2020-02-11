package main

import "fmt"

func main() {
	nums := []int{32, 35, 67, 732, 268, 332, 48}

	for i, num := range nums {
		fmt.Printf("Index %d - num: %d\n", i, num)
	}

	// Not using index
	for _, num := range nums {
		fmt.Printf("num: %d\n", num)
	}

	// Add numbers tongether
	sum := 0
	for _, num := range nums {
		sum += num
	}

	fmt.Println("Sum", sum)

	// Range with map
	locations := map[string]string{"Kawasaki": "London", "Mitsubishi": "Germany"}

	for key, value := range locations {
		fmt.Printf("%s: %s\n", key, value)
	}

	for key := range locations {
		fmt.Println("Model: ", key)
	}
}
