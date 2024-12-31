package main

func mincostTickets(days []int, costs []int) int {
	cache := map[int]map[int]int{}

	var dp func(i int, validUntil int) int
	dp = func(i int, validUntil int) int {
		if i >= len(days) {
			return 0
		}

		day := days[i]
		if validUntil >= day {
			return dp(i+1, validUntil)
		}

		if _, ok := cache[i]; !ok {
			cache[i] = map[int]int{}
		}

		if cachedResult, ok := cache[i][validUntil]; ok {
			return cachedResult
		}

		one := costs[0] + dp(i+1, day)
		seven := costs[1] + dp(i+1, day+6)
		thirty := costs[2] + dp(i+1, day+29)

		cache[i][validUntil] = min(one, min(seven, thirty))

		return cache[i][validUntil]
	}

	result := dp(0, 0)

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
