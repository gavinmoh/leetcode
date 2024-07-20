package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestoreMatrix(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		rowSum := []int{3, 8}
		colSum := []int{4, 7}
		matrix := restoreMatrix(rowSum, colSum)

		rows, cols := len(rowSum), len(colSum)

		for row := 0; row < rows; row++ {
			sum := 0
			for col := 0; col < cols; col++ {
				sum += matrix[row][col]
			}
			assert.Equal(t, rowSum[row], sum)
		}

		for col := 0; col < cols; col++ {
			sum := 0
			for row := 0; row < rows; row++ {
				sum += matrix[row][col]
			}
			assert.Equal(t, colSum[col], sum)
		}
	})

	t.Run("test case 2", func(t *testing.T) {
		rowSum := []int{5, 7, 10}
		colSum := []int{8, 6, 8}
		matrix := restoreMatrix(rowSum, colSum)

		rows, cols := len(rowSum), len(colSum)

		for row := 0; row < rows; row++ {
			sum := 0
			for col := 0; col < cols; col++ {
				sum += matrix[row][col]
			}
			assert.Equal(t, rowSum[row], sum)
		}

		for col := 0; col < cols; col++ {
			sum := 0
			for row := 0; row < rows; row++ {
				sum += matrix[row][col]
			}
			assert.Equal(t, colSum[col], sum)
		}
	})
}
