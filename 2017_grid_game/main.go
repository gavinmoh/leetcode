package main

import "math"

func gridGame(grid [][]int) int64 {
	firstRowSum, secondRowSum := int64(0), int64(0)
	for _, num := range grid[0] {
		firstRowSum += int64(num)
	}

	var minimumSum int64 = math.MaxInt64
	for turnIndex := 0; turnIndex < len(grid[0]); turnIndex++ {
		firstRowSum -= int64(grid[0][turnIndex])
		minimumSum = min(minimumSum, max(firstRowSum, secondRowSum))
		secondRowSum += int64(grid[1][turnIndex])
	}

	return minimumSum
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
