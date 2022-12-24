package main

import (
	"fmt"
)

//gets the power series of integer a and returns
//tuple of square of a and cube of a
// powerSeries method
func powerSeries(a int) (int, int) {
	return a * a, a * a * a
}

// powerSeriesN method
func powerSeriesN(a int) (square int, cube int) {
	square = a * a
	cube = square * a
	return
}

// powerSeriesE method
func powerSeriesE(a int) (int, int, error) {
	var square int
	square = a * a
	var cube int
	cube = square * a
	return square, cube, nil
}

// main method
func main() {
	var square int
	var cube int
	square, cube = powerSeries(3)
	fmt.Println("Square ", square, "Cube", cube)
	fmt.Println(powerSeriesN(4))
	fmt.Println(powerSeriesE(5))
}
