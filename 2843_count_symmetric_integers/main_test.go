package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountSymmetricIntegers(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		low := 1
		high := 100
		expected := 9

		assert.Equal(t, expected, countSymmetricIntegers(low, high))
	})

	t.Run("test case 2", func(t *testing.T) {
		low := 1200
		high := 1230
		expected := 4

		assert.Equal(t, expected, countSymmetricIntegers(low, high))
	})
}
