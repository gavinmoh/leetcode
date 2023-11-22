package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindDiagonalOrder(t *testing.T) {
	t.Run("should return [1,4,2,7,5,3,8,6,9]", func(t *testing.T) {
		nums := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
		expected := []int{1, 4, 2, 7, 5, 3, 8, 6, 9}

		assert.Equal(t, expected, findDiagonalOrder(nums))
	})

	t.Run("should return [1,6,2,8,7,3,9,4,12,10,5,13,11,14,15,16]", func(t *testing.T) {
		nums := [][]int{{1, 2, 3, 4, 5}, {6, 7}, {8}, {9, 10, 11}, {12, 13, 14, 15, 16}}
		expected := []int{1, 6, 2, 8, 7, 3, 9, 4, 12, 10, 5, 13, 11, 14, 15, 16}

		assert.Equal(t, expected, findDiagonalOrder(nums))
	})
}
