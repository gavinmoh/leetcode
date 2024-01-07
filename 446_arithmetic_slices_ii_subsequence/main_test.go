package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberOfArithmeticSlices(t *testing.T) {
	t.Run("should return 7", func(t *testing.T) {
		nums := []int{2, 4, 6, 8, 10}
		expected := 7
		assert.Equal(t, expected, numberOfArithmeticSlices(nums))
	})

	t.Run("should return 16", func(t *testing.T) {
		nums := []int{7, 7, 7, 7, 7}
		expected := 16
		assert.Equal(t, expected, numberOfArithmeticSlices(nums))
	})
}
