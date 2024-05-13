package main

func matrixScore(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])

	for row := 0; row < rows; row++ {
		if grid[row][0] == 0 {
			for col := 0; col < cols; col++ {
				grid[row][col] = flip(grid[row][col])
			}
		}
	}

	for col := 0; col < cols; col++ {
		zeroes, ones := 0, 0
		for row := 0; row < rows; row++ {
			if grid[row][col] == 0 {
				zeroes++
			} else {
				ones++
			}
		}
		if zeroes > ones {
			for row := 0; row < rows; row++ {
				grid[row][col] = flip(grid[row][col])
			}
		}
	}

	sum := 0
	for row := 0; row < rows; row++ {
		score := 0
		for col := 0; col < cols; col++ {
			score += grid[row][col] << (cols - col - 1)
		}
		sum += score
	}

	return sum
}

func flip(x int) int {
	if x == 0 {
		return 1
	}
	return 0
}
