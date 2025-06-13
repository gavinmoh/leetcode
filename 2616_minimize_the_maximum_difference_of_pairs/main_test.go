package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimizeMax(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{10, 1, 2, 7, 1, 3}
		p := 2
		expected := 1

		assert.Equal(t, expected, minimizeMax(nums, p))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{4, 2, 1, 2}
		p := 1
		expected := 0

		assert.Equal(t, expected, minimizeMax(nums, p))
	})
}
