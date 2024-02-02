package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSequentialDigits(t *testing.T) {
	t.Run("should return [123,234]", func(t *testing.T) {
		low := 100
		high := 300
		expected := []int{123, 234}

		assert.Equal(t, expected, sequentialDigits(low, high))
	})

	t.Run("should return [1234,2345,3456,4567,5678,6789]", func(t *testing.T) {
		low := 1000
		high := 10000
		expected := []int{1234, 2345, 3456, 4567, 5678, 6789}

		assert.Equal(t, expected, sequentialDigits(low, high))
	})
}
