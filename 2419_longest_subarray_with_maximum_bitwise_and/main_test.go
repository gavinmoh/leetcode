package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestSubarray(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{1, 2, 3, 3, 2, 2}
		expected := 2

		assert.Equal(t, expected, longestSubarray(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{1, 2, 3, 4}
		expected := 1

		assert.Equal(t, expected, longestSubarray(nums))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 395808, 470030, 470030, 470030, 470030, 470030, 470030, 470030, 470030, 470030, 470030, 470030, 470030, 470030, 470030, 470030, 470030, 470030, 470030, 470030, 470030, 470030, 470030, 470030, 470030, 153490, 330001, 330001, 330001, 330001, 330001, 330001, 330001, 37284, 470030, 470030, 470030, 470030, 470030, 470030, 156542, 226743}
		expected := 24

		assert.Equal(t, expected, longestSubarray(nums))
	})
}
