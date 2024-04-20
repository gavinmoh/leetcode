package main

func findFarmland(land [][]int) [][]int {
	rows, cols := len(land), len(land[0])
	groups := [][]int{}

	var dfs func(row int, col int, group *[]int)
	dfs = func(row int, col int, group *[]int) {
		// mark as seen
		land[row][col] = 0

		// update of maxRow and maxCol
		if row > (*group)[2] {
			(*group)[2] = row
		}
		if col > (*group)[3] {
			(*group)[3] = col
		}

		// move right
		if col+1 < cols && land[row][col+1] == 1 {
			dfs(row, col+1, group)
		}

		// move down
		if row+1 < rows && land[row+1][col] == 1 {
			dfs(row+1, col, group)
		}
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if land[row][col] == 0 {
				continue
			}
			group := []int{row, col, row, col}
			dfs(row, col, &group)
			groups = append(groups, group)
		}
	}

	return groups
}
