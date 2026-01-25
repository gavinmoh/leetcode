package main

import (
	"math"
	"sort"
)

func minimumDifference(nums []int, k int) int {
	sort.Ints(nums)
	minDiff := math.MaxInt

	for i := 0; i+k <= len(nums); i++ {
		diff := nums[i+k-1] - nums[i]
		if diff < minDiff {
			minDiff = diff
		}
	}

	return minDiff
}
