package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKthDistinct(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		arr := []string{"d", "b", "c", "b", "c", "a"}
		k := 2
		expected := "a"

		assert.Equal(t, expected, kthDistinct(arr, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		arr := []string{"aaa", "aa", "a"}
		k := 1
		expected := "aaa"

		assert.Equal(t, expected, kthDistinct(arr, k))
	})

	t.Run("test case 3", func(t *testing.T) {
		arr := []string{"a", "b", "a"}
		k := 3
		expected := ""

		assert.Equal(t, expected, kthDistinct(arr, k))
	})

}
