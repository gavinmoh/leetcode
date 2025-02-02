package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheck(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{3, 4, 5, 1, 2}

		assert.True(t, check(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{2, 1, 3, 4}

		assert.False(t, check(nums))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{1, 2, 3}

		assert.True(t, check(nums))
	})
}
