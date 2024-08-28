package main

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	rows, cols := len(grid1), len(grid1[0])
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid2[row][col] > 0 {
				grid2[row][col] += grid1[row][col]
			}
		}
	}

	islands := [][][]int{{}}

	var dfs func(int, int)
	dfs = func(row int, col int) {
		if grid2[row][col] == 0 {
			return
		}

		islands[len(islands)-1] = append(islands[len(islands)-1], []int{row, col, grid2[row][col]})
		grid2[row][col] = 0

		if row < rows-1 {
			dfs(row+1, col)
		}

		if row > 0 {
			dfs(row-1, col)
		}

		if col < cols-1 {
			dfs(row, col+1)
		}

		if col > 0 {
			dfs(row, col-1)
		}
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid2[row][col] > 0 {
				dfs(row, col)
				islands = append(islands, [][]int{})
			}
		}
	}

	count := 0
	for _, island := range islands {
		isSubIsland := true
		if len(island) < 1 {
			continue
		}

		for _, cell := range island {
			if cell[2] == 1 {
				isSubIsland = false
				break
			}
		}
		if isSubIsland {
			count++
		}
	}

	return count
}
