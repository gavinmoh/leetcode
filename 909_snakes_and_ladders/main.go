package main

func snakesAndLadders(board [][]int) int {
	n := len(board)
	queue := []int{1}
	visited := make([]bool, (n*n)+1)
	count := 0
	for len(queue) > 0 {
		newQueue := []int{}

		for _, label := range queue {
			if visited[label] {
				continue
			}
			visited[label] = true

			if label == n*n {
				return count
			}

			for i := 1; i <= 6; i++ {
				next := label + i
				if next > n*n {
					break
				}

				row, col := labelToCoordinate(next, n)
				if board[row][col] != -1 {
					newQueue = append(newQueue, board[row][col])
				} else {
					newQueue = append(newQueue, next)
				}
			}
		}

		count += 1
		queue = newQueue
	}

	return -1
}

func labelToCoordinate(x, n int) (row int, col int) {
	// work out which row
	row = x / n
	if x%n == 0 {
		row -= 1
	}

	// check if row is even or odd
	// even is left to right
	// odd is right to left
	col = x - (n * row) - 1
	if row%2 != 0 {
		col = n - col - 1
	}

	row = n - (row + 1)

	return
}
