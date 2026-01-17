package main

func largestSquareArea(bottomLeft [][]int, topRight [][]int) int64 {
	n := len(bottomLeft)
	maxArea := int64(0)

	for i := 0; i < n+1; i++ {
		for j := i + 1; j < n; j++ {
			xl := max(bottomLeft[i][0], bottomLeft[j][0])
			yl := max(bottomLeft[i][1], bottomLeft[j][1])
			xr := min(topRight[i][0], topRight[j][0])
			yr := min(topRight[i][1], topRight[j][1])

			// return if no overlap
			if xl >= xr || yl >= yr {
				continue
			}

			l := min(xr-xl, yr-yl)
			area := int64(l) * int64(l)
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
