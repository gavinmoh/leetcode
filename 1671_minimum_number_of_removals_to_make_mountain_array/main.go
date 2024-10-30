package main

import "math"

func minimumMountainRemovals(nums []int) int {
	n := len(nums)
	lis := make([]int, n)
	lds := make([]int, n)
	for i := 0; i < n; i++ {
		lis[i] = 1
		lds[i] = 1
	}

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				lis[i] = max(lis[i], lis[j]+1)
			}
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if nums[i] > nums[j] {
				lds[i] = max(lds[i], lds[j]+1)
			}
		}
	}

	minRemovals := math.MaxInt
	for i := 0; i < n; i++ {
		if lis[i] > 1 && lds[i] > 1 {
			minRemovals = min(minRemovals, n-lis[i]-lds[i]+1)
		}
	}

	return minRemovals
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
