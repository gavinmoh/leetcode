package main

func pacificAtlantic(heights [][]int) [][]int {
	rows, cols := len(heights), len(heights[0])
	result := [][]int{}

	pacific := make([][]bool, rows)
	atlantic := make([][]bool, rows)
	for r := 0; r < rows; r++ {
		pacific[r] = make([]bool, cols)
		atlantic[r] = make([]bool, cols)
	}

	var dfs func(r int, c int, visited [][]bool, prevHeight int)
	dfs = func(r int, c int, visited [][]bool, prevHeight int) {
		if r < 0 || c < 0 || r == rows || c == cols || heights[r][c] < prevHeight || visited[r][c] {
			return
		}

		visited[r][c] = true
		dfs(r+1, c, visited, heights[r][c])
		dfs(r-1, c, visited, heights[r][c])
		dfs(r, c+1, visited, heights[r][c])
		dfs(r, c-1, visited, heights[r][c])
	}

	for c := 0; c < cols; c++ {
		dfs(0, c, pacific, heights[0][c])
		dfs(rows-1, c, atlantic, heights[rows-1][c])
	}
	for r := 0; r < rows; r++ {
		dfs(r, 0, pacific, heights[r][0])
		dfs(r, cols-1, atlantic, heights[r][cols-1])
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if pacific[r][c] && atlantic[r][c] {
				result = append(result, []int{r, c})
			}
		}
	}

	return result
}
