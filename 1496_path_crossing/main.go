package main

func isPathCrossing(path string) bool {
	visited := [][]int{{0, 0}}
	current := []int{0, 0}

	for _, direction := range path {
		switch direction {
		case 'N':
			current[1]++
		case 'S':
			current[1]--
		case 'E':
			current[0]++
		case 'W':
			current[0]--
		}
		for _, point := range visited {
			if point[0] == current[0] && point[1] == current[1] {
				return true
			}
		}
		visited = append(visited, []int{current[0], current[1]})
	}

	return false
}
