package main

func countVowelPermutation(n int) int {

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, 5)
	}

	for i := 0; i < 5; i++ {
		dp[1][i] = 1
	}

	modulo := 1000000007

	for i := 2; i <= n; i++ {
		// a, e, i, o, u
		dp[i][0] = (dp[i-1][1] + dp[i-1][2] + dp[i-1][4]) % modulo
		dp[i][1] = (dp[i-1][0] + dp[i-1][2]) % modulo
		dp[i][2] = (dp[i-1][1] + dp[i-1][3]) % modulo
		dp[i][3] = dp[i-1][2] % modulo
		dp[i][4] = (dp[i-1][2] + dp[i-1][3]) % modulo
	}

	return (dp[n][0] + dp[n][1] + dp[n][2] + dp[n][3] + dp[n][4]) % modulo
}
