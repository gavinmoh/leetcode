package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClimbStairs(t *testing.T) {
	t.Run("should return 2", func(t *testing.T) {
		n := 2
		expected := 2
		assert.Equal(t, expected, climbStairs(n))
	})

	t.Run("should return ", func(t *testing.T) {
		n := 3
		expected := 3
		assert.Equal(t, expected, climbStairs(n))
	})
}
