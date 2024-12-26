package main

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	cache := map[int]map[int]int{}
	var dfs func(idx int, acc int) int
	dfs = func(idx int, acc int) int {
		if idx >= len(nums) {
			if acc == target {
				return 1
			}

			return 0
		}

		if _, ok := cache[idx]; !ok {
			cache[idx] = map[int]int{}
		}

		if _, ok := cache[idx][acc+sum]; ok {
			return cache[idx][acc+sum]
		}

		add := dfs(idx+1, acc+nums[idx])
		subtract := dfs(idx+1, acc-nums[idx])

		cache[idx][acc+sum] = add + subtract

		return cache[idx][acc+sum]
	}

	return dfs(0, 0)
}
