package main

import (
	"fmt"
)

// Factor method
func Factor(num int) int {
	if num <= 1 {
		return 1
	}
	return num * Factor(num-1)
}

//main method
func main() {
	var num int
	num = 5
	fmt.Printf("Factorial: %d is %d\n", num, Factor(num))
}
