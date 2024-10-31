package main

import "sort"

func minimumTotalDistance(robot []int, factory [][]int) int64 {
	sort.Ints(robot)
	sort.SliceStable(factory, func(i, j int) bool {
		return factory[i][0] < factory[j][0]
	})
	factoryPositions := []int{}
	for _, f := range factory {
		for i := 0; i < f[1]; i++ {
			factoryPositions = append(factoryPositions, f[0])
		}
	}

	cache := map[int]map[int]int64{}
	var dfs func(robotIdx, factoryIdx int) int64
	dfs = func(robotIdx, factoryIdx int) int64 {
		if robotIdx == len(robot) {
			return 0
		}

		if factoryIdx == len(factoryPositions) {
			return 1_000_000_000_000
		}

		if _, ok := cache[robotIdx]; !ok {
			cache[robotIdx] = map[int]int64{}
		}

		if cachedResult, ok := cache[robotIdx][factoryIdx]; ok {
			return cachedResult
		}

		left := int64(abs(robot[robotIdx]-factoryPositions[factoryIdx])) + dfs(robotIdx+1, factoryIdx+1)
		right := dfs(robotIdx, factoryIdx+1)

		result := min(left, right)
		cache[robotIdx][factoryIdx] = result

		return result
	}

	return dfs(0, 0)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
