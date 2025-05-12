package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindEvenNumbers(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		digits := []int{2, 1, 3, 0}
		expected := []int{102, 120, 130, 132, 210, 230, 302, 310, 312, 320}

		assert.Equal(t, expected, findEvenNumbers(digits))
	})

	t.Run("test case 2", func(t *testing.T) {
		digits := []int{2, 2, 8, 8, 2}
		expected := []int{222, 228, 282, 288, 822, 828, 882}

		assert.Equal(t, expected, findEvenNumbers(digits))
	})

	t.Run("test case 3", func(t *testing.T) {
		digits := []int{3, 7, 5}
		expected := []int{}

		assert.Equal(t, expected, findEvenNumbers(digits))
	})
}
