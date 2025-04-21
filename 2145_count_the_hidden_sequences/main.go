package main

func numberOfArrays(differences []int, lower int, upper int) int {
	x, y, current := 0, 0, 0
	for _, diff := range differences {
		current += diff
		x, y = min(x, current), max(y, current)
		if y-x > upper-lower {
			return 0
		}
	}

	return (upper - lower) - (y - x) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
