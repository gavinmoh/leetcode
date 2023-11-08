package main

func isReachableAtTime(sx int, sy int, fx int, fy int, t int) bool {
	// edge case
	if sx == fx && sy == fy && t == 1 {
		return false
	}

	// calculate chebyshev distance
	dx := abs(sx - fx)
	dy := abs(sy - fy)
	d := max(dx, dy)

	return t >= d
}

// branchless style
func abs(x int) int {
	mask := x >> 31
	return (x + mask) ^ mask
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
