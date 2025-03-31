package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPutMarbles(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		weights := []int{1, 3, 5, 1}
		k := 2
		expected := int64(4)

		assert.Equal(t, expected, putMarbles(weights, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		weights := []int{1, 3}
		k := 2
		expected := int64(0)

		assert.Equal(t, expected, putMarbles(weights, k))
	})
}
