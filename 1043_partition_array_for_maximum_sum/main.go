package main

func maxSumAfterPartitioning(arr []int, k int) int {
	n := len(arr)
	cache := make(map[int]int)

	var dfs func(int) int
	dfs = func(i int) int {
		if i >= n {
			return 0
		}

		if cachedResult, ok := cache[i]; ok {
			return cachedResult
		}

		maxSum := 0
		for j := 1; j <= k; j++ {
			if i+j > n {
				break
			}

			subArr := arr[i : i+j]
			currentMax := -1
			for _, num := range subArr {
				currentMax = max(currentMax, num)
			}
			sum := currentMax*j + dfs(i+j)
			if sum > maxSum {
				maxSum = sum
			}
		}

		cache[i] = maxSum
		return maxSum
	}

	return dfs(0)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
