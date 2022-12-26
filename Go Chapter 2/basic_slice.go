package main

import (
	"fmt"
)

// main method
func main() {
	var slice = []int{1, 3, 5, 6}
	slice = append(slice, 8)
	fmt.Println("Capacity", cap(slice))
	fmt.Println("Length", len(slice))
}
