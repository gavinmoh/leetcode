package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumFourDivisors(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{21, 4, 7}
		expected := 32

		assert.Equal(t, expected, sumFourDivisors(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{21, 21}
		expected := 64

		assert.Equal(t, expected, sumFourDivisors(nums))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		expected := 0

		assert.Equal(t, expected, sumFourDivisors(nums))
	})
}
