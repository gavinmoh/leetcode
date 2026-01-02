package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepeatedNTimes(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{1, 2, 3, 3}
		expected := 3

		assert.Equal(t, expected, repeatedNTimes(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{2, 1, 2, 5, 3, 2}
		expected := 2

		assert.Equal(t, expected, repeatedNTimes(nums))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{5, 1, 5, 2, 5, 3, 5, 4}
		expected := 5

		assert.Equal(t, expected, repeatedNTimes(nums))
	})
}
