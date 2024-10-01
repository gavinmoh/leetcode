package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanArrange(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5, 10, 6, 7, 8, 9}
		k := 5

		assert.True(t, canArrange(arr, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5, 6}
		k := 7

		assert.True(t, canArrange(arr, k))
	})

	t.Run("test case 3", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5, 6}
		k := 10

		assert.False(t, canArrange(arr, k))
	})
}
