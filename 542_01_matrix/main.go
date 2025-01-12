package main

func updateMatrix(mat [][]int) [][]int {
	rows, cols := len(mat), len(mat[0])
	queue := [][]int{}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if mat[row][col] == 0 {
				queue = append(queue, []int{row, col})
			}
		}
	}

	// bfs
	level := 0
	visited := map[int]map[int]bool{}
	for len(queue) > 0 {
		newQueue := [][]int{}
		for _, cell := range queue {
			row, col := cell[0], cell[1]

			// check if visited
			if _, ok := visited[row]; !ok {
				visited[row] = map[int]bool{}
			}
			if _, ok := visited[row][col]; ok {
				continue
			} else {
				visited[row][col] = true
			}

			mat[row][col] = level

			if row > 0 {
				newQueue = append(newQueue, []int{row - 1, col})
			}

			if row < rows-1 {
				newQueue = append(newQueue, []int{row + 1, col})
			}

			if col > 0 {
				newQueue = append(newQueue, []int{row, col - 1})
			}

			if col < cols-1 {
				newQueue = append(newQueue, []int{row, col + 1})
			}
		}
		level++
		queue = newQueue
	}

	return mat
}
