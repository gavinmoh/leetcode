package main

import "sort"

func maxWidthRamp(nums []int) int {
	indices := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		indices[i] = []int{i, nums[i]}
	}

	sort.SliceStable(indices, func(i, j int) bool {
		return indices[i][1] < indices[j][1]
	})

	minIndex := len(nums)
	maxWidth := 0

	for _, item := range indices {
		index := item[0]
		maxWidth = max(maxWidth, index-minIndex)
		minIndex = min(minIndex, index)
	}

	return maxWidth
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
