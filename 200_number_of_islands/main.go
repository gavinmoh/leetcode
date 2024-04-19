package main

func numIslands(grid [][]byte) int {
	rows, cols := len(grid), len(grid[0])
	var dfs func(row, col int)
	dfs = func(row, col int) {
		if grid[row][col] == '0' {
			return
		}

		// mark as visited
		grid[row][col] = '0'

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
			if grid[row][col] == '1' {
				islands++
				dfs(row, col)
			}
		}
	}
	return islands
}
