package main

import "slices"

func minimumSize(nums []int, maxOperations int) int {
	// assume penalty is between 1 to max of number
	// check if nums can be distributed into smaller than penalty
	// use binary search to find the min penalty
	left, right := 1, slices.Max(nums)

	helper := func(maxBalls int) bool {
		totalOperations := 0
		for _, num := range nums {
			totalOperations += (1 + (num-1)/maxBalls) - 1 // ceil division

			if totalOperations > maxOperations {
				return false
			}
		}

		return true
	}

	for left < right {
		mid := (left + right) / 2

		if helper(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}
