package main

import (
	"algos/bubblesort"
	"algos/fibonacci"
	"algos/insertionsort"
	"algos/mergesort"
	"algos/quicksort"
	"algos/structures"
	"fmt"
)

func main() {
	// Fibonacci Sequence
	ordinal := 6
	value := fibonacci.Run(ordinal)
	fmt.Println("value", value)
	fmt.Printf("This is the %dth element of the fibonacci sequence: %d\n", ordinal, value)

	// Linked List
	list := structures.NewLinkedList()
	list.Append(3)
	list.Append(5)
	list.Append(9)
	list.Append(1)
	list.Append(1)
	list.Traverse()

	// Stack
	stack := structures.NewStack()
	stack.Push("this")
	stack.Push("is")
	stack.Push("sparta!!")

	for !stack.IsEmpty() {
		x, err := stack.Pop()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(x)
		}
	}

	// Queue
	queue := structures.NewQueue()
	queue.Add("this")
	queue.Add("queue")
	queue.Add("is")
	queue.Add("boss!")

	for !queue.IsEmpty() {
		x, err := queue.Remove()
		if err != nil {
			fmt.Println(err)
		} else {
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
