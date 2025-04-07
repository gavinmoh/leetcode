package main

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum%2 != 0 {
		return false
	}

	cache := map[int]map[int]bool{}

	var dfs func(i, target int) bool
	dfs = func(i, target int) bool {
		if i >= len(nums) {
			return target == 0
		}

		if target < 0 {
			return false
		}

		if _, ok := cache[i]; !ok {
			cache[i] = map[int]bool{}
		}

		if cachedResult, ok := cache[i][target]; ok {
			return cachedResult
		}

		cache[i][target] = dfs(i+1, target-nums[i]) || dfs(i+1, target)

		return cache[i][target]
	}

	return dfs(0, sum/2)
}
