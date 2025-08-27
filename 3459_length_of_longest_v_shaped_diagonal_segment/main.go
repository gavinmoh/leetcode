package main

func lenOfVDiagonal(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	directions := [][]int{{-1, -1}, {-1, 1}, {1, 1}, {1, -1}}

	isWithinBound := func(row, col int) bool {
		if row >= 0 && row < rows && col >= 0 && col < cols {
			return true
		}

		return false
	}

	nextSeq := func(num int) int {
		if num == 0 {
			return 2
		}
		return 0
	}

	var dfs func(row, col, dir int, turned bool) int
	dfs = func(row, col, dir int, turned bool) int {
		left, right := 0, 0
		nextNum := nextSeq(grid[row][col])

		// continue
		nextRow, nextCol := row+directions[dir][0], col+directions[dir][1]
		if isWithinBound(nextRow, nextCol) && grid[nextRow][nextCol] == nextNum {
			left = dfs(nextRow, nextCol, dir, turned) + 1
		}

		// turn if not yet turn
		if !turned {
			newDir := (dir + 1) % 4
			nextRow, nextCol := row+directions[newDir][0], col+directions[newDir][1]
			if isWithinBound(nextRow, nextCol) && grid[nextRow][nextCol] == nextNum {
				right = dfs(nextRow, nextCol, newDir, !turned) + 1
			}
		}

		if left > right {
			return left
		}

		return right
	}

	maxLength := 0
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] != 1 {
				continue
			}

			if maxLength == 0 {
				maxLength = 1
			}

			for dir := range directions {
				newRow, newCol := row+directions[dir][0], col+directions[dir][1]
				if !isWithinBound(newRow, newCol) {
					continue
				}

				if grid[newRow][newCol] != 2 {
					continue
				}

				length := dfs(newRow, newCol, dir, false) + 2
				if length > maxLength {
					maxLength = length
				}
			}
		}
	}

	return maxLength
}
