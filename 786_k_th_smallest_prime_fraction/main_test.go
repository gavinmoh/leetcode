package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKthSmallestPrimeFraction(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		arr := []int{1, 2, 3, 5}
		k := 3
		expected := []int{2, 5}

		assert.Equal(t, expected, kthSmallestPrimeFraction(arr, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		arr := []int{1, 7}
		k := 1
		expected := []int{1, 7}

		assert.Equal(t, expected, kthSmallestPrimeFraction(arr, k))
	})
}
