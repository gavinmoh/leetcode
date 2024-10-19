package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindKthBit(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		n := 3
		k := 1
		expected := byte('0')

		assert.Equal(t, expected, findKthBit(n, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		n := 4
		k := 11
		expected := byte('1')

		assert.Equal(t, expected, findKthBit(n, k))
	})
}
