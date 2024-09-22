package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindKthNumber(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		n := 13
		k := 2
		expected := 10

		assert.Equal(t, expected, findKthNumber(n, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		n := 1
		k := 1
		expected := 1

		assert.Equal(t, expected, findKthNumber(n, k))
	})

	t.Run("test case 3", func(t *testing.T) {
		n := 9885387
		k := 8786251
		expected := 8907623

		assert.Equal(t, expected, findKthNumber(n, k))
	})
}
