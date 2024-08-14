package main

import (
	"sort"
)

func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)

	helper := func(distance int) int {
		left := 0
		result := 0
		for right := 0; right < len(nums); right++ {
			for nums[right]-nums[left] > distance {
				left++
			}
			result += right - left
		}
		return result
	}

	left, right := 0, nums[len(nums)-1]
	for left < right {
		mid := (left + right) / 2
		pairs := helper(mid)
		if pairs >= k {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return right
}
