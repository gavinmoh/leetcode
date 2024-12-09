package main

import "sort"

func maximumBeauty(nums []int, k int) int {
	sort.Ints(nums)

	maxBeauty := 0
	left := 0
	for right := 0; right < len(nums); right++ {
		for nums[right]-nums[left] > (2 * k) {
			left++
		}
		beauty := right - left + 1
		if beauty > maxBeauty {
			maxBeauty = beauty
		}
	}

	return maxBeauty
}
