package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindNumbers(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{12, 345, 2, 6, 7896}
		expected := 2

		assert.Equal(t, expected, findNumbers(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{555, 901, 482, 1171}
		expected := 1

		assert.Equal(t, expected, findNumbers(nums))
	})
}
