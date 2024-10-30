package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimizedMaximum(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		n := 6
		quantities := []int{11, 6}
		expected := 3

		assert.Equal(t, expected, minimizedMaximum(n, quantities))
	})

	t.Run("test case 2", func(t *testing.T) {
		n := 7
		quantities := []int{15, 10, 10}
		expected := 5

		assert.Equal(t, expected, minimizedMaximum(n, quantities))
	})

	t.Run("test case 3", func(t *testing.T) {
		n := 1
		quantities := []int{100000}
		expected := 100000

		assert.Equal(t, expected, minimizedMaximum(n, quantities))
	})
}
