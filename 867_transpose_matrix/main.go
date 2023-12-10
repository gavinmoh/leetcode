package main

func transpose(matrix [][]int) [][]int {
	rows := len(matrix)
	cols := len(matrix[0])

	transposed := make([][]int, cols)
	for col := 0; col < cols; col++ {
		transposed[col] = make([]int, rows)
		for row := 0; row < rows; row++ {
			transposed[col][row] = matrix[row][col]
		}
	}

	return transposed
}
