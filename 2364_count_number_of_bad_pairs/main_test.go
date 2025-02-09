package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountBadPairs(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{4, 1, 3, 3}
		expected := int64(5)

		assert.Equal(t, expected, countBadPairs(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		expected := int64(0)

		assert.Equal(t, expected, countBadPairs(nums))
	})

}
