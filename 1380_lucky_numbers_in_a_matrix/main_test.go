package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLuckyNumbers(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		matrix := [][]int{
			{3, 7, 8},
			{9, 11, 13},
			{15, 16, 17},
		}
		expected := []int{15}

		assert.Equal(t, expected, luckyNumbers(matrix))
	})

	t.Run("test case 2", func(t *testing.T) {
		matrix := [][]int{
			{1, 10, 4, 2},
			{9, 3, 8, 7},
			{15, 16, 17, 12},
		}
		expected := []int{12}

		assert.Equal(t, expected, luckyNumbers(matrix))
	})

	t.Run("test case 3", func(t *testing.T) {
		matrix := [][]int{
			{7, 8},
			{1, 2},
		}
		expected := []int{7}

		assert.Equal(t, expected, luckyNumbers(matrix))
	})
}
