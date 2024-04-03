package main

func exist(board [][]byte, word string) bool {
	m, n, l := len(board), len(board[0]), len(word)

	var dfs func(row int, col int, i int, visited map[int]map[int]bool) bool
	dfs = func(row int, col int, i int, visited map[int]map[int]bool) bool {
		if board[row][col] != word[i] {
			return false
		}

		if i == l-1 {
			return board[row][col] == word[i]
		}

		visited[row][col] = true
		// move left
		if (col > 0 && !visited[row][col-1]) && dfs(row, col-1, i+1, visited) {
			return true
		}

		// move right
		if (col < n-1 && !visited[row][col+1]) && dfs(row, col+1, i+1, visited) {
			return true
		}

		// move up
		if (row > 0 && !visited[row-1][col]) && dfs(row-1, col, i+1, visited) {
			return true
		}

		// move down
		if (row < m-1 && !visited[row+1][col]) && dfs(row+1, col, i+1, visited) {
			return true
		}

		visited[row][col] = false
		return false
	}

	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			visited := make(map[int]map[int]bool, m)
			for i := 0; i < m; i++ {
				visited[i] = make(map[int]bool, n)
			}
			if dfs(row, col, 0, visited) {
				return true
			}
		}
	}

	return false
}
