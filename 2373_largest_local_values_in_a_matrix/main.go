package main

import "math"

func largestLocal(grid [][]int) [][]int {
	rows, cols := len(grid), len(grid[0])
	result := [][]int{}

	for row := 0; row < rows-2; row++ {
		resultRow := []int{}
		for col := 0; col < cols-2; col++ {
			currentMax := math.MinInt
			for i := row; i < row+3; i++ {
				for j := col; j < col+3; j++ {
					if grid[i][j] > currentMax {
						currentMax = grid[i][j]
					}
				}
			}
			resultRow = append(resultRow, currentMax)
		}
		result = append(result, resultRow)
	}

	return result
}
