package main

func countServers(grid [][]int) int {
	rowsCount := map[int]int{}
	colsCount := map[int]int{}

	for row, columns := range grid {
		for col, cell := range columns {
			if cell == 1 {
				rowsCount[row]++
				colsCount[col]++
			}
		}
	}

	count := 0
	for row, columns := range grid {
		for col, cell := range columns {
			if cell == 1 && (rowsCount[row] > 1 || colsCount[col] > 1) {
				count++
			}
		}
	}

	return count
}
