package main

import "sort"

func largestPerimeter(nums []int) int {
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	for i := 0; i < len(nums)-2; i++ {
		if nums[i+1]+nums[i+2] > nums[i] {
			return nums[i] + nums[i+1] + nums[i+2]
		}
	}

	return 0
}
