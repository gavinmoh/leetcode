package main

import "bytes"

func maxEqualRowsAfterFlips(matrix [][]int) int {
	patternFreq := map[string]int{}

	for _, row := range matrix {
		var pattern bytes.Buffer
		for _, num := range row {
			if num == row[0] {
				pattern.WriteString("T")
			} else {
				pattern.WriteString("F")
			}
		}
		patternFreq[pattern.String()]++
	}

	maxFreq := 0
	for _, freq := range patternFreq {
		if freq > maxFreq {
			maxFreq = freq
		}
	}

	return maxFreq
}
