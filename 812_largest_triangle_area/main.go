package main

func largestTriangleArea(points [][]int) float64 {
	n := len(points)
	maxArea := float64(-1)

	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				x1, x2, x3 := points[i][0], points[j][0], points[k][0]
				y1, y2, y3 := points[i][1], points[j][1], points[k][1]

				area := 0.5 * float64(abs(x1*(y2-y3)+x2*(y3-y1)+x3*(y1-y2)))
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}

	return maxArea
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
