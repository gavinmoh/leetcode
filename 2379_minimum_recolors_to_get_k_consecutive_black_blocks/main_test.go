package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumRecolors(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		blocks := "WBBWWBBWBW"
		k := 7
		expected := 3

		assert.Equal(t, expected, minimumRecolors(blocks, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		blocks := "WBWBBBW"
		k := 2
		expected := 0

		assert.Equal(t, expected, minimumRecolors(blocks, k))
	})
}
