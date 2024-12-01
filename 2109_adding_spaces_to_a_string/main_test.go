package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddSpaces(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s := "LeetcodeHelpsMeLearn"
		spaces := []int{8, 13, 15}
		expected := "Leetcode Helps Me Learn"

		assert.Equal(t, expected, addSpaces(s, spaces))
	})

	t.Run("test case 2", func(t *testing.T) {
		s := "icodeinpython"
		spaces := []int{1, 5, 7, 9}
		expected := "i code in py thon"

		assert.Equal(t, expected, addSpaces(s, spaces))
	})

	t.Run("test case 3", func(t *testing.T) {
		s := "spacing"
		spaces := []int{0, 1, 2, 3, 4, 5, 6}
		expected := " s p a c i n g"

		assert.Equal(t, expected, addSpaces(s, spaces))
	})
}
