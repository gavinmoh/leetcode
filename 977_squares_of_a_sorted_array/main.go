package main

import "sort"

func sortedSquares(nums []int) []int {
	for i, num := range nums {
		nums[i] = num * num
	}
	sort.Ints(nums)
	return nums
}
