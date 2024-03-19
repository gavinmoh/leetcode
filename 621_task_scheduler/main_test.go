package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeastInterval(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		tasks := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
		n := 2
		expected := 8

		assert.Equal(t, expected, leastInterval(tasks, n))
	})

	t.Run("test case 2", func(t *testing.T) {
		tasks := []byte{'A', 'C', 'A', 'B', 'D', 'B'}
		n := 1
		expected := 6

		assert.Equal(t, expected, leastInterval(tasks, n))
	})

	t.Run("test case 3", func(t *testing.T) {
		tasks := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
		n := 3
		expected := 10

		assert.Equal(t, expected, leastInterval(tasks, n))
	})
}
