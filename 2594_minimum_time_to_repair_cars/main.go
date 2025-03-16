package main

import "math"

func repairCars(ranks []int, cars int) int64 {
	maxRank := -1
	for _, rank := range ranks {
		maxRank = max(maxRank, rank)
	}

	check := func(time int64) bool {
		repairs := 0
		for _, rank := range ranks {
			repairs += int(math.Sqrt(float64(time / int64(rank))))
			if repairs >= cars {
				return true
			}
		}
		return false
	}

	left, right := int64(1), int64(maxRank)*int64(cars)*int64(cars)
	for left <= right {
		mid := left + (right-left)/int64(2)
		if check(mid) {
			right = mid - 1
		} else {
			left = mid + 1
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
