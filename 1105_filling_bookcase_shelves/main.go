package main

import "math"

func minHeightShelves(books [][]int, shelfWidth int) int {
	n := len(books)
	cache := make(map[int]int)
	var dp func(i int) int
	dp = func(i int) int {
		if i == n {
			return 0
		}

		if cachedResult, ok := cache[i]; ok {
			return cachedResult
		}

		currentWidth := shelfWidth
		maxHeight := 0
		result := math.MaxInt
		for j := i; j < n; j++ {
			width, height := books[j][0], books[j][1]
			if currentWidth < width {
				break
			}

			currentWidth -= width
			maxHeight = max(maxHeight, height)
			result = min(result, dp(j+1)+maxHeight)
		}
		cache[i] = result

		return result
	}

	return dp(0)
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
