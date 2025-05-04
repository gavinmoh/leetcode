package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumEquivDominoPairs(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		dominoes := [][]int{{1, 2}, {2, 1}, {3, 4}, {5, 6}}
		expected := 1

		assert.Equal(t, expected, numEquivDominoPairs(dominoes))
	})

	t.Run("test case 2", func(t *testing.T) {
		dominoes := [][]int{{1, 2}, {1, 2}, {1, 1}, {1, 2}, {2, 2}}
		expected := 3

		assert.Equal(t, expected, numEquivDominoPairs(dominoes))
	})
}
