package main

func solveSudoku(board [][]byte) {
	var dfs func(x, y int) bool
	dfs = func(x, y int) bool {
		if x == 8 && y == 9 {
			return true
		}

		// move to next row
		if y == 9 {
			x += 1
			y = 0
		}

		// continue next cell
		if board[x][y] != '.' {
			return dfs(x, y+1)
		}

		for i := 1; i <= 9; i++ {
			board[x][y] = '0' + byte(i)
			if isValidSudoku(board) && dfs(x, y+1) {
				return true
			}
			board[x][y] = '.'
		}

		return false
	}

	dfs(0, 0)
}

func isValidSudoku(board [][]byte) bool {
	// check rows
	for i := 0; i < 9; i++ {
		count := make([]int, 9)
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			num := int(board[i][j]-'0') - 1
			count[num]++

			if count[num] > 1 {
				return false
			}
		}
	}

	// check cols
	for i := 0; i < 9; i++ {
		count := make([]int, 9)
		for j := 0; j < 9; j++ {
			if board[j][i] == '.' {
				continue
			}

			num := int(board[j][i]-'0') - 1
			count[num]++

			if count[num] > 1 {
				return false
			}
		}
	}

	// check grid
	gridStarts := [][]int{
		{0, 0}, {0, 3}, {0, 6},
		{3, 0}, {3, 3}, {3, 6},
		{6, 0}, {6, 3}, {6, 6},
	}
	for _, start := range gridStarts {
		count := make([]int, 9)
		row, col := start[0], start[1]
		for i := row; i < row+3; i++ {
			for j := col; j < col+3; j++ {
				if board[i][j] == '.' {
					continue
				}

				num := int(board[i][j]-'0') - 1
				count[num]++

				if count[num] > 1 {
					return false
				}
			}
		}
	}

	return true
}
