package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMaximizedCapital(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		k := 2
		w := 0
		profits := []int{1, 2, 3}
		capital := []int{0, 1, 1}
		expected := 4

		assert.Equal(t, expected, findMaximizedCapital(k, w, profits, capital))
	})

	t.Run("test case 2", func(t *testing.T) {
		k := 3
		w := 0
		profits := []int{1, 2, 3}
		capital := []int{0, 1, 2}
		expected := 6

		assert.Equal(t, expected, findMaximizedCapital(k, w, profits, capital))
	})
}
