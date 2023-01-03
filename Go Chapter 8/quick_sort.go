package main

import (
	"fmt"
)

// QuickSorter method
func QuickSorter(elements []int, below int, upper int) {
	if below < upper {
		var part int
		part = divideParts(elements, below, upper)
		QuickSorter(elements, below, part-1)
		QuickSorter(elements, part+1, upper)
	}
}

// divideParts method
func divideParts(elements []int, below int, upper int) int {
	var center int
	center = elements[upper]
	var i int
	i = below
	var j int
	for j = below; j < upper; j++ {
		if elements[j] <= center {
			swap(&elements[i], &elements[j])
			i += 1
		}
	}
	swap(&elements[i], &elements[upper])
	return i
}

// swap method
func swap(element1 *int, element2 *int) {
	var val int
	val = *element1
	*element1 = *element2
	*element2 = val
}

// main method
func main() {
	var num int

	fmt.Print("Enter Number of Elements: ")
	_, err := fmt.Scan(&num)
	if err != nil {
		return
	}

	var array = make([]int, num)

	var i int
	for i = 0; i < num; i++ {
		fmt.Print("array[", i, "]: ")
		_, err := fmt.Scan(&array[i])
		if err != nil {
			return
		}
	}

	fmt.Print("Elements: ", array, "\n")
	QuickSorter(array, 0, num-1)
	fmt.Print("Sorted Elements: ", array, "\n")
}
