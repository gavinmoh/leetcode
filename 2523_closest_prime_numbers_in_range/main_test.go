package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClosestPrimes(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		left := 10
		right := 19
		expected := []int{11, 13}

		assert.Equal(t, expected, closestPrimes(left, right))
	})

	t.Run("test case 2", func(t *testing.T) {
		left := 4
		right := 6
		expected := []int{-1, -1}

		assert.Equal(t, expected, closestPrimes(left, right))
	})
}
