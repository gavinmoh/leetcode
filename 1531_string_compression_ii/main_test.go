package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLengthOfOptimalCompression(t *testing.T) {
	t.Run("should return 4", func(t *testing.T) {
		s := "aaabcccd"
		k := 2
		expected := 4
		assert.Equal(t, expected, getLengthOfOptimalCompression(s, k))
	})

	t.Run("should return 2", func(t *testing.T) {
		s := "aabbaa"
		k := 2
		expected := 2
		assert.Equal(t, expected, getLengthOfOptimalCompression(s, k))
	})

	t.Run("should return 3", func(t *testing.T) {
		s := "aaaaaaaaaaa"
		k := 0
		expected := 3
		assert.Equal(t, expected, getLengthOfOptimalCompression(s, k))
	})

}
