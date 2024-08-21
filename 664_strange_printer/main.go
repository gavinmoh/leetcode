package main

func strangePrinter(s string) int {
	processed := []byte{s[0]}
	for i := 1; i < len(s); i++ {
		if s[i-1] == s[i] {
			continue
		}
		processed = append(processed, s[i])
	}
	s = string(processed)

	cache := make(map[int]map[int]int)

	var dfs func(int, int) int
	dfs = func(start int, end int) int {
		if start > end {
			return 0
		}

		if _, ok := cache[start]; !ok {
			cache[start] = make(map[int]int)
		}

		if cachedResult, ok := cache[start][end]; ok {
			return cachedResult
		}

		minTurns := 1 + dfs(start+1, end)

		for i := start + 1; i <= end; i++ {
			if s[i] == s[start] {
				turns := dfs(start, i-1) + dfs(i+1, end)
				if turns < minTurns {
					minTurns = turns
				}
			}
		}

		cache[start][end] = minTurns
		return minTurns
	}

	return dfs(0, len(s)-1)
}
