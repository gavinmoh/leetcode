package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildArray(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{0, 2, 1, 5, 3, 4}
		expected := []int{0, 1, 2, 4, 5, 3}

		assert.Equal(t, expected, buildArray(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{5, 0, 1, 2, 3, 4}
		expected := []int{4, 5, 0, 1, 2, 3}

		assert.Equal(t, expected, buildArray(nums))
	})
}
