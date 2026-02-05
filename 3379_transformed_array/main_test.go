package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructTransformedArray(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{3, -2, 1, 1}
		expected := []int{1, 1, 1, 3}

		assert.Equal(t, expected, constructTransformedArray(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{-1, 4, -1}
		expected := []int{-1, -1, 4}

		assert.Equal(t, expected, constructTransformedArray(nums))
	})
}
