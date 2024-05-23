package main

import "math"

func beautifulSubsets(nums []int, k int) int {
	count := make(map[int]int)

	for _, num := range nums {
		count[num] += 1
	}

	groups := []map[int]int{}
	visited := make(map[int]bool)

	for num := range count {
		if _, ok := visited[num]; ok {
			continue
		}
		group := make(map[int]int)
		for {
			if _, ok := count[num-k]; ok {
				num -= k
			} else {
				break
			}
		}

		for {
			if _, ok := count[num]; ok {
				group[num] = count[num]
				visited[num] = true
				num += k
			} else {
				break
			}
		}

		groups = append(groups, group)
	}

	cache := make(map[int]int)
	var helper func(n int, group map[int]int) int
	helper = func(n int, group map[int]int) int {
		if _, ok := group[n]; !ok {
			return 1
		}
		if cachedResult, ok := cache[n]; ok {
			return cachedResult
		}
		skip := helper(n+k, group)
		freq := group[n]
		include := combinations(freq) * helper(n+(2*k), group)
		res := skip + include
		cache[n] = res
		return res
	}

	result := 1
	for _, group := range groups {
		groupMin := math.MaxInt
		for k := range group {
			if k < groupMin {
				groupMin = k
			}
		}
		result *= helper(groupMin, group)
	}

	return result - 1
}

func combinations(n int) int {
	if n == 1 {
		return 2 - 1
	}

	result := 2
	for i := 2; i <= n; i++ {
		result *= 2
	}

	return result - 1
}
