package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindXSum(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{1, 1, 2, 2, 3, 4, 2, 3}
		k := 6
		x := 2
		expected := []int{6, 10, 12}

		assert.Equal(t, expected, findXSum(nums, k, x))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{3, 8, 7, 8, 7, 5}
		k := 2
		x := 2
		expected := []int{11, 15, 15, 15, 12}

		assert.Equal(t, expected, findXSum(nums, k, x))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{9, 2, 2}
		k := 3
		x := 3
		expected := []int{13}

		assert.Equal(t, expected, findXSum(nums, k, x))
	})
}
