package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstruct2DArray(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		original := []int{1, 2, 3, 4}
		m := 2
		n := 2
		expected := [][]int{{1, 2}, {3, 4}}

		assert.Equal(t, expected, construct2DArray(original, m, n))
	})

	t.Run("test case 2", func(t *testing.T) {
		original := []int{1, 2, 3}
		m := 2
		n := 2
		expected := [][]int{}

		assert.Equal(t, expected, construct2DArray(original, m, n))
	})

	t.Run("test case 3", func(t *testing.T) {
		original := []int{1, 2}
		m := 1
		n := 1
		expected := [][]int{}

		assert.Equal(t, expected, construct2DArray(original, m, n))
	})
}
