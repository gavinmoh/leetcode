package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindLucky(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		arr := []int{2, 2, 3, 4}
		expected := 2

		assert.Equal(t, expected, findLucky(arr))
	})

	t.Run("test case 2", func(t *testing.T) {
		arr := []int{1, 2, 2, 3, 3, 3}
		expected := 3

		assert.Equal(t, expected, findLucky(arr))
	})

	t.Run("test case 3", func(t *testing.T) {
		arr := []int{2, 2, 2, 3, 3}
		expected := -1

		assert.Equal(t, expected, findLucky(arr))
	})
}
