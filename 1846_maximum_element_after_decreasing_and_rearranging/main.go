package main

import "sort"

func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	n := len(arr)

	// sort the array
	sort.Ints(arr)

	// make sure first element is 1
	arr[0] = 1

	// loop and decrease each element to statisfy
	// abs(arr[n] - arr[n-1]) <= 1
	for i := 1; i < n; i++ {
		diff := arr[i] - arr[i-1]

		if diff > 1 {
			arr[i] = arr[i-1] + 1
		}
	}

	return arr[n-1]
}
