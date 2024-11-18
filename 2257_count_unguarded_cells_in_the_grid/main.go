package main

const (
	EMPTY = iota
	GUARDED
	GUARD
	WALL
)

func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
	rows, cols := m, n
	grid := make([][]int, rows)
	for row := 0; row < rows; row++ {
		grid[row] = make([]int, cols)
	}

	for _, guard := range guards {
		row, col := guard[0], guard[1]
		grid[row][col] = GUARD
	}

	for _, wall := range walls {
		row, col := wall[0], wall[1]
		grid[row][col] = WALL
	}

	for _, guard := range guards {
		row, col := guard[0], guard[1]
		// check top
		for i := row - 1; i >= 0; i-- {
			if grid[i][col] == WALL || grid[i][col] == GUARD {
				break
			}

			grid[i][col] = GUARDED
		}

		// check bottom
		for i := row + 1; i < rows; i++ {
			if grid[i][col] == WALL || grid[i][col] == GUARD {
				break
			}

			grid[i][col] = GUARDED
		}

		// check left
		for i := col - 1; i >= 0; i-- {
			if grid[row][i] == WALL || grid[row][i] == GUARD {
				break
			}

			grid[row][i] = GUARDED
		}

		// check right
		for i := col + 1; i < cols; i++ {
			if grid[row][i] == WALL || grid[row][i] == GUARD {
				break
			}

			grid[row][i] = GUARDED
		}
	}

	count := 0
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == EMPTY {
				count++
			}
		}
	}

	return count
}
