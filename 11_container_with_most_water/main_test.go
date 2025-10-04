package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxArea(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
		assert.Equal(t, 49, maxArea(height))
	})

	t.Run("test case 2", func(t *testing.T) {
		height := []int{1, 1}
		assert.Equal(t, 1, maxArea(height))
	})
}
