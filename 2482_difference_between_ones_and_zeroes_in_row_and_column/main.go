package main

func onesMinusZeros(grid [][]int) [][]int {
	m := len(grid)
	n := len(grid[0])

	// calculate all the ones and zeroes
	onesRow := make(map[int]int, m)
	zerosRow := make(map[int]int, m)
	onesCol := make(map[int]int, n)
	zerosCol := make(map[int]int, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				onesRow[i]++
				onesCol[j]++
			} else {
				zerosRow[i]++
				zerosCol[j]++
			}
		}
	}

	// initialize diff
	diff := make([][]int, m)
	for i := 0; i < m; i++ {
		diff[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			diff[i][j] = onesRow[i] + onesCol[j] - zerosRow[i] - zerosCol[j]
		}
	}

	return diff
}
