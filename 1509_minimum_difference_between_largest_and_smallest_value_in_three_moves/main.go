package main

import "sort"

func minDifference(nums []int) int {
	n := len(nums)
	if n < 5 {
		return 0
	}

	sort.Ints(nums)

	// delete 3 from start
	result := nums[n-1] - nums[3]
	// delete 2 from start and 1 from end
	result = min(result, nums[n-2]-nums[2])
	// delete 1 from start and 2 from end
	result = min(result, nums[n-3]-nums[1])
	// delete 3 from end
	result = min(result, nums[n-4]-nums[0])

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
