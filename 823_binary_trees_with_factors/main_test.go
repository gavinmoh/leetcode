package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumFactoredBinaryTrees(t *testing.T) {
	t.Run("should return 3", func(t *testing.T) {
		arr := []int{2, 4}
		assert.Equal(t, 3, numFactoredBinaryTrees(arr))
	})

	t.Run("should return 7", func(t *testing.T) {
		arr := []int{2, 4, 5, 10}
		assert.Equal(t, 7, numFactoredBinaryTrees(arr))
	})

	t.Run("should return 12", func(t *testing.T) {
		arr := []int{18, 3, 6, 2}
		assert.Equal(t, 12, numFactoredBinaryTrees(arr))
	})
}
