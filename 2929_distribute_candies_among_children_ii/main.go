package main

func distributeCandies(n int, limit int) int64 {
	answer := int64(0)
	for i := 0; i <= min(n, limit); i++ {
		leftOver := n - i
		if leftOver > limit*2 {
			continue
		}

		answer += int64(max(min(limit, leftOver)-max(0, leftOver-limit)+1, 0))
	}

	return answer
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
