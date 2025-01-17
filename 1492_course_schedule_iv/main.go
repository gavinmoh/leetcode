package main

func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	courses := map[int]map[int]bool{}
	for i := 0; i < numCourses; i++ {
		courses[i] = map[int]bool{}
	}

	for _, prerequisite := range prerequisites {
		left, right := prerequisite[0], prerequisite[1]
		courses[right][left] = true
	}

	// dfs
	// determine is course a is prerequisite of course b
	cache := map[int]map[int]bool{}
	var isPrerequisite func(a, b int) bool
	isPrerequisite = func(a, b int) bool {
		if _, ok := cache[a]; !ok {
			cache[a] = map[int]bool{}
		}

		if cachedResult, ok := cache[a][b]; ok {
			return cachedResult
		}

		for prerequisite := range courses[b] {
			if prerequisite == a || isPrerequisite(a, prerequisite) {
				cache[a][b] = true
				return true
			}
		}

		cache[a][b] = false
		return false
	}

	result := make([]bool, len(queries))
	for i, query := range queries {
		left, right := query[0], query[1]
		result[i] = isPrerequisite(left, right)
	}

	return result
}
