package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortPeople(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		names := []string{"Mary", "John", "Emma"}
		heights := []int{180, 165, 170}
		expected := []string{"Mary", "Emma", "John"}

		assert.Equal(t, expected, sortPeople(names, heights))
	})

	t.Run("test case 2", func(t *testing.T) {
		names := []string{"Alice", "Bob", "Bob"}
		heights := []int{155, 185, 150}
		expected := []string{"Bob", "Alice", "Bob"}

		assert.Equal(t, expected, sortPeople(names, heights))
	})
}
