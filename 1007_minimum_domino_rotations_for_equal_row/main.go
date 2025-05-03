package main

import "math"

func minDominoRotations(tops []int, bottoms []int) int {
	n := len(tops)
	result := math.MaxInt

	candidates := []int{tops[0], bottoms[0]}
	for _, candidate := range candidates {
		presentInAllColumn := true
		for i := 0; i < n; i++ {
			if tops[i] != candidate && bottoms[i] != candidate {
				presentInAllColumn = false
				break
			}
		}

		if !presentInAllColumn {
			continue
		}

		rotateTop, rotateBottom := 0, 0
		for i := 0; i < n; i++ {
			if tops[i] != candidate {
				rotateTop++
			}

			if bottoms[i] != candidate {
				rotateBottom++
			}
		}

		result = min(result, min(rotateTop, rotateBottom))
	}

	if result == math.MaxInt {
		return -1
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
