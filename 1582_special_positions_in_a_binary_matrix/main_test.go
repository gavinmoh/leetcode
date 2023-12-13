package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumSpecial(t *testing.T) {
	t.Run("should return 1", func(t *testing.T) {
		mat := [][]int{{1, 0, 0}, {0, 0, 1}, {1, 0, 0}}
		expected := 1
		assert.Equal(t, expected, numSpecial(mat))
	})

	// t.Run("should return 1", func(t *testing.T) {
	// 	mat := [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	// 	expected := 3
	// 	assert.Equal(t, expected, numSpecial(mat))
	// })
}
