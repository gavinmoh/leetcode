package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNthUglyNumber(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		n := 10
		expected := 12

		assert.Equal(t, expected, nthUglyNumber(n))
	})

	t.Run("test case 2", func(t *testing.T) {
		n := 1
		expected := 1

		assert.Equal(t, expected, nthUglyNumber(n))
	})
}
