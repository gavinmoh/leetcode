package main

func setZeroes(matrix [][]int) {
	rows, cols := len(matrix), len(matrix[0])
	rowsToClear, columnsToClear := make([]bool, rows), make([]bool, cols)
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if matrix[row][col] == 0 {
				rowsToClear[row] = true
				columnsToClear[col] = true
			}
		}
	}

	for row, clear := range rowsToClear {
		if clear {
			for col := 0; col < cols; col++ {
				matrix[row][col] = 0
			}
		}
	}

	for col, clear := range columnsToClear {
		if clear {
			for row := 0; row < rows; row++ {
				matrix[row][col] = 0
			}
		}
	}
}
