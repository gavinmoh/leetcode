package main

import "math"

func maximumCandies(candies []int, k int64) int {
	left, right := 0, math.MinInt
	for _, pile := range candies {
		right = max(right, pile)
	}

	for left < right {
		mid := (left + right + 1) / 2
		remaining := k
		for _, pile := range candies {
			remaining -= int64(pile / mid)

			if remaining <= 0 {
				break
			}
		}

		if remaining > 0 {
			right = mid - 1
		} else {
			left = mid
		}
	}

	return left
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
