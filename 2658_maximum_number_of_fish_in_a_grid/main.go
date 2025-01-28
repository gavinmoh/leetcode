package main

func findMaxFish(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	maxFish := 0

	var dfs func(row, col int) int
	dfs = func(row, col int) int {
		if grid[row][col] == 0 {
			return 0
		}

		fish := grid[row][col]
		grid[row][col] = 0

		// top
		if row > 0 {
			fish += dfs(row-1, col)
		}

		// bottom
		if row < rows-1 {
			fish += dfs(row+1, col)
		}

		// left
		if col > 0 {
			fish += dfs(row, col-1)
		}

		// right
		if col < cols-1 {
			fish += dfs(row, col+1)
		}

		return fish
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			maxFish = max(maxFish, dfs(row, col))
		}
	}

	return maxFish
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
