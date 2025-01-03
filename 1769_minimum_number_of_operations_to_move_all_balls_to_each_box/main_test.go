package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinOperations(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		boxes := "110"
		expected := []int{1, 1, 3}

		assert.Equal(t, expected, minOperations(boxes))
	})

	t.Run("test case 2", func(t *testing.T) {
		boxes := "001011"
		expected := []int{11, 8, 5, 4, 3, 4}

		assert.Equal(t, expected, minOperations(boxes))
	})
}
