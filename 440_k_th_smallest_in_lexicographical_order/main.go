package main

func findKthNumber(n int, k int) int {
	current, i := 1, 1

	count := func(cur int) int {
		result, neighbour := 0, cur+1
		for cur <= n {
			result += min(neighbour, n+1) - cur
			cur *= 10
			neighbour *= 10
		}
		return result
	}

	for i < k {
		steps := count(current)
		if i+steps <= k {
			current += 1
			i += steps
		} else {
			current *= 10
			i += 1
		}
	}

	return current
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
