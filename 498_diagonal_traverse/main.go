package main

func findDiagonalOrder(mat [][]int) []int {
	rows, cols := len(mat), len(mat[0])
	up := true

	result := make([]int, rows*cols)
	for row, col, i := 0, 0, 0; i < rows*cols; i++ {
		result[i] = mat[row][col]

		nextRow, nextCol := -1, -1
		if up {
			nextRow = row - 1
			nextCol = col + 1
		} else {
			nextRow = row + 1
			nextCol = col - 1
		}

		if nextRow < 0 || nextRow >= rows || nextCol < 0 || nextCol >= cols {
			if up {
				if col == cols-1 {
					row++
				} else if col < cols-1 {
					col++
				}
			} else {
				if row == rows-1 {
					col++
				} else if row < rows-1 {
					row++
				}
			}

			// flip direction
			up = !up
		} else {
			row, col = nextRow, nextCol
		}
	}

	return result
}
