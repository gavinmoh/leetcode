package main

func minimumDeletions(s string) int {
	n := len(s)
	aCount := make([]int, n)
	bCount := make([]int, n)
	aTotal, bTotal := 0, 0

	for i := 0; i < n; i++ {
		bCount[i] = bTotal
		if s[i] == 'b' {
			bTotal++
		}
	}

	for i := n - 1; i >= 0; i-- {
		aCount[i] = aTotal
		if s[i] == 'a' {
			aTotal++
		}
	}

	minDeletions := n
	for i := 0; i < n; i++ {
		minDeletions = min(minDeletions, aCount[i]+bCount[i])
	}

	return minDeletions
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
