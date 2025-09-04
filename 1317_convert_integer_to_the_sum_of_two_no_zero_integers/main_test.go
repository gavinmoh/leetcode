package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNoZeroIntegers(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		n := 2
		expected := []int{1, 1}

		assert.Equal(t, expected, getNoZeroIntegers(n))
	})

	t.Run("test case 2", func(t *testing.T) {
		n := 11
		expected := []int{2, 9}

		assert.Equal(t, expected, getNoZeroIntegers(n))
	})
}
