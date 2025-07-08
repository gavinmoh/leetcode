package main

import "sort"

func maxValue(events [][]int, k int) int {
	n := len(events)

	sort.SliceStable(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})

	nextEvents := make([]int, n)
	binarySearch := func(startDay int) int {
		left, right := 0, len(events)
		for left < right {
			mid := left + (right-left)/2
			if events[mid][0] <= startDay {
				left = mid + 1
			} else {
				right = mid
			}
		}

		return left
	}

	// precompute the next event for each event
	for i, event := range events {
		nextEvents[i] = binarySearch(event[1])
	}

	// initialize dp
	dp := make([][]int, k+1)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = -1
		}
	}

	var dfs func(idx, k int) int
	dfs = func(idx, k int) int {
		// skip if out of bound or k is 0 (cannot attend anymore)
		if idx == n || k == 0 {
			return 0
		}

		// if result already computed, use cached result
		if dp[k][idx] != -1 {
			return dp[k][idx]
		}

		// find the next event index
		next := nextEvents[idx]

		// attend current event
		left := events[idx][2] + dfs(next, k-1)

		// skip current event
		right := dfs(idx+1, k)

		if left > right {
			dp[k][idx] = left
		} else {
			dp[k][idx] = right
		}

		return dp[k][idx]
	}

	return dfs(0, k)
}
