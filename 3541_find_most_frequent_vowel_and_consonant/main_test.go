package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxFreqSum(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s := "successes"
		expected := 6

		assert.Equal(t, expected, maxFreqSum(s))
	})

	t.Run("test case 2", func(t *testing.T) {
		s := "aeiaeia"
		expected := 3

		assert.Equal(t, expected, maxFreqSum(s))
	})
}
