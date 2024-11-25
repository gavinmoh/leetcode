package main

import (
	"fmt"
	"strings"
)

// [0, 1, 2]
// [3, 4, 5]
var directions = map[int][]int{
	0: {1, 3},
	1: {0, 2, 4},
	2: {1, 5},
	3: {0, 4},
	4: {1, 3, 5},
	5: {2, 4},
}

func slidingPuzzle(board [][]int) int {
	count := 0
	boards := []string{toString(board)}
	visited := map[string]bool{}

	for len(boards) > 0 {
		newBoards := []string{}
		for _, b := range boards {
			if isSolved(b) {
				return count
			}

			if _, ok := visited[b]; ok {
				continue
			} else {
				visited[b] = true
			}

			// find position of 0
			var position int
			for i, char := range b {
				if char == '0' {
					position = i
					break
				}
			}

			// generate all variants
			for _, swap := range directions[position] {
				newBoard := []byte(b)
				newBoard[position], newBoard[swap] = newBoard[swap], newBoard[position]
				newBoards = append(newBoards, string(newBoard))
			}
		}
		count++
		boards = newBoards
	}

	return -1
}

func toString(board [][]int) string {
	var result strings.Builder
	for _, row := range board {
		for _, num := range row {
			result.WriteString(fmt.Sprintf("%d", num))
		}
	}
	return result.String()
}

func isSolved(boardStr string) bool {
	return boardStr == "123450"
}
