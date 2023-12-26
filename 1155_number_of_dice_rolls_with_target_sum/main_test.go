package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumRollsToTarget(t *testing.T) {
	t.Run("should return 1", func(t *testing.T) {
		n := 1
		k := 6
		target := 3
		expected := 1
		assert.Equal(t, expected, numRollsToTarget(n, k, target))
	})

	t.Run("should return 6", func(t *testing.T) {
		n := 2
		k := 6
		target := 7
		expected := 6
		assert.Equal(t, expected, numRollsToTarget(n, k, target))
	})

	t.Run("should return 222616187", func(t *testing.T) {
		n := 30
		k := 30
		target := 500
		expected := 222616187
		assert.Equal(t, expected, numRollsToTarget(n, k, target))
	})

}
