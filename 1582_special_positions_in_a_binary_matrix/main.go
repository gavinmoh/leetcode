package main

func numSpecial(mat [][]int) int {
	m := len(mat)
	n := len(mat[0])
	count := 0
	ones := [][]int{}

	// keep track of all the ones
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 1 {
				ones = append(ones, []int{i, j})
			}
		}
	}

	// check for each ones
	for _, one := range ones {
		row := one[0]
		col := one[1]
		// checking row
		special := true
		for i := 0; i < n; i++ {
			if i == col {
				continue
			}
			cell := mat[row][i]
			if cell == 1 {
				special = false
				break
			}
		}
		if !special {
			continue
		}
		// check for cols
		for i := 0; i < m; i++ {
			if i == row {
				continue
			}
			cell := mat[i][col]
			if cell == 1 {
				special = false
				break
			}
		}
		if special {
			count++
		}
	}

	return count
}
