package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxTaskAssign(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		tasks := []int{3, 2, 1}
		workers := []int{0, 3, 3}
		pills := 1
		strength := 1
		expected := 3

		assert.Equal(t, expected, maxTaskAssign(tasks, workers, pills, strength))
	})

	t.Run("test case 2", func(t *testing.T) {
		tasks := []int{5, 4}
		workers := []int{0, 0, 0}
		pills := 1
		strength := 5
		expected := 1

		assert.Equal(t, expected, maxTaskAssign(tasks, workers, pills, strength))
	})

	t.Run("test case 3", func(t *testing.T) {
		tasks := []int{10, 15, 30}
		workers := []int{0, 10, 10, 10, 10}
		pills := 3
		strength := 10
		expected := 2

		assert.Equal(t, expected, maxTaskAssign(tasks, workers, pills, strength))
	})

	t.Run("test case 4", func(t *testing.T) {
		tasks := []int{5, 9, 8, 5, 9}
		workers := []int{1, 6, 4, 2, 6}
		pills := 1
		strength := 5
		expected := 3

		assert.Equal(t, expected, maxTaskAssign(tasks, workers, pills, strength))
	})
}
