package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiffWaysToCompute(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		expression := "2-1-1"
		expected := []int{2, 0}

		assert.Equal(t, expected, diffWaysToCompute(expression))
	})

	t.Run("test case 2", func(t *testing.T) {
		expression := "2*3-4*5"
		expected := []int{-34, -10, -14, -10, 10}

		assert.Equal(t, expected, diffWaysToCompute(expression))
	})

	t.Run("test case 3", func(t *testing.T) {
		expression := "0"
		expected := []int{0}

		assert.Equal(t, expected, diffWaysToCompute(expression))
	})

	t.Run("test case 4", func(t *testing.T) {
		expression := "2-1-1-1-1"
		expected := []int{2, 0, 2, 2, 4, 0, 2, 2, 0, 0, 2, 0, 0, -2}

		assert.Equal(t, expected, diffWaysToCompute(expression))
	})
}
