package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivideArray(t *testing.T) {
	t.Run("should return [[1,1,3],[3,4,5],[7,8,9]]", func(t *testing.T) {
		nums := []int{1, 3, 4, 8, 7, 9, 3, 5, 1}
		k := 2
		expected := [][]int{{1, 1, 3}, {3, 4, 5}, {7, 8, 9}}
		assert.Equal(t, expected, divideArray(nums, k))
	})

	t.Run("should return empty array", func(t *testing.T) {
		nums := []int{1, 3, 3, 2, 7, 3}
		k := 3
		expected := [][]int(nil)
		assert.Equal(t, expected, divideArray(nums, k))
	})
}
