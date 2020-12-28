package main

import (
	"algos/mergesort"
	"fmt"
)

func main() {
	unsortedArray := []int{5, 7, 2, -3, 9, -5}
	fmt.Println("Unsorted Array", unsortedArray)
	// sortedArray := quicksort.Sort(unsortedArray, 0, len(unsortedArray)-1)
	sortedArray := mergesort.Sort(unsortedArray)
	fmt.Println("Sorted Array", sortedArray)
}
