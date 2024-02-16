package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindLeastNumOfUniqueInts(t *testing.T) {
	t.Run("should return 1", func(t *testing.T) {
		arr := []int{5, 5, 4}
		k := 1
		expected := 1

		assert.Equal(t, expected, findLeastNumOfUniqueInts(arr, k))
	})

	t.Run("should return 2", func(t *testing.T) {
		arr := []int{4, 3, 1, 1, 3, 3, 2}
		k := 3
		expected := 2

		assert.Equal(t, expected, findLeastNumOfUniqueInts(arr, k))
	})
}
