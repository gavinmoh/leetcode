package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountGoodNumbers(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		n := int64(1)
		expected := 5

		assert.Equal(t, expected, countGoodNumbers(n))
	})

	t.Run("test case 2", func(t *testing.T) {
		n := int64(4)
		expected := 400

		assert.Equal(t, expected, countGoodNumbers(n))
	})

	t.Run("test case 1", func(t *testing.T) {
		n := int64(50)
		expected := 564908303

		assert.Equal(t, expected, countGoodNumbers(n))
	})
}
