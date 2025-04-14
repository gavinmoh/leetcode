package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountGoodTriplets(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		arr := []int{3, 0, 1, 1, 9, 7}
		a := 7
		b := 2
		c := 3
		expected := 4

		assert.Equal(t, expected, countGoodTriplets(arr, a, b, c))
	})

	t.Run("test case 2", func(t *testing.T) {
		arr := []int{1, 1, 2, 2, 3}
		a := 0
		b := 0
		c := 1
		expected := 0

		assert.Equal(t, expected, countGoodTriplets(arr, a, b, c))
	})
}
