package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinEnd(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		x := 4
		n := 3
		expected := int64(6)

		assert.Equal(t, expected, minEnd(n, x))
	})

	t.Run("test case 2", func(t *testing.T) {
		x := 7
		n := 2
		expected := int64(15)

		assert.Equal(t, expected, minEnd(n, x))
	})
}
