package main

import "strings"

func pushDominoes(dominoes string) string {
	n := len(dominoes)
	lefts, rights := make([]int, n), make([]int, n)

	// left to right
	force := 0
	for i := 0; i < n; i++ {
		switch dominoes[i] {
		case 'L':
			force = 0
		case 'R':
			force = n
		default:
			force = max(force-1, 0)
		}

		rights[i] = force
	}

	// right to left
	force = 0
	for i := n - 1; i >= 0; i-- {
		switch dominoes[i] {
		case 'L':
			force = n
		case 'R':
			force = 0
		default:
			force = max(force-1, 0)
		}

		lefts[i] = force
	}

	var result strings.Builder
	for i := 0; i < n; i++ {
		if lefts[i] == rights[i] {
			result.WriteRune('.')
		} else if lefts[i] > rights[i] {
			result.WriteRune('L')
		} else {
			result.WriteRune('R')
		}
	}

	return result.String()
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
