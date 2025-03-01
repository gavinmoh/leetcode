package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApplyOperations(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{1, 2, 2, 1, 1, 0}
		expected := []int{1, 4, 2, 0, 0, 0}

		assert.Equal(t, expected, applyOperations(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{0, 1}
		expected := []int{1, 0}

		assert.Equal(t, expected, applyOperations(nums))
	})
}
