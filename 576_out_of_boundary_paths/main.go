package main

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	modulo := 1_000_000_007
	cache := make(map[int]map[int]map[int]int)

	var dfs func(moves int, row int, col int) int
	dfs = func(moves int, row int, col int) int {
		// if we run out of move, then return 0
		if moves < 0 {
			return 0
		}
		// if we are out of bound, return 1
		if row > m-1 || col > n-1 || row < 0 || col < 0 {
			return 1
		}

		if _, ok := cache[row]; !ok {
			cache[row] = make(map[int]map[int]int)
		}

		if _, ok := cache[row][col]; !ok {
			cache[row][col] = make(map[int]int)
		}

		if cachedResult, ok := cache[row][col][moves]; ok {
			return cachedResult
		}

		top := dfs(moves-1, row+1, col)
		bottom := dfs(moves-1, row-1, col)
		left := dfs(moves-1, row, col-1)
		right := dfs(moves-1, row, col+1)
		result := (top + bottom + left + right) % modulo

		cache[row][col][moves] = result
		return result
	}

	return dfs(maxMove, startRow, startColumn)
}
