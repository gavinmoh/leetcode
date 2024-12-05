package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxCount(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		banned := []int{1, 6, 5}
		n := 5
		maxSum := 6
		expected := 2

		assert.Equal(t, expected, maxCount(banned, n, maxSum))
	})

	t.Run("test case 2", func(t *testing.T) {
		banned := []int{1, 2, 3, 4, 5, 6, 7}
		n := 8
		maxSum := 1
		expected := 0

		assert.Equal(t, expected, maxCount(banned, n, maxSum))
	})

	t.Run("test case 3", func(t *testing.T) {
		banned := []int{11}
		n := 7
		maxSum := 50
		expected := 7

		assert.Equal(t, expected, maxCount(banned, n, maxSum))
	})

	t.Run("test case 4", func(t *testing.T) {
		banned := []int{176, 36, 104, 125, 188, 152, 101, 47, 51, 65, 39, 174, 29, 55, 13, 138, 79, 81, 175, 178, 42, 108, 24, 80, 183, 190, 123, 20, 139, 22, 140, 62, 58, 137, 68, 148, 172, 76, 173, 189, 151, 186, 153, 57, 142, 105, 133, 114, 165, 118, 56, 59, 124, 82, 49, 94, 8, 146, 109, 14, 85, 44, 60, 181, 95, 23, 150, 97, 28, 182, 157, 46, 160, 155, 12, 67, 135, 117, 2, 25, 74, 91, 71, 98, 127, 120, 130, 107, 168, 18, 69, 110, 61, 147, 145, 38}
		n := 3016
		maxSum := 240
		expected := 17

		assert.Equal(t, expected, maxCount(banned, n, maxSum))
	})
}
