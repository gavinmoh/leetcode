package main

func countSquares(matrix [][]int) int {
	rows, cols := len(matrix), len(matrix[0])
	cache := map[int]map[int]int{}

	var dfs func(row, col int) int
	dfs = func(row, col int) int {
		if row == rows || col == cols || matrix[row][col] == 0 {
			return 0
		}
		if _, ok := cache[row]; !ok {
			cache[row] = map[int]int{}
		}
		if cachedResult, ok := cache[row][col]; ok {
			return cachedResult
		}

		length := 1 + min(
			dfs(row+1, col),
			min(
				dfs(row, col+1),
				dfs(row+1, col+1),
			),
		)
		cache[row][col] = length

		return length
	}

	result := 0
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			result += dfs(row, col)
		}
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
