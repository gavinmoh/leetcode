package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstUniqChar(t *testing.T) {
	t.Run("should return 0", func(t *testing.T) {
		s := "leetcode"
		expected := 0

		assert.Equal(t, expected, firstUniqChar(s))
	})

	t.Run("should return 2", func(t *testing.T) {
		s := "loveleetcode"
		expected := 2

		assert.Equal(t, expected, firstUniqChar(s))
	})

	t.Run("should return -1", func(t *testing.T) {
		s := "aabb"
		expected := -1

		assert.Equal(t, expected, firstUniqChar(s))
	})
}
