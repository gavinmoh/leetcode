package main

import "math"

func totalFruit(fruits []int) int {
	freq := map[int]int{}
	result, currentLength := math.MinInt, 0

	left := 0
	for right := 0; right < len(fruits); right++ {
		freq[fruits[right]]++
		currentLength++

		for len(freq) > 2 {
			freq[fruits[left]]--
			currentLength--
			if freq[fruits[left]] == 0 {
				delete(freq, fruits[left])
			}
			left++
		}

		result = max(result, currentLength)
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
