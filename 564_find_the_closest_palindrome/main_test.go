package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNearestPalindromic(t *testing.T) {
	// t.Run("test case 1", func(t *testing.T) {
	// 	n := "123"
	// 	expected := "121"

	// 	assert.Equal(t, expected, nearestPalindromic(n))
	// })

	t.Run("test case 2", func(t *testing.T) {
		n := "1"
		expected := "0"

		assert.Equal(t, expected, nearestPalindromic(n))
	})

	t.Run("test case 3", func(t *testing.T) {
		n := "10"
		expected := "9"

		assert.Equal(t, expected, nearestPalindromic(n))
	})

	t.Run("test case 3", func(t *testing.T) {
		n := "11"
		expected := "9"

		assert.Equal(t, expected, nearestPalindromic(n))
	})

	t.Run("test case 5", func(t *testing.T) {
		n := "807045053224792883"
		expected := "807045053350540708"

		assert.Equal(t, expected, nearestPalindromic(n))
	})
}
