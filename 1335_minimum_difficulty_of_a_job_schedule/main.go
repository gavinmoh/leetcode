package main

import "math"

type Key struct {
	i          int
	d          int
	currentMax int
}

func minDifficulty(jobDifficulty []int, d int) int {
	n := len(jobDifficulty)
	if n < d {
		return -1
	}

	cache := make(map[Key]int)

	var dfs func(i int, d int, currentMax int) int
	dfs = func(i int, d int, currentMax int) int {
		if i == n {
			if d == 0 {
				return 0
			} else {
				return math.MaxInt32
			}
		}

		if d == 0 {
			return math.MaxInt32
		}

		key := Key{i: i, d: d, currentMax: currentMax}
		if cachedResult, ok := cache[key]; ok {
			return cachedResult
		}

		currentMax = max(currentMax, jobDifficulty[i])

		result := min(dfs(i+1, d, currentMax), dfs(i+1, d-1, -1)+currentMax)
		cache[key] = result
		return result
	}

	return dfs(0, d, -1)
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
