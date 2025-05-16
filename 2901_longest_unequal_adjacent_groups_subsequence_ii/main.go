package main

func getWordsInLongestSubsequence(words []string, groups []int) []string {
	n := len(groups)
	dp := make([]int, n)
	prev := make([]int, n)
	for i := range dp {
		dp[i] = 1
		prev[i] = -1
	}
	maxIndex := 0

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if differentByOneCharacter(words[i], words[j]) && dp[j]+1 > dp[i] && groups[i] != groups[j] {
				dp[i] = dp[j] + 1
				prev[i] = j
			}

		}

		if dp[i] > dp[maxIndex] {
			maxIndex = i
		}
	}

	result := []string{}
	for i := maxIndex; i >= 0; i = prev[i] {
		result = append(result, words[i])
	}
	reverse(result)

	return result
}

func differentByOneCharacter(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	count := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
		}

		if count > 1 {
			return false
		}
	}

	return count == 1
}

func reverse(arr []string) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
