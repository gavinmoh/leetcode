package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartition(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s := "aab"
		expected := [][]string{{"a", "a", "b"}, {"aa", "b"}}

		assert.Equal(t, expected, partition(s))
	})

	t.Run("test case 2", func(t *testing.T) {
		s := "a"
		expected := [][]string{{"a"}}

		assert.Equal(t, expected, partition(s))
	})

	t.Run("test case 3", func(t *testing.T) {
		s := "bb"
		expected := [][]string{{"b", "b"}, {"bb"}}

		assert.Equal(t, expected, partition(s))
	})
}
