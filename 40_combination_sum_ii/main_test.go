package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombinationSum2(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		candidates := []int{10, 1, 2, 7, 6, 1, 5}
		target := 8
		expected := [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}}

		assert.Equal(t, expected, combinationSum2(candidates, target))
	})

	t.Run("test case 2", func(t *testing.T) {
		candidates := []int{2, 5, 2, 1, 2}
		target := 5
		expected := [][]int{{1, 2, 2}, {5}}

		assert.Equal(t, expected, combinationSum2(candidates, target))
	})
}
