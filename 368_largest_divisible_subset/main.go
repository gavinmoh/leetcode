package main

import "sort"

func largestDivisibleSubset(nums []int) []int {
	n := len(nums)
	sort.Ints(nums)

	cache := make(map[int]map[int][]int)

	var dfs func(int, int) []int
	dfs = func(i int, prev int) []int {
		if i == n {
			return []int{}
		}

		if _, ok := cache[i]; !ok {
			cache[i] = make(map[int][]int)
		}

		if cachedResult, ok := cache[i][prev]; ok {
			return cachedResult
		}

		num := nums[i]
		left := dfs(i+1, prev)
		if num%prev == 0 {
			right := append([]int{num}, dfs(i+1, num)...)

			if len(left) > len(right) {
				cache[i][prev] = left
			} else {
				cache[i][prev] = right
			}
		} else {
			cache[i][prev] = left
		}

		return cache[i][prev]
	}

	return dfs(0, 1)
}
