package main

func regionsBySlashes(grid []string) int {
	n := len(grid)
	matrix := make([][]int, n*3)
	for i := 0; i < n*3; i++ {
		matrix[i] = make([]int, n*3)
	}

	for i, row := range grid {
		for j, char := range row {
			switch char {
			case ' ':
				matrix[i*3][j*3] = 1
				matrix[i*3][j*3+1] = 1
				matrix[i*3][j*3+2] = 1
				matrix[i*3+1][j*3] = 1
				matrix[i*3+1][j*3+1] = 1
				matrix[i*3+1][j*3+2] = 1
				matrix[i*3+2][j*3] = 1
				matrix[i*3+2][j*3+1] = 1
				matrix[i*3+2][j*3+2] = 1
			case '/':
				matrix[i*3][j*3] = 1
				matrix[i*3][j*3+1] = 1
				matrix[i*3][j*3+2] = 0
				matrix[i*3+1][j*3] = 1
				matrix[i*3+1][j*3+1] = 0
				matrix[i*3+1][j*3+2] = 1
				matrix[i*3+2][j*3] = 0
				matrix[i*3+2][j*3+1] = 1
				matrix[i*3+2][j*3+2] = 1
			case '\\':
				matrix[i*3][j*3] = 0
				matrix[i*3][j*3+1] = 1
				matrix[i*3][j*3+2] = 1
				matrix[i*3+1][j*3] = 1
				matrix[i*3+1][j*3+1] = 0
				matrix[i*3+1][j*3+2] = 1
				matrix[i*3+2][j*3] = 1
				matrix[i*3+2][j*3+1] = 1
				matrix[i*3+2][j*3+2] = 0
			}
		}
	}

	return numIslands(matrix)
}

func numIslands(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	var dfs func(row, col int)
	dfs = func(row, col int) {
		if grid[row][col] == 0 {
			return
		}

		// mark as visited
		grid[row][col] = 0

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
			if grid[row][col] == 1 {
				islands++
				dfs(row, col)
			}
		}
	}
	return islands
}
