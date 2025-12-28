package main

func countNegatives(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])

	count := 0
	for row := rows - 1; row >= 0; row-- {
		for col := cols - 1; col >= 0; col-- {
			if grid[row][col] >= 0 {
				break
			}

			count++
		}
	}

	return count
}
