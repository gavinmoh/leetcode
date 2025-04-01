package main

func mostPoints(questions [][]int) int64 {
	n := len(questions)
	cache := map[int]int64{}

	var dfs func(i int) int64
	dfs = func(i int) int64 {
		if i >= n {
			return int64(0)
		}

		if cachedResult, ok := cache[i]; ok {
			return cachedResult
		}

		question := questions[i]
		points, brainPower := int64(question[0]), question[1]

		left := points + dfs(i+brainPower+1)
		right := dfs(i + 1)

		cache[i] = max(left, right)
		return cache[i]
	}

	return dfs(0)
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
