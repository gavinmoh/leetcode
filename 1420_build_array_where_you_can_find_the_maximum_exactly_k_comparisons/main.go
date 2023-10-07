package main

func numOfArrays(n int, m int, k int) int {
	// n: size of array
	// m: max value of array
	// k: search cost/number of elements to search
	// m must be at the position of k
	// arr[:k+1] and arr[k:] must be less than m
	if k == 0 {
		return 0
	}

	mod := 1000000007

	// initialize dp
	// dp[i][j][k] = number of ways to build an array of size i, with j elements, and max value k
	dp := make([][][]int, n+1)
	for i := range dp {
		dp[i] = make([][]int, k+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, m+1)
		}
	}

	for i := 1; i <= m; i++ {
		dp[1][1][i] = 1
	}

	for i := 2; i <= n; i++ {
		for j := 1; j <= k; j++ {
			for l := 1; l <= m; l++ {
				dp[i][j][l] = (dp[i-1][j][l] * l) % mod
				for l0 := 1; l0 < l; l0++ {
					dp[i][j][l] = (dp[i][j][l] + dp[i-1][j-1][l0]) % mod
				}
			}
		}
	}

	result := 0
	for i := 1; i <= m; i++ {
		result = (result + dp[n][k][i]) % mod
	}

	return result
}
