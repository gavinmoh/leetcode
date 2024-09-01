package main

func construct2DArray(original []int, m int, n int) [][]int {
	if len(original) != (m * n) {
		return [][]int{}
	}

	result := make([][]int, m)
	for row := 0; row < m; row++ {
		result[row] = make([]int, n)
		offset := row * n
		for col := 0; col < n; col++ {
			result[row][col] = original[offset+col]
		}
	}

	return result
}
