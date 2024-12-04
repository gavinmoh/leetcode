package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsArraySpecial(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{3, 4, 1, 2, 6}
		queries := [][]int{{0, 4}}
		expected := []bool{false}

		assert.Equal(t, expected, isArraySpecial(nums, queries))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{4, 3, 1, 6}
		queries := [][]int{{0, 2}, {2, 3}}
		expected := []bool{false, true}

		assert.Equal(t, expected, isArraySpecial(nums, queries))
	})
}
