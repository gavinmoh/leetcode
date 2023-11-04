package main

import "math"

func getLastMoment(n int, left []int, right []int) int {
	if len(left) == 0 && len(right) == 0 {
		return 0
	} else if len(left) == 0 {
		return n - min(right...)
	} else if len(right) == 0 {
		return max(left...)
	}

	var maxLeft, minRight int
	if len(left) > 0 {
		maxLeft = max(left...)
	}
	if len(right) > 0 {
		minRight = min(right...)
	}
	return int(math.Max(float64(maxLeft), float64(n-minRight)))
}

func max(nums ...int) int {
	max := math.MinInt32
	for _, n := range nums {
		if n > max {
			max = n
		}
	}
	return max
}

func min(nums ...int) int {
	min := math.MaxInt32
	for _, n := range nums {
		if n < min {
			min = n
		}
	}
	return min
}
