package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinDeletionSize(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		strs := []string{"cba", "daf", "ghi"}
		expected := 1

		assert.Equal(t, expected, minDeletionSize(strs))
	})

	t.Run("test case 2", func(t *testing.T) {
		strs := []string{"a", "b"}
		expected := 0

		assert.Equal(t, expected, minDeletionSize(strs))
	})

	t.Run("test case 3", func(t *testing.T) {
		strs := []string{"zyx", "wvu", "tsr"}
		expected := 3

		assert.Equal(t, expected, minDeletionSize(strs))
	})
}
