package main

func countGoodStrings(low int, high int, zero int, one int) int {
	dp := make([]int, high+1)
	dp[0] = 1
	mod := 1_000_000_007

	for end := 1; end <= high; end++ {
		if end >= zero {
			dp[end] += dp[end-zero]
		}
		if end >= one {
			dp[end] += dp[end-one]
		}

		dp[end] %= mod
	}

	sum := 0
	for _, count := range dp[low : high+1] {
		sum += count
		sum %= mod
	}

	return sum
}
