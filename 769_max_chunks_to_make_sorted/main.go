package main

import "math"

func maxChunksToSorted(arr []int) int {
	count := 0
	currentMax := math.MinInt
	for i, num := range arr {
		if num > currentMax {
			currentMax = num
		}
		if i >= currentMax {
			count++
			currentMax = math.MinInt
		}
	}

	return count
}
