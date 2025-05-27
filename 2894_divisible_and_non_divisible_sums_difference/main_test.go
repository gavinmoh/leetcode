package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDifferenceOfSums(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		n := 10
		m := 3
		expected := 19

		assert.Equal(t, expected, differenceOfSums(n, m))
	})

	t.Run("test case 2", func(t *testing.T) {
		n := 5
		m := 6
		expected := 15

		assert.Equal(t, expected, differenceOfSums(n, m))
	})

	t.Run("test case 3", func(t *testing.T) {
		n := 5
		m := 1
		expected := -15

		assert.Equal(t, expected, differenceOfSums(n, m))
	})
}
