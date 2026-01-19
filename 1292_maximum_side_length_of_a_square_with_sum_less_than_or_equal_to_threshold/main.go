package main

func maxSideLength(mat [][]int, threshold int) int {
	rows, cols := len(mat), len(mat[0])
	prefixSum := make([][]int, rows)
	for i := range prefixSum {
		prefixSum[i] = make([]int, cols+1)
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			prefixSum[row][col+1] = prefixSum[row][col] + mat[row][col]
		}
	}

	checkSquare := func(i, j, l int) bool {
		sum := 0
		for row := i; row < i+l; row++ {
			sum += prefixSum[row][j+l] - prefixSum[row][j]
		}

		return sum <= threshold
	}

	result := 0
	left, right := 0, min(rows, cols)
	for left <= right {
		mid := (left + right) / 2
		found := false

	outer:
		for row := 0; row+mid <= rows; row++ {
			for col := 0; col+mid <= cols; col++ {
				if checkSquare(row, col, mid) {
					found = true
					break outer
				}
			}
		}

		if found {
			result = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
