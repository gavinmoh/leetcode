package main

func firstCompleteIndex(arr []int, mat [][]int) int {
	rows, cols := len(mat), len(mat[0])

	coordinates := map[int][]int{}
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			cell := mat[row][col]
			coordinates[cell] = []int{row, col}
		}
	}

	rowsCount, colsCount := map[int]int{}, map[int]int{}
	for i, cell := range arr {
		row, col := coordinates[cell][0], coordinates[cell][1]
		rowsCount[row]++
		if rowsCount[row] == cols {
			return i
		}
		colsCount[col]++
		if colsCount[col] == rows {
			return i
		}
	}

	return -1
}
