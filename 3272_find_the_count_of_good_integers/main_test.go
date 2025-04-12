package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountGoodIntegers(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		n := 3
		k := 5
		expected := int64(27)

		assert.Equal(t, expected, countGoodIntegers(n, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		n := 1
		k := 4
		expected := int64(2)

		assert.Equal(t, expected, countGoodIntegers(n, k))
	})

	t.Run("test case 3", func(t *testing.T) {
		n := 5
		k := 6
		expected := int64(2468)

		assert.Equal(t, expected, countGoodIntegers(n, k))
	})
}
