package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRob(t *testing.T) {
	t.Run("should return 4", func(t *testing.T) {
		nums := []int{1, 2, 3, 1}
		expected := 4
		assert.Equal(t, expected, rob(nums))
	})

	t.Run("should return 12", func(t *testing.T) {
		nums := []int{2, 7, 9, 3, 1}
		expected := 12
		assert.Equal(t, expected, rob(nums))
	})
}
