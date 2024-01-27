package main

func kInversePairs(n int, k int) int {
	const modulo = 1e9 + 7
	dp := make([]int, k+1)
	dp[0] = 1

	for i := 1; i < n+1; i++ {
		current := make([]int, k+1)
		total := 0 // sliding window
		for j := 0; j < k+1; j++ {
			if j >= i {
				total = (total - dp[j-i] + modulo) % modulo
			}
			total = (total + dp[j]) % modulo
			current[j] = total
		}
		dp = current
	}

	return dp[k]
}
