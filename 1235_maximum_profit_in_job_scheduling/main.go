package main

import (
	"sort"
)

type Job struct {
	startTime int
	endTime   int
	profit    int
}

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	n := len(startTime)
	jobs := make([]Job, n)

	for i := 0; i < n; i++ {
		jobs[i] = Job{startTime: startTime[i], endTime: endTime[i], profit: profit[i]}
	}

	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].startTime < jobs[j].startTime
	})

	binarySearch := func(left, target int) int {
		right := n - 1

		for left <= right {
			mid := (left + right) / 2
			if jobs[mid].startTime < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		if left >= n {
			return -1
		}
		return left
	}

	cache := make(map[int]int)
	var dfs func(i int) int
	dfs = func(i int) int {
		if i >= n {
			return 0
		}

		job := jobs[i]
		if _, ok := cache[i]; !ok {
			next := binarySearch(i+1, job.endTime)
			left := job.profit
			if next != -1 {
				left += dfs(next)
			}
			right := dfs(i + 1)
			cache[i] = max(left, right)
		}

		return cache[i]
	}

	return dfs(0)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
