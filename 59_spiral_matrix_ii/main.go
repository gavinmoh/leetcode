package main

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	direction := 'R'
	row, col := 0, 0
	for i := 1; i <= n*n; i++ {
		matrix[row][col] = i
		switch direction {
		case 'R':
			if col+1 == n || matrix[row][col+1] != 0 {
				row++
				direction = 'D'
			} else {
				col++
			}
		case 'D':
			if row+1 == n || matrix[row+1][col] != 0 {
				col--
				direction = 'L'
			} else {
				row++
			}
		case 'L':
			if col-1 < 0 || matrix[row][col-1] != 0 {
				row--
				direction = 'U'
			} else {
				col--
			}
		case 'U':
			if row-1 < 0 || matrix[row-1][col] != 0 {
				col++
				direction = 'R'
			} else {
				row--
			}
		}
	}

	return matrix
}
