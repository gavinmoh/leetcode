package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAverageWaitingTime(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		customers := [][]int{{1, 2}, {2, 5}, {4, 3}}
		expected := 5.0

		assert.Equal(t, expected, averageWaitingTime(customers))
	})

	t.Run("test case 2", func(t *testing.T) {
		customers := [][]int{{5, 2}, {5, 4}, {10, 3}, {20, 1}}
		expected := 3.25

		assert.Equal(t, expected, averageWaitingTime(customers))
	})
}
