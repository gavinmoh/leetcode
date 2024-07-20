package main

func restoreMatrix(rowSum []int, colSum []int) [][]int {
	rows, cols := len(rowSum), len(colSum)
	matrix := make([][]int, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, cols)
	}

	// initialize first column
	for row := 0; row < rows; row++ {
		matrix[row][0] = rowSum[row]
	}

	for col := 0; col < cols-1; col++ {
		sum := 0
		for row := 0; row < rows; row++ {
			sum += matrix[row][col]
		}

		for row := 0; row < rows; row++ {
			diff := sum - colSum[col]
			matrix[row][col+1] = min(matrix[row][col], diff)
			matrix[row][col] -= matrix[row][col+1]
			sum -= matrix[row][col+1]
		}
	}

	return matrix
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
