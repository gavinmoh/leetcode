package main

import "math"

func minCapability(nums []int, k int) int {
	left, right := 1, math.MinInt
	for _, num := range nums {
		right = max(right, num)
	}

	for left < right {
		mid := (left + right) / 2
		possible := 0

		i := 0
		for i < len(nums) {
			if nums[i] <= mid {
				possible++
				i += 2
			} else {
				i++
			}
		}

		if possible >= k {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
