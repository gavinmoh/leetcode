package main

import "math"

func stoneGameII(piles []int) int {
	cache := make(map[bool]map[int]map[int]int)

	var dfs func(alice bool, i int, m int) int
	dfs = func(alice bool, i int, m int) int {
		if i == len(piles) {
			return 0
		}

		if _, ok := cache[alice]; !ok {
			cache[alice] = make(map[int]map[int]int)
		}

		if _, ok := cache[alice][i]; !ok {
			cache[alice][i] = make(map[int]int)
		}

		if cachedResult, ok := cache[alice][i][m]; ok {
			return cachedResult
		}

		total := 0
		result := 0
		if !alice {
			result = math.MaxInt
		}

		for x := 1; x <= m*2; x++ {
			if i+x > len(piles) {
				break
			}
			total += piles[i+x-1]
			if alice {
				result = max(result, total+dfs(!alice, i+x, max(m, x)))
			} else {
				result = min(result, dfs(!alice, i+x, max(m, x)))
			}
		}
		cache[alice][i][m] = result
		return result
	}

	return dfs(true, 0, 1)
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
