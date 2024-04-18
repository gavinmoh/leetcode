package main

func islandPerimeter(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	perimeter := 0
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == 0 {
				continue
			}

			// check left
			if col-1 < 0 || grid[row][col-1] == 0 {
				perimeter++
			}

			// check right
			if col+1 >= cols || grid[row][col+1] == 0 {
				perimeter++
			}

			// check top
			if row-1 < 0 || grid[row-1][col] == 0 {
				perimeter++
			}

			// check bottom
			if row+1 >= rows || grid[row+1][col] == 0 {
				perimeter++
			}
		}
	}

	return perimeter
}
