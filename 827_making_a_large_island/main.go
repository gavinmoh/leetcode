package main

func largestIsland(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	islandSize := map[int]int{}
	sea := [][]int{}

	var dfs func(row, col, islandIdx int) int
	dfs = func(row, col, islandIdx int) int {
		if grid[row][col] == 0 || grid[row][col] > 1 {
			return 0
		}

		grid[row][col] = islandIdx
		size := 1

		// top
		if row > 0 {
			size += dfs(row-1, col, islandIdx)
		}

		// bottom
		if row < rows-1 {
			size += dfs(row+1, col, islandIdx)
		}

		// left
		if col > 0 {
			size += dfs(row, col-1, islandIdx)
		}

		// right
		if col < cols-1 {
			size += dfs(row, col+1, islandIdx)
		}

		return size
	}

	islandIdx := 2
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == 0 {
				sea = append(sea, []int{row, col})
			} else if grid[row][col] > 1 {
				continue
			} else {
				islandSize[islandIdx] = dfs(row, col, islandIdx)
				islandIdx++
			}
		}
	}

	if len(sea) == 0 {
		return rows * cols
	}

	if len(islandSize) == 0 {
		return 1
	}

	maxSize := 0
	for _, size := range islandSize {
		maxSize = max(maxSize, size)
	}
	maxSize += 1

	for _, cell := range sea {
		row, col := cell[0], cell[1]
		neighbourIslands := map[int]bool{}

		// top
		if row > 0 && grid[row-1][col] > 1 {
			neighbourIslands[grid[row-1][col]] = true
		}

		// bottom
		if row < rows-1 && grid[row+1][col] > 1 {
			neighbourIslands[grid[row+1][col]] = true
		}

		// left
		if col > 0 && grid[row][col-1] > 1 {
			neighbourIslands[grid[row][col-1]] = true
		}

		// right
		if col < cols-1 && grid[row][col+1] > 1 {
			neighbourIslands[grid[row][col+1]] = true
		}

		size := 1
		for island, _ := range neighbourIslands {
			size += islandSize[island]
		}

		maxSize = max(maxSize, size)
	}

	return maxSize
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
