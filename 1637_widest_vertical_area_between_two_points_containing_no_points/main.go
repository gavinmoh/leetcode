package main

import (
	"math"
	"sort"
)

func maxWidthOfVerticalArea(points [][]int) int {
	// sort by x
	sort.SliceStable(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	// find the max difference between two x
	maxDiff := -math.MaxInt

	for i := 1; i < len(points); i++ {
		diffX := abs(points[i-1][0] - points[i][0])
		maxDiff = max(maxDiff, diffX)
	}

	return maxDiff
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
