package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombinationSum(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		candidates := []int{2, 3, 6, 7}
		target := 7
		expected := [][]int{{2, 2, 3}, {7}}

		assert.Equal(t, expected, combinationSum(candidates, target))
	})

	t.Run("test case 2", func(t *testing.T) {
		candidates := []int{2, 3, 5}
		target := 8
		expected := [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}

		assert.Equal(t, expected, combinationSum(candidates, target))
	})

	t.Run("test case 3", func(t *testing.T) {
		candidates := []int{2}
		target := 1
		expected := [][]int{}

		assert.Equal(t, expected, combinationSum(candidates, target))
	})

	t.Run("test case 4", func(t *testing.T) {
		candidates := []int{3, 2, 6, 7}
		target := 11
		expected := [][]int{{3, 3, 3, 2}, {3, 2, 2, 2, 2}, {3, 2, 6}, {2, 2, 7}}

		assert.Equal(t, expected, combinationSum(candidates, target))
	})
}
