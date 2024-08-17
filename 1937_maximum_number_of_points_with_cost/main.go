package main

import "math"

func maxPoints(points [][]int) int64 {
	rows, cols := len(points), len(points[0])

	// initialize prevRows
	prevRows := make([]int64, cols)
	for col := 0; col < cols; col++ {
		prevRows[col] = int64(points[0][col])
	}

	for row := 1; row < rows; row++ {
		left := make([]int64, cols)
		for col := 0; col < cols; col++ {
			left[col] = prevRows[col]
			if col > 0 {
				left[col] = max(left[col], left[col-1]-1)
			}
		}
		right := make([]int64, cols)
		for col := cols - 1; col >= 0; col-- {
			right[col] = prevRows[col]
			if col < cols-1 {
				right[col] = max(right[col], right[col+1]-1)
			}
		}
		currentRows := make([]int64, cols)
		for col := 0; col < cols; col++ {
			currentRows[col] = int64(points[row][col]) + max(left[col], right[col])
		}
		prevRows = currentRows
	}

	maxPoint := int64(math.MinInt64)
	for _, point := range prevRows {
		maxPoint = max(maxPoint, point)
	}

	return maxPoint
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a int64, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
