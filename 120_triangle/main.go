package main

import (
	"math"
	"slices"
)

func minimumTotal(triangle [][]int) int {
	for row := range triangle {
		if row == 0 {
			continue
		}

		for col := range triangle[row] {
			left, right := math.MaxInt, math.MaxInt

			if col != len(triangle[row])-1 {
				right = triangle[row][col] + triangle[row-1][col]
			}

			if col != 0 {
				left = triangle[row][col] + triangle[row-1][col-1]
			}

			triangle[row][col] = min(left, right)
		}
	}

	lastRow := triangle[len(triangle)-1]
	return slices.Min(lastRow)
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
