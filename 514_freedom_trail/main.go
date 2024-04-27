package main

func findRotateSteps(ring string, key string) int {
	m, n := len(ring), len(key)
	cache := make(map[int]map[int]map[int]int)

	// two pointers i & j represents ring[i] and key[j] respectively
	// when clockwise: step == 1
	// when counter-clockwise: step == -1
	var dfs func(i int, j int, step int) int
	dfs = func(i int, j int, step int) int {
		// out of bound
		if j >= n {
			return 0
		}

		if _, ok := cache[i]; !ok {
			cache[i] = make(map[int]map[int]int)
		}

		if _, ok := cache[i][j]; !ok {
			cache[i][j] = make(map[int]int)
		}

		if cachedResult, ok := cache[i][j][step]; ok {
			return cachedResult
		}

		current, count := i, 1
		for ring[current] != key[j] {
			current += step
			count++

			// reset current
			if current < 0 {
				current = m - 1
			} else if current >= m {
				current = 0
			}
		}

		left := dfs(current, j+1, 1) + count
		right := dfs(current, j+1, -1) + count
		result := min(left, right)
		cache[i][j][step] = result

		return result
	}

	return min(dfs(0, 0, 1), dfs(0, 0, -1))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
