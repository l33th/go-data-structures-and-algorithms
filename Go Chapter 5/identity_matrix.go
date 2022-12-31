package main

import (
	"fmt"
)

// Identity method
func Identity(order int) [][]float64 {
	var matrix [][]float64
	matrix = make([][]float64, order)
	var i int
	for i = 0; i < order; i++ {
		var temp []float64
		temp = make([]float64, order)
		temp[i] = 1
		matrix[i] = temp
	}
	return matrix
}

// main method
func main() {
	fmt.Println(Identity(4))
}
