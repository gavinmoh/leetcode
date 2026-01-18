package main

func largestMagicSquare(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])

	// construct prefix sum
	rowsPrefix := make([][]int, rows)
	colsPrefix := make([][]int, rows+1)

	for i := range rowsPrefix {
		rowsPrefix[i] = make([]int, cols+1)
	}

	for i := range colsPrefix {
		colsPrefix[i] = make([]int, cols)
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			rowsPrefix[row][col+1] = rowsPrefix[row][col] + grid[row][col]
			colsPrefix[row+1][col] = colsPrefix[row][col] + grid[row][col]
		}
	}

	checkSquare := func(i, j, k int) bool {
		// use left diagonal sum as magic sum
		magicSum := 0
		for row, col := i, j; row < i+k && col < j+k; {
			magicSum += grid[row][col]
			row++
			col++
		}

		// check right diagonal sum
		dSum := 0
		for row, col := i, j+k-1; row < i+k && col >= j; {
			dSum += grid[row][col]
			row++
			col--
		}

		if dSum != magicSum {
			return false
		}

		// check all row sum
		for row := i; row < i+k; row++ {
			sum := rowsPrefix[row][j+k] - rowsPrefix[row][j]
			if sum != magicSum {
				return false
			}
		}

		// check all col sum
		for col := j; col < j+k; col++ {
			sum := colsPrefix[i+k][col] - colsPrefix[i][col]
			if sum != magicSum {
				return false
			}
		}

		return true
	}

	maxK := 1
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			k := min(rows-row, cols-col)
			for k > 1 {
				if checkSquare(row, col, k) {
					break
				}
				k--
			}
			if k > maxK {
				maxK = k
			}
		}
	}

	return maxK
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
