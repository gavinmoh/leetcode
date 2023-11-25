package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSumAbsoluteDifferences(t *testing.T) {
	t.Run("should return [4,3,5]", func(t *testing.T) {
		nums := []int{2, 3, 5}
		expected := []int{4, 3, 5}
		assert.Equal(t, expected, getSumAbsoluteDifferences(nums))
	})

	t.Run("should return [24,15,13,15,21]", func(t *testing.T) {
		nums := []int{1, 4, 6, 8, 10}
		expected := []int{24, 15, 13, 15, 21}
		assert.Equal(t, expected, getSumAbsoluteDifferences(nums))
	})
}
