package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxDistance(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		position := []int{1, 2, 3, 4, 7}
		m := 3
		expected := 3

		assert.Equal(t, expected, maxDistance(position, m))
	})

	t.Run("test case 2", func(t *testing.T) {
		position := []int{5, 4, 3, 2, 1, 1000000000}
		m := 2
		expected := 999999999

		assert.Equal(t, expected, maxDistance(position, m))
	})
}
