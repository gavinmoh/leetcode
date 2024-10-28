package main

func maxMoves(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	cache := map[int]map[int]int{}

	var dfs func(row, col int) int
	dfs = func(row, col int) int {
		if _, ok := cache[row]; !ok {
			cache[row] = map[int]int{}
		}

		if cachedResult, ok := cache[row][col]; ok {
			return cachedResult
		}

		move := 0
		if col+1 != cols {
			if row > 0 && grid[row-1][col+1] > grid[row][col] {
				move = max(move, dfs(row-1, col+1)+1)
			}

			if grid[row][col+1] > grid[row][col] {
				move = max(move, dfs(row, col+1)+1)
			}

			if row < rows-1 && grid[row+1][col+1] > grid[row][col] {
				move = max(move, dfs(row+1, col+1)+1)
			}
		}

		cache[row][col] = move
		return move
	}

	maxMove := -1
	for row := 0; row < rows; row++ {
		maxMove = max(maxMove, dfs(row, 0))
	}

	return maxMove
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
