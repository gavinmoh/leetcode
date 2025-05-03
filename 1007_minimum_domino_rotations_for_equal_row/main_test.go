package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinDominoRotations(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		tops := []int{2, 1, 2, 4, 2, 2}
		bottoms := []int{5, 2, 6, 2, 3, 2}
		expected := 2

		assert.Equal(t, expected, minDominoRotations(tops, bottoms))
	})

	t.Run("test case 2", func(t *testing.T) {
		tops := []int{3, 5, 1, 2, 3}
		bottoms := []int{3, 6, 3, 3, 4}
		expected := -1

		assert.Equal(t, expected, minDominoRotations(tops, bottoms))
	})
}
