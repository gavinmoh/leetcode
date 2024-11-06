package main

import "sort"

func countFairPairs(nums []int, lower int, upper int) int64 {
	sort.Ints(nums)

	lowerBound := func(target int) int64 {
		left, right := 0, len(nums)-1
		count := int64(0)

		for left < right {
			sum := nums[left] + nums[right]
			if sum < target {
				count += int64(right - left)
				left++
			} else {
				right--
			}
		}

		return count
	}

	return lowerBound(upper+1) - lowerBound(lower)
}
