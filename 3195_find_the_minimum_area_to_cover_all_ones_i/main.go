package main

func minimumArea(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	minRow, maxRow := rows, 0
	minCol, maxCol := cols, 0

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == 1 {
				minRow = min(minRow, row)
				maxRow = max(maxRow, row)
				minCol = min(minCol, col)
				maxCol = max(maxCol, col)
			}
		}
	}

	height := maxRow - minRow + 1
	width := maxCol - minCol + 1

	return height * width
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
