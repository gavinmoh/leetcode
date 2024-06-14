package main

import "sort"

func minIncrementForUnique(nums []int) int {
	sort.Ints(nums)

	count := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] < nums[i]+1 {
			count += nums[i] - nums[i+1] + 1
			nums[i+1] = nums[i] + 1
		}
	}

	return count
}
