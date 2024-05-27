package main

import "sort"

func specialArray(nums []int) int {
	sort.Ints(nums)
	n := len(nums)

	for x := n; x >= 0; x-- {
		count := 0
		for i := n - 1; i >= 0; i-- {
			if nums[i] >= x {
				count++
			} else {
				break
			}
		}

		if count == x {
			return x
		}
	}

	return -1
}
