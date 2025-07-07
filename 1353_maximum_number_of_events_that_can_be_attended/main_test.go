package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxEvents(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		events := [][]int{{1, 2}, {2, 3}, {3, 4}}
		expected := 3

		assert.Equal(t, expected, maxEvents(events))
	})

	t.Run("test case 2", func(t *testing.T) {
		events := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 2}}
		expected := 4

		assert.Equal(t, expected, maxEvents(events))
	})

	t.Run("test case 3", func(t *testing.T) {
		events := [][]int{{1, 4}, {4, 4}, {2, 2}, {3, 4}, {1, 1}}
		expected := 4

		assert.Equal(t, expected, maxEvents(events))
	})
}
