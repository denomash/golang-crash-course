package main

import (
	"fmt"
	"math"

	util "github.com/denomash/go_crash_coarse/03_packages/utils"
)

func main() {
	fmt.Println(math.Floor(4.5))
	fmt.Println(math.Ceil(3.9))
	fmt.Println(math.Sqrt(25))
	fmt.Println(util.Reverse("denomash"))
}
