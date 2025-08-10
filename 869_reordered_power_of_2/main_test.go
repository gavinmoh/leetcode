package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReorderedPowerOf2(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		n := 1

		assert.True(t, reorderedPowerOf2(n))
	})

	t.Run("test case 2", func(t *testing.T) {
		n := 10

		assert.False(t, reorderedPowerOf2(n))
	})
}
