package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountValidSelections(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{1, 0, 2, 0, 3}
		expected := 2

		assert.Equal(t, expected, countValidSelections(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{2, 3, 4, 0, 4, 1, 0}
		expected := 0

		assert.Equal(t, expected, countValidSelections(nums))
	})
}
