package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSatisfied(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		customers := []int{1, 0, 1, 2, 1, 1, 7, 5}
		grumpy := []int{0, 1, 0, 1, 0, 1, 0, 1}
		minutes := 3
		expected := 16

		assert.Equal(t, expected, maxSatisfied(customers, grumpy, minutes))
	})

	t.Run("test case 2", func(t *testing.T) {
		customers := []int{1}
		grumpy := []int{0}
		minutes := 1
		expected := 1

		assert.Equal(t, expected, maxSatisfied(customers, grumpy, minutes))
	})

	t.Run("test case 3", func(t *testing.T) {
		customers := []int{4, 10, 10}
		grumpy := []int{1, 1, 0}
		minutes := 2
		expected := 24

		assert.Equal(t, expected, maxSatisfied(customers, grumpy, minutes))
	})
}
