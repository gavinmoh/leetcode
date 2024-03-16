package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMaxLength(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{0, 1}
		expected := 2

		assert.Equal(t, expected, findMaxLength(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{0, 1, 0}
		expected := 2

		assert.Equal(t, expected, findMaxLength(nums))
	})
}
