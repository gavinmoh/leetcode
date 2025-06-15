package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxDiff(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		num := 555
		expected := 888

		assert.Equal(t, expected, maxDiff(num))
	})

	t.Run("test case 2", func(t *testing.T) {
		num := 9
		expected := 8

		assert.Equal(t, expected, maxDiff(num))
	})

	t.Run("test case 3", func(t *testing.T) {
		num := 123456
		expected := 820000

		assert.Equal(t, expected, maxDiff(num))
	})
}
