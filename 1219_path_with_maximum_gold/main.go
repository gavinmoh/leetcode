package main

func getMaximumGold(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])

	var dfs func(row int, col int) int
	dfs = func(row int, col int) int {
		gold := grid[row][col]
		grid[row][col] = 0 // marked as visited
		currentMax := 0

		// move up
		if row > 0 && grid[row-1][col] != 0 {
			currentMax = max(currentMax, dfs(row-1, col))
		}

		// move down
		if row < rows-1 && grid[row+1][col] != 0 {
			currentMax = max(currentMax, dfs(row+1, col))
		}

		// move left
		if col > 0 && grid[row][col-1] != 0 {
			currentMax = max(currentMax, dfs(row, col-1))
		}

		// move right
		if col < cols-1 && grid[row][col+1] != 0 {
			currentMax = max(currentMax, dfs(row, col+1))
		}
		result := gold + currentMax
		grid[row][col] = gold

		return result
	}

	maxGold := 0
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] != 0 {
				maxGold = max(maxGold, dfs(row, col))
			}
		}
	}

	return maxGold
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
