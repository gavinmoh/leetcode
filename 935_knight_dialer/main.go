package main

func knightDialer(n int) int {
	jumps := [][]int{
		{4, 6},
		{6, 8},
		{7, 9},
		{4, 8},
		{3, 9, 0},
		{},
		{1, 7, 0},
		{2, 6},
		{1, 3},
		{2, 4},
	}
	cache := make(map[int]map[int]int)
	modulo := 1_000_000_000 + 7

	var dp func(remain, square int) (ans int)
	dp = func(remain, square int) (ans int) {
		if remain == 0 {
			return 1
		}
		if _, ok := cache[remain][square]; ok {
			return cache[remain][square]
		}
		for _, nextSquare := range jumps[square] {
			ans = (ans + dp(remain-1, nextSquare)) % modulo
		}
		if _, ok := cache[remain]; ok {
			cache[remain][square] = ans
		} else {
			cacheItem := make(map[int]int)
			cacheItem[square] = ans
			cache[remain] = cacheItem
		}
		return ans
	}

	ans := 0
	for square := 0; square < 10; square++ {
		ans = (ans + dp(n-1, square)) % modulo
	}

	return ans
}
