package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMatrix(t *testing.T) {
	t.Run("should return [[1,3,4,2],[1,3],[1]]", func(t *testing.T) {
		nums := []int{1, 3, 4, 1, 2, 3, 1}
		expected := [][]int{{1, 3, 4, 2}, {1, 3}, {1}}
		assert.Equal(t, expected, findMatrix(nums))
	})

}
