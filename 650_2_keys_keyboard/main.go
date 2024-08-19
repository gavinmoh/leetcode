package main

func minSteps(n int) int {
	var dfs func(steps int, clipboard int, count int) int
	dfs = func(steps int, clipboard int, count int) int {
		if count == n {
			return steps
		}
		if count > n {
			return -1
		}

		// paste
		left := -1
		if clipboard > 0 {
			left = dfs(steps+1, clipboard, count+clipboard)
		}
		// copy and paste
		right := dfs(steps+2, count, count*2)

		if left == -1 {
			return right
		}
		if right == -1 {
			return left
		}
		return min(left, right)
	}

	return dfs(0, 0, 1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
