package main

import "math"

func luckyNumbers(matrix [][]int) []int {
	rows, cols := len(matrix), len(matrix[0])
	rowsMin := make(map[int]int)
	for i := 0; i < rows; i++ {
		rowsMin[i] = math.MaxInt
	}
	colsMax := make(map[int]int)
	for i := 0; i < cols; i++ {
		colsMax[i] = math.MinInt
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			num := matrix[row][col]
			if num < rowsMin[row] {
				rowsMin[row] = num
			}
			if num > colsMax[col] {
				colsMax[col] = num
			}
		}
	}

	result := []int{}
	nums := make(map[int]bool)
	for _, num := range rowsMin {
		nums[num] = true
	}

	for _, num := range colsMax {
		if _, ok := nums[num]; ok {
			result = append(result, num)
		}
	}

	return result
}
