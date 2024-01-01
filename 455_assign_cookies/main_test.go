package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindContentChildren(t *testing.T) {
	t.Run("should return 1", func(t *testing.T) {
		g := []int{1, 2, 3}
		s := []int{1, 1}
		expected := 1
		assert.Equal(t, expected, findContentChildren(g, s))
	})

	t.Run("should return 2", func(t *testing.T) {
		g := []int{1, 2}
		s := []int{1, 2, 3}
		expected := 2
		assert.Equal(t, expected, findContentChildren(g, s))
	})
}
