package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsTrionic(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{1, 3, 5, 4, 2, 6}

		assert.True(t, isTrionic(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{2, 1, 3}

		assert.False(t, isTrionic(nums))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{5, 9, 1, 7}

		assert.True(t, isTrionic(nums))
	})
}
