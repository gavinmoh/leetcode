package main

import (
	"math"
)

func minFallingPathSum(grid [][]int) int {
	n := len(grid)
	cache := make(map[int]map[int]int)

	var dfs func(row int, prevCol int) int
	dfs = func(row int, prevCol int) int {
		if row >= n {
			return 0
		}

		if _, ok := cache[row]; !ok {
			cache[row] = make(map[int]int)
		}

		if cachedResult, ok := cache[row][prevCol]; ok {
			return cachedResult
		}

		result := math.MaxInt
		for i, num := range grid[row] {
			if i == prevCol {
				continue
			}
			result = min(result, dfs(row+1, i)+num)
		}

		cache[row][prevCol] = result
		return result
	}

	minSum := dfs(0, -1)
	for i := 1; i < n; i++ {
		minSum = min(minSum, dfs(0, i))
	}

	return minSum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
