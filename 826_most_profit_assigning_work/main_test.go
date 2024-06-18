package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfitAssignment(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		difficulty := []int{2, 4, 6, 8, 10}
		profit := []int{10, 20, 30, 40, 50}
		worker := []int{4, 5, 6, 7}
		expected := 100

		assert.Equal(t, expected, maxProfitAssignment(difficulty, profit, worker))
	})

	t.Run("test case 2", func(t *testing.T) {
		difficulty := []int{85, 47, 57}
		profit := []int{24, 66, 99}
		worker := []int{40, 25, 25}
		expected := 0

		assert.Equal(t, expected, maxProfitAssignment(difficulty, profit, worker))
	})
}
