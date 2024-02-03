package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSumAfterPartitioning(t *testing.T) {
	t.Run("should return 84", func(t *testing.T) {
		arr := []int{1, 15, 7, 9, 2, 5, 10}
		k := 3
		expected := 84

		assert.Equal(t, expected, maxSumAfterPartitioning(arr, k))
	})

	t.Run("should return 83", func(t *testing.T) {
		arr := []int{1, 4, 1, 5, 7, 3, 6, 1, 9, 9, 3}
		k := 4
		expected := 83

		assert.Equal(t, expected, maxSumAfterPartitioning(arr, k))
	})

	t.Run("should return 1", func(t *testing.T) {
		arr := []int{1}
		k := 1
		expected := 1

		assert.Equal(t, expected, maxSumAfterPartitioning(arr, k))
	})
}
