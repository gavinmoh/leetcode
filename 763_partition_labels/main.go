package main

func partitionLabels(s string) []int {
	last := map[rune]int{}
	for i, c := range s {
		last[c] = i
	}

	result := []int{}

	left, right := 0, -1
	for i, c := range s {
		right = max(right, last[c])

		if i == right {
			result = append(result, right-left+1)
			left, right = i+1, -1
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
