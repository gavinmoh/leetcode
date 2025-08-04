package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTotalFruit(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		fruits := []int{1, 2, 1}
		expected := 3

		assert.Equal(t, expected, totalFruit(fruits))
	})

	t.Run("test case 2", func(t *testing.T) {
		fruits := []int{0, 1, 2, 2}
		expected := 3

		assert.Equal(t, expected, totalFruit(fruits))
	})

	t.Run("test case 3", func(t *testing.T) {
		fruits := []int{1, 2, 3, 2, 2}
		expected := 4

		assert.Equal(t, expected, totalFruit(fruits))
	})

	t.Run("test case 4", func(t *testing.T) {
		fruits := []int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}
		expected := 5

		assert.Equal(t, expected, totalFruit(fruits))
	})
}
