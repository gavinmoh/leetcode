package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEliminateMaximum(t *testing.T) {
	t.Run("should return 3", func(t *testing.T) {
		dist := []int{1, 3, 4}
		speed := []int{1, 1, 1}
		expected := 3
		assert.Equal(t, expected, eliminateMaximum(dist, speed))
	})

	t.Run("should return 1", func(t *testing.T) {
		dist := []int{1, 1, 2, 3}
		speed := []int{1, 1, 1, 1}
		expected := 1
		assert.Equal(t, expected, eliminateMaximum(dist, speed))
	})

	t.Run("should return 1", func(t *testing.T) {
		dist := []int{3, 2, 4}
		speed := []int{5, 3, 2}
		expected := 1
		assert.Equal(t, expected, eliminateMaximum(dist, speed))
	})

	t.Run("should return 2", func(t *testing.T) {
		dist := []int{3, 5, 7, 4, 5}
		speed := []int{2, 3, 6, 3, 2}
		expected := 2
		assert.Equal(t, expected, eliminateMaximum(dist, speed))
	})

}
