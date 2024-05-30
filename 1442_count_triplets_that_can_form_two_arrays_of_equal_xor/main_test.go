package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountTriplets(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		arr := []int{2, 3, 1, 6, 7}
		expected := 4

		assert.Equal(t, expected, countTriplets(arr))
	})

	t.Run("test case 2", func(t *testing.T) {
		arr := []int{1, 1, 1, 1, 1}
		expected := 10

		assert.Equal(t, expected, countTriplets(arr))
	})
}
