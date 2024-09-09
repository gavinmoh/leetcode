package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func spiralMatrix(m int, n int, head *ListNode) [][]int {
	matrix := make([][]int, m)
	for i := 0; i < m; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			matrix[i][j] = -1
		}
	}

	row, col := 0, 0
	direction := 'L'
	for head != nil {
		matrix[row][col] = head.Val
		head = head.Next
		if head == nil {
			break
		}

		for head != nil && matrix[row][col] != -1 {
			switch direction {
			case 'L':
				if col < n-1 && matrix[row][col+1] == -1 {
					col++
				} else {
					direction = 'D'
				}
			case 'D':
				if row < m-1 && matrix[row+1][col] == -1 {
					row++
				} else {
					direction = 'R'
				}
			case 'R':
				if col > 0 && matrix[row][col-1] == -1 {
					col--
				} else {
					direction = 'U'
				}
			case 'U':
				if row > 0 && matrix[row-1][col] == -1 {
					row--
				} else {
					direction = 'L'
				}
			}
		}
	}

	return matrix
}
