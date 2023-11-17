package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinPairSum(t *testing.T) {
	t.Run("should return 7", func(t *testing.T) {
		nums := []int{3, 5, 2, 3}
		expected := 7
		assert.Equal(t, expected, minPairSum(nums))
	})

	t.Run("should return 8", func(t *testing.T) {
		nums := []int{3, 5, 4, 2, 4, 6}
		expected := 8
		assert.Equal(t, expected, minPairSum(nums))
	})
}
