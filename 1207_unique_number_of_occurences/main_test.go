package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniqueOccurrences(t *testing.T) {
	t.Run("should return true", func(t *testing.T) {
		arr := []int{1, 2, 2, 1, 1, 3}
		assert.True(t, uniqueOccurrences(arr))
	})

	t.Run("should return false", func(t *testing.T) {
		arr := []int{1, 2}
		assert.False(t, uniqueOccurrences(arr))
	})

	t.Run("should return true", func(t *testing.T) {
		arr := []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}
		assert.True(t, uniqueOccurrences(arr))
	})

}
