package main

func numRollsToTarget(n int, k int, target int) int {
	modulo := 1_000_000_007
	cache := make(map[int]map[int]int)

	var dfs func(n, target int) int
	dfs = func(n, target int) int {
		if n == 0 {
			if target == 0 {
				return 1
			} else {
				return 0
			}
		}

		if _, ok := cache[n]; !ok {
			cache[n] = make(map[int]int)
		}

		if _, ok := cache[n][target]; ok {
			return cache[n][target]
		}

		count := 0
		for i := 1; i <= k; i++ {
			count += (dfs(n-1, target-i) % modulo)
		}
		cache[n][target] = count
		return cache[n][target]
	}

	return dfs(n, target) % modulo

}
