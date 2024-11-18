package main

import "math"

func maxMatrixSum(matrix [][]int) int64 {
	rows, cols := len(matrix), len(matrix[0])
	count, sum, smallest := 0, int64(0), math.MaxInt

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			num := matrix[row][col]
			if num < 0 {
				count++
				num = -num
			}
			sum += int64(num)
			smallest = min(smallest, num)
		}
	}

	if count%2 == 0 {
		return sum
	}

	return sum - int64(2*smallest)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
