package main

func findOrder(numCourses int, prerequisites [][]int) []int {
	courses := make([][]int, numCourses) // adjacency list
	for i := 0; i < numCourses; i++ {
		courses[i] = []int{}
	}

	for _, edge := range prerequisites {
		courses[edge[0]] = append(courses[edge[0]], edge[1])
	}

	taken := make([]bool, numCourses)
	result := []int{}
	var dfs func(course int, visited []bool) bool
	dfs = func(course int, visited []bool) bool {
		visited[course] = true
		for _, prereq := range courses[course] {
			if taken[prereq] {
				continue
			}

			if visited[prereq] || !dfs(prereq, visited) {
				return false
			}
		}

		if !taken[course] {
			result = append(result, course)
			taken[course] = true
		}

		return true
	}

	for i := 0; i < numCourses; i++ {
		if taken[i] {
			continue
		}

		visited := make([]bool, numCourses)
		if !dfs(i, visited) {
			return []int{}
		}
	}

	return result
}
