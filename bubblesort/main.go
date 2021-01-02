package bubblesort

// Bubblesort
//
// 1. Compare each value with the next in the array
// 2. If the previous value is greater than the next, swap the values
// 3. Continue along the array toward the end, repeating 1 and 2
// 4. Repeat the process, the high values will bubble to the top with each
//    successive iteration.
//
// The time complexity of bubble sort is O(n2).

// Sort an array using the bubblesort algorithm
func Sort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 1; j < len(arr)-i; j++ {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
	return arr
}
