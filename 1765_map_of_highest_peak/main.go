package main

func highestPeak(isWater [][]int) [][]int {
	rows, cols := len(isWater), len(isWater[0])
	result := make([][]int, rows)
	queue := [][]int{}
	for row := 0; row < rows; row++ {
		result[row] = make([]int, cols)
		for col := 0; col < cols; col++ {
			if isWater[row][col] == 1 {
				queue = append(queue, []int{row, col})
			}
		}
	}

	visited := map[int]map[int]bool{}
	level := 0
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

			result[row][col] = level

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

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
