package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumAbsDifference(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		arr := []int{4, 2, 1, 3}
		expected := [][]int{{1, 2}, {2, 3}, {3, 4}}

		assert.Equal(t, expected, minimumAbsDifference(arr))
	})

	t.Run("test case 2", func(t *testing.T) {
		arr := []int{1, 3, 6, 10, 15}
		expected := [][]int{{1, 3}}

		assert.Equal(t, expected, minimumAbsDifference(arr))
	})

	t.Run("test case 3", func(t *testing.T) {
		arr := []int{3, 8, -10, 23, 19, -4, -14, 27}
		expected := [][]int{{-14, -10}, {19, 23}, {23, 27}}

		assert.Equal(t, expected, minimumAbsDifference(arr))
	})
}
