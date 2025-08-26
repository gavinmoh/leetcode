package main

import "math"

func areaOfMaxDiagonal(dimensions [][]int) int {
	maxArea, maxLength := 0, float64(0)
	for _, dim := range dimensions {
		a, b := dim[0], dim[1]
		length := diagonal(a, b)
		if length > maxLength {
			maxLength = length
			maxArea = calculateArea(a, b)
		} else if length == maxLength {
			area := calculateArea(a, b)
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}

func diagonal(a, b int) float64 {
	return math.Sqrt(float64(a*a + b*b))
}

func calculateArea(a, b int) int {
	return a * b
}
