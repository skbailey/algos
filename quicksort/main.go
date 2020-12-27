package quicksort

// Quicksort:
//
// 1. Pick a pivot (in this case the highest index in the array)
// 2. Partition the array (items less than pivot to the left)
// 3. Perform quicksort on subarrays
//
// In the worst case (pivot not well-selected/array sorted in reverse order)
// the algorithm runs in O(n2) because the subtrees are uneven. One has n - 1
// elements and the other has 0. The algorithm has to go through n - 1 levels
// of partitioning.
//
// In the average case, the algorithm has a time complexity O(nlogn). Because
// the pivot is well selected, the subtrees are even and process n/2 elements each.
// This creates a balanced tree where instead of going through n - 1 levels of partitioning
// the task is well divided among logn (base 2) levels of execution in the binary tree.

func partition(arr []int, low int, high int) int {
	pivot := arr[high]
	leftwall := low

	i := leftwall
	for ; i < high-1; i++ {
		if arr[i] < pivot {
			// swap lefwall and i
			arr[leftwall], arr[i] = arr[i], arr[leftwall]

			// increment leftwall
			leftwall++
		}
	}

	arr[high], arr[leftwall] = arr[leftwall], arr[high]

	return leftwall
}

// Sort an array using the quicksort algorithm
func Sort(arr []int, low int, high int) []int {
	// Base case:
	// If low == high, then they point to the same value
	// If low > high, there is nothing to sort
	if low >= high {
		return arr
	}

	pivot := partition(arr, low, high)
	Sort(arr, low, pivot-1)
	Sort(arr, pivot+1, high)

	return arr
}
