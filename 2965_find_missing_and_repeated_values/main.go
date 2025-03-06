package main

func findMissingAndRepeatedValues(grid [][]int) []int {
	n := len(grid)
	count := make([]int, (n*n)+1)
	for _, row := range grid {
		for _, num := range row {
			count[num]++
		}
	}

	result := make([]int, 2)
	for i := 1; i <= n*n; i++ {
		if count[i] > 1 {
			result[0] = i
		}

		if count[i] == 0 {
			result[1] = i
		}
	}

	return result
}
