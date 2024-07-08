package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindTheWinner(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		n := 5
		k := 2
		expected := 3

		assert.Equal(t, expected, findTheWinner(n, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		n := 6
		k := 5
		expected := 1

		assert.Equal(t, expected, findTheWinner(n, k))
	})
}
