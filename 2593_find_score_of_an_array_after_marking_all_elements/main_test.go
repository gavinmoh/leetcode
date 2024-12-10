package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindScore(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{2, 1, 3, 4, 5, 2}
		expected := int64(7)

		assert.Equal(t, expected, findScore(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{2, 3, 5, 1, 3, 2}
		expected := int64(5)

		assert.Equal(t, expected, findScore(nums))
	})

}
