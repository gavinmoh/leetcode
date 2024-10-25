package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveSubfolders(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		folder := []string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"}
		expected := []string{"/a", "/c/d", "/c/f"}

		assert.Equal(t, expected, removeSubfolders(folder))
	})

	t.Run("test case 2", func(t *testing.T) {
		folder := []string{"/a", "/a/b/c", "/a/b/d"}
		expected := []string{"/a"}

		assert.Equal(t, expected, removeSubfolders(folder))
	})

	t.Run("test case 3", func(t *testing.T) {
		folder := []string{"/a/b/c", "/a/b/ca", "/a/b/d"}
		expected := []string{"/a/b/c", "/a/b/ca", "/a/b/d"}

		assert.Equal(t, expected, removeSubfolders(folder))
	})
}
