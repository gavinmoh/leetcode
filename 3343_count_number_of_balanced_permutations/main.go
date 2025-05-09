package main

const MOD = 1e9 + 7

func countBalancedPermutations(num string) int {
	n := len(num)
	freq := make([]int, 10)
	sum := 0
	for _, c := range num {
		x := int(c - '0')
		sum += x
		freq[x]++
	}

	if sum%2 != 0 {
		return 0
	}

	target := sum / 2
	maxOdd := (n + 1) / 2

	comb := make([][]int, maxOdd+1)
	for i := 0; i < len(comb); i++ {
		comb[i] = make([]int, maxOdd+1)
		comb[i][i], comb[i][0] = 1, 1
		for j := 1; j < i; j++ {
			comb[i][j] = (comb[i-1][j] + comb[i-1][j-1]) % MOD
		}
	}

	pSum := make([]int, 11)
	for i := 9; i >= 0; i-- {
		pSum[i] = pSum[i+1] + freq[i]
	}

	cache := make([][][]int, 10)
	for i := 0; i < len(cache); i++ {
		cache[i] = make([][]int, target+1)
		for j := 0; j < len(cache[i]); j++ {
			cache[i][j] = make([]int, maxOdd+1)
			for k := 0; k < len(cache[i][j]); k++ {
				cache[i][j][k] = -1
			}
		}
	}

	var dfs func(pos, current, oddCount int) int
	dfs = func(pos, current, oddCount int) int {
		if oddCount < 0 || pSum[pos] < oddCount || current > target {
			return 0
		}

		if pos > 9 {
			if current == target && oddCount == 0 {
				return 1
			}
			return 0
		}

		if cachedResult := cache[pos][current][oddCount]; cachedResult != -1 {
			return cachedResult
		}

		evenCount := pSum[pos] - oddCount
		result := 0
		start := max(0, freq[pos]-evenCount)
		end := min(freq[pos], oddCount)
		for i := start; i <= end; i++ {
			ways := comb[oddCount][i] * comb[evenCount][freq[pos]-i] % MOD
			result = (result + ways*dfs(pos+1, current+i*pos, oddCount-i)) % MOD
		}
		cache[pos][current][oddCount] = result
		return result
	}

	return dfs(0, 0, maxOdd)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
