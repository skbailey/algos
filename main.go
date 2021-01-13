package main

import (
	"algos/bubblesort"
	"algos/insertionsort"
	"algos/mergesort"
	"algos/quicksort"
	"algos/structures"
	"fmt"
)

func main() {
	stack := structures.NewStack()
	stack.Push("this")
	stack.Push("is")
	stack.Push("sparta!!")

	for len(stack) > 0 {
		x, y := stack.Pop()
		if y == true {
			fmt.Println(x)
		}
	}

	unsortedArray := []int{5, 7, 2, -3, 9, -5}
	fmt.Println("Unsorted Array", unsortedArray)

	// Quick Sort
	sortedArray := quicksort.Sort(unsortedArray, 0, len(unsortedArray)-1)
	fmt.Println("quick sort complete", sortedArray)

	// Merge Sort
	unsortedArray = []int{5, 7, 2, -3, 9, -5}
	sortedArray = mergesort.Sort(unsortedArray)
	fmt.Println("merge sort complete", sortedArray)

	// Insertion Sort
	unsortedArray = []int{5, 7, 2, -3, 9, -5}
	sortedArray = insertionsort.Sort(unsortedArray)
	fmt.Println("insertion sort complete", sortedArray)

	// Bubble Sort
	unsortedArray = []int{5, 7, 2, -3, 9, -5}
	sortedArray = bubblesort.Sort(unsortedArray)
	fmt.Println("bubble sort complete", sortedArray)
}
