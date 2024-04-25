package main

func longestIdealString(s string, k int) int {
	dp := [26]int{}
	result := 0

	for _, char := range s {
		current := int(char - 'a')
		longest := 1
		for previous := 0; previous < 26; previous++ {
			if abs(current-previous) <= k {
				longest = max(longest, dp[previous]+1)
			}
		}
		dp[current] = max(dp[current], longest)
		result = max(result, dp[current])
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
