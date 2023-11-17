package main

import "sort"

func minPairSum(nums []int) int {
	n := len(nums)
	sort.Ints(nums)

	output := -1
	j := n - 1
	for i := 0; i < n/2; i++ {
		left := nums[i]
		right := nums[j]

		output = max(output, left+right)

		j--
	}

	return output
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
