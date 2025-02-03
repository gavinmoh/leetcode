package main

import "math"

func minimumDeletions(word string, k int) int {
	freq := make([]int, 26)
	for _, char := range word {
		freq[char-'a']++
	}

	smallest := math.MaxInt
	for _, count := range freq {
		smallest = min(smallest, count)
	}

	minDeletions := math.MaxInt
	for _, x := range freq {
		deletions := 0
		for _, y := range freq {
			if y < x {
				deletions += y
			}
			if y > x+k {
				deletions += (y - x - k)
			}
		}
		minDeletions = min(minDeletions, deletions)
	}

	return minDeletions
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
