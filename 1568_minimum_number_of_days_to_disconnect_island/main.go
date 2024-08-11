package main

func minDays(grid [][]int) int {
	islands := numIslands(grid)

	if islands == 0 || islands > 1 {
		return 0
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				continue
			}

			grid[i][j] = 0
			count := numIslands(grid)
			if count == 0 || count > 1 {
				return 1
			}
			grid[i][j] = 1
		}
	}

	return 2
}

func numIslands(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])

	newGrid := make([][]int, rows)
	for i := 0; i < rows; i++ {
		newGrid[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			newGrid[i][j] = grid[i][j]
		}
	}

	var dfs func(row, col int)
	dfs = func(row, col int) {
		if newGrid[row][col] == 0 {
			return
		}

		// mark as visited
		newGrid[row][col] = 0

		// check top
		if row > 0 {
			dfs(row-1, col)
		}
		// check bottom
		if row+1 < rows {
			dfs(row+1, col)
		}
		// check left
		if col > 0 {
			dfs(row, col-1)
		}
		// check right
		if col+1 < cols {
			dfs(row, col+1)
		}
	}

	islands := 0
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if newGrid[row][col] == 1 {
				islands++
				dfs(row, col)
			}
		}
	}
	return islands
}
