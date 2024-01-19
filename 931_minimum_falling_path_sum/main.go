package main

func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	cache := make(map[int]map[int]int)

	var dfs func(int, int) int
	dfs = func(row int, col int) int {
		if row >= n {
			return 0
		}

		if _, ok := cache[row]; !ok {
			cache[row] = make(map[int]int)
		}

		if _, ok := cache[row][col]; ok {
			return cache[row][col]
		}

		element := matrix[row][col]
		bottom := element + dfs(row+1, col)

		if col-1 < 0 {
			right := element + dfs(row+1, col+1)
			cache[row][col] = min(right, bottom)
		} else if col+1 >= n {
			left := element + dfs(row+1, col-1)
			cache[row][col] = min(left, bottom)
		} else {
			right := element + dfs(row+1, col+1)
			left := element + dfs(row+1, col-1)
			cache[row][col] = min(left, min(right, bottom))
		}

		return cache[row][col]
	}

	var minSum int
	for col := 0; col < n; col++ {
		currentMin := dfs(0, col)

		if col == 0 {
			minSum = currentMin
		} else {
			minSum = min(minSum, currentMin)
		}
	}

	return minSum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
