package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountOdds(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		low := 3
		high := 7
		expected := 3

		assert.Equal(t, expected, countOdds(low, high))
	})

	t.Run("test case 2", func(t *testing.T) {
		low := 8
		high := 10
		expected := 1

		assert.Equal(t, expected, countOdds(low, high))
	})
}
