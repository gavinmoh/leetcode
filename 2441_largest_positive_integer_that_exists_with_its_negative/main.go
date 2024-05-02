package main

import "sort"

func findMaxK(nums []int) int {
	sort.Ints(nums)
	left, right := 0, len(nums)-1
	for left < right {
		if nums[right] == -nums[left] {
			return nums[right]
		}

		if nums[right] > -nums[left] {
			right--
		} else {
			left++
		}
	}

	return -1
}
