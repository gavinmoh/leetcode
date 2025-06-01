package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistributeCandies(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		n := 5
		limit := 2
		expected := int64(3)

		assert.Equal(t, expected, distributeCandies(n, limit))
	})

	t.Run("test case 2", func(t *testing.T) {
		n := 3
		limit := 3
		expected := int64(10)

		assert.Equal(t, expected, distributeCandies(n, limit))
	})
}
