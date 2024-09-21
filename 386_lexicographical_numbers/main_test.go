package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLexicalOrder(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		n := 13
		expected := []int{1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9}

		assert.Equal(t, expected, lexicalOrder(n))
	})

	t.Run("test case 2", func(t *testing.T) {
		n := 2
		expected := []int{1, 2}

		assert.Equal(t, expected, lexicalOrder(n))
	})

}
