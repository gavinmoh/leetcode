package main

func cherryPickup(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	cache := make(map[int]map[int]map[int]int)

	var dfs func(int, int, int) int
	dfs = func(i int, j1 int, j2 int) int {
		count := 0
		count += grid[i][j1]
		if j1 != j2 {
			count += grid[i][j2]
		}

		// when reaching the bottom
		if i == rows-1 {
			return count
		}

		if _, ok := cache[i]; !ok {
			cache[i] = make(map[int]map[int]int)
		}

		if _, ok := cache[i][j1]; !ok {
			cache[i][j1] = make(map[int]int)
		}

		if cachedResult, ok := cache[i][j1][j2]; ok {
			return cachedResult
		}

		next1, next2 := []int{j1}, []int{j2}

		if j1-1 >= 0 {
			next1 = append(next1, j1-1)
		}

		if j1+1 < cols {
			next1 = append(next1, j1+1)
		}

		if j2-1 >= 0 {
			next2 = append(next2, j2-1)
		}

		if j2+1 < cols {
			next2 = append(next2, j2+1)
		}

		maxCount := 0
		for _, a := range next1 {
			for _, b := range next2 {
				maxCount = max(maxCount, dfs(i+1, a, b))
			}
		}

		result := count + maxCount
		cache[i][j1][j2] = result
		return result
	}

	return dfs(0, 0, cols-1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
