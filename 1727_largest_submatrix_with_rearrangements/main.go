package main

func largestSubmatrix(matrix [][]int) int {
	m := len(matrix)
	n := len(matrix[0])
	previousHeights := [][]int{}
	largestArea := 0

	for row := 0; row < m; row++ {
		heights := [][]int{}
		seen := make([]bool, n)

		for _, prevHeight := range previousHeights {
			height := prevHeight[0]
			col := prevHeight[1]
			if matrix[row][col] == 1 {
				heights = append(heights, []int{height + 1, col})
				seen[col] = true
			}
		}

		for col := 0; col < n; col++ {
			if !seen[col] && matrix[row][col] == 1 {
				heights = append(heights, []int{1, col})
			}
		}

		for i, height := range heights {
			largestArea = max(largestArea, height[0]*(i+1))
		}

		previousHeights = heights
	}

	return largestArea
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
