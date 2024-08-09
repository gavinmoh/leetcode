package main

func numMagicSquaresInside(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	count := 0
	for row := 0; row < rows-2; row++ {
		for col := 0; col < cols-2; col++ {
			subgrid := make([][]int, 3)
			for i := 0; i < 3; i++ {
				subgrid[i] = grid[row+i][col : col+3]
			}

			if !isDistinct(subgrid) {
				continue
			}

			rowSum := sumRow(subgrid)
			if rowSum == -1 {
				continue
			}

			colSum := sumCol(subgrid)
			if colSum == -1 {
				continue
			}
			if rowSum != colSum {
				continue
			}

			diagonalSum := sumDiagonal(subgrid)
			if diagonalSum == -1 {
				continue
			}
			if colSum != diagonalSum {
				continue
			}
			count += 1
		}
	}

	return count
}

func isDistinct(grid [][]int) bool {
	freq := make([]int, 9)
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			num := grid[row][col] - 1
			if num < 0 || num > 8 {
				return false
			}
			if freq[num] != 0 {
				return false
			}
			freq[num] = 1
		}
	}
	return true
}

func sumRow(grid [][]int) int {
	sum := -1
	for row := 0; row < 3; row++ {
		currentSum := 0
		for col := 0; col < 3; col++ {
			currentSum += grid[row][col]
		}
		if sum == -1 {
			sum = currentSum
		} else if sum != currentSum {
			return -1
		}
	}
	return sum
}

func sumCol(grid [][]int) int {
	sum := -1
	for col := 0; col < 3; col++ {
		currentSum := 0
		for row := 0; row < 3; row++ {
			currentSum += grid[row][col]
		}
		if sum == -1 {
			sum = currentSum
		} else if sum != currentSum {
			return -1
		}
	}
	return sum
}

func sumDiagonal(grid [][]int) int {
	left, right := 0, 0
	for i := 0; i < 3; i++ {
		left += grid[i][i]
		right += grid[i][2-i]
	}

	if left != right {
		return -1
	}

	return left
}
