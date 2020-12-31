package mergesort

// Mergesort
//
// 1. Divide the list in two equal lengths
// 2. Recursively divide the sublists into equal lengths
// 3. Merge the sublists, placing items in order
//
// The time complexity in the worst and average case: O(nlogn)
// The problem is solved recursively by splitting the input
// in half at each level. Once the base case is reached (a
// list of size 1) then the results are merged in order.

func merge(left, right []int) []int {
	combined := make([]int, 0)
	for (len(left) != 0) && (len(right) != 0) {
		firstLeft := left[0]
		firstRight := right[0]

		if firstLeft < firstRight {
			combined = append(combined, firstLeft)
			left = left[1:]
		} else {
			combined = append(combined, firstRight)
			right = right[1:]
		}
	}

	if len(left) != 0 {
		combined = append(combined, left...)
	}

	if len(right) != 0 {
		combined = append(combined, right...)
	}

	return combined
}

// Sort an array using the mergesort algorithm
func Sort(arr []int) []int {
	// Base case
	if len(arr) <= 1 {
		return arr
	}

	// Recursive case
	// Divide the list into equal-sized sublists
	left := make([]int, 0)
	right := make([]int, 0)

	for i, v := range arr {
		if i < len(arr)/2 {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	left = Sort(left)
	right = Sort(right)

	return merge(left, right)
}
