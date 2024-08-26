package sortimpl

// selection sort implementation
// sort the array and return new one
func Selection(array []int) []int {
	// 1. Iterate through the array: Start from the beginning of the array and compare the current element with all the remaining elements
	// 2. Find the minimum: Find the index of the minimum element in the unsorted part of the array
	// 3. Swap: Swap the current element with the minimum element
	// 4. Repeat: Repeat steps 1-3 until the entire array is sorted
	for i := 0; i < len(array); i++ {
		for j := i; j < len(array); j++ {
			// move the min data towards start index at each iteration
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}

	return array
}

// bubble sort implementation
// return new array
func Bubble(array []int) []int {
	// 1. Iterate through the array: Start from the beginning of the array and compare adjacent elements
	// 2. Swap if necessary: If the current element is greater than the next element, swap them
	// 3. Repeat: Continue iterating through the array until no swaps are made during a pass
	var swapped bool
	for i := 0; i < len(array); i++ {
		swapped = false
		for j := 1; j < len(array)-i; j++ {
			if array[j-1] > array[j] {
				swapped = true
				array[j-1], array[j] = array[j], array[j-1]
			}
		}
		// if not swap happend then the array get sorted
		// so break the loop, no need to check remaining iterations
		if !swapped {
			break
		}
	}

	return array
}

// insertion sort
func Insertion(array []int) []int {
	// 1. Start with the second element: Begin with the second element of the array (index 1)
	// 2. Compare with previous elements: Compare this element with the elements to its left (already sorted subarray)
	// 3. Shift elements: If the current element is smaller than the previous element, shift the previous element one position to the right. Repeat this process until the current element is no longer smaller than the previous element
	// 4. Insert: Insert the current element into the empty position
	// 5. Repeat: Move to the next element and repeat steps 2-4 until all elements have been sorted
	for i := 1; i < len(array); i++ {
		for j := i; j > 0; j-- {
			if array[j-1] > array[j] {
				array[j-1], array[j] = array[j], array[j-1]
			}
		}
	}

	return array
}

// help function for combine splitted array as sorted order
func mergeSort(left, right []int) []int {
	result := make([]int, len(left)+len(right))
	var i, j, k int
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			j++
		}
		k++
	}
	for i < len(left) {
		result[k] = left[i]
		i++
		k++
	}
	for j < len(right) {
		result[k] = right[j]
		j++
		k++
	}
	return result
}

// merge sort
func Merge(array []int) []int {
	// 1. Divide: Divide the unsorted array into two halves
	// 2. Conquer: Recursively sort each half using merge sort
	// 3. Combine: Merge the sorted halves into a single sorted array
	if len(array) <= 1 {
		return array
	}
	// split the array as two subarray
	mid := len(array) / 2
	left := Merge(array[:mid])
	right := Merge(array[mid:])

	// combine the subarray as sorted
	return mergeSort(left, right)
}
