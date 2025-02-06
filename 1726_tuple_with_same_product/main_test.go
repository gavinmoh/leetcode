package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTupleSameProduct(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{2, 3, 4, 6}
		expected := 8

		assert.Equal(t, expected, tupleSameProduct(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{1, 2, 4, 5, 10}
		expected := 16

		assert.Equal(t, expected, tupleSameProduct(nums))
	})
}
