package main

func findJudge(n int, trust [][]int) int {
	trusts := make(map[int]int)
	trustedBy := make(map[int]int)
	for i := 1; i <= n; i++ {
		trusts[i] = 0
		trustedBy[i] = 0
	}

	for _, relationship := range trust {
		trusts[relationship[0]]++
		trustedBy[relationship[1]]++
	}

	for i := 1; i <= n; i++ {
		if trusts[i] == 0 && trustedBy[i] == n-1 {
			return i
		}
	}

	return -1
}
