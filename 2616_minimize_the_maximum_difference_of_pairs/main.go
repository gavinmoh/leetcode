package main

import "sort"

func minimizeMax(nums []int, p int) int {
	sort.Ints(nums)

	n := len(nums)
	left, right := 0, nums[n-1]-nums[0]

	for left < right {
		mid := left + (right-left)/2

		count := 0
		for i := 0; i < n-1; i++ {
			if nums[i+1]-nums[i] <= mid {
				count++
				i++
			}
		}

		if count >= p {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}
