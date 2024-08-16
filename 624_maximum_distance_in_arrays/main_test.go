package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxDistance(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		arrays := [][]int{{1, 2, 3}, {4, 5}, {1, 2, 3}}
		expected := 4

		assert.Equal(t, expected, maxDistance(arrays))
	})

	t.Run("test case 2", func(t *testing.T) {
		arrays := [][]int{{1}, {1}}
		expected := 0

		assert.Equal(t, expected, maxDistance(arrays))
	})
}
