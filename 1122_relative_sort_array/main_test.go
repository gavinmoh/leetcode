package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRelativeSortArray(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		arr1 := []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}
		arr2 := []int{2, 1, 4, 3, 9, 6}
		expected := []int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19}

		assert.Equal(t, expected, relativeSortArray(arr1, arr2))
	})

	t.Run("test case 2", func(t *testing.T) {
		arr1 := []int{28, 6, 22, 8, 44, 17}
		arr2 := []int{22, 28, 8, 6}
		expected := []int{22, 28, 8, 6, 17, 44}

		assert.Equal(t, expected, relativeSortArray(arr1, arr2))
	})
}
