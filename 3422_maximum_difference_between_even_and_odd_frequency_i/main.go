package main

import "math"

func maxDifference(s string) int {
	freq := make([]int, 26)
	for _, c := range s {
		freq[c-'a']++
	}

	maxOdd, minEven := math.MinInt, math.MaxInt
	for _, count := range freq {
		if count == 0 {
			continue
		}

		if count%2 == 0 {
			minEven = min(minEven, count)
		} else {
			maxOdd = max(maxOdd, count)
		}
	}

	return maxOdd - minEven
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
