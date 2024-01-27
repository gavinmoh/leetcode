package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKInversePairs(t *testing.T) {
	t.Run("should return 1", func(t *testing.T) {
		n := 3
		k := 0
		expected := 1

		assert.Equal(t, expected, kInversePairs(n, k))
	})

	t.Run("should return 2", func(t *testing.T) {
		n := 3
		k := 1
		expected := 2

		assert.Equal(t, expected, kInversePairs(n, k))
	})

	t.Run("should return 663677020", func(t *testing.T) {
		n := 1000
		k := 1000
		expected := 663677020

		assert.Equal(t, expected, kInversePairs(n, k))
	})

}
