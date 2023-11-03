package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildArray(t *testing.T) {
	t.Run("should return [Push, Push, Pop, Push] when target is [1, 3] and n is 3", func(t *testing.T) {
		target := []int{1, 3}
		n := 3
		expected := []string{"Push", "Push", "Pop", "Push"}

		assert.Equal(t, expected, buildArray(target, n))
	})

	t.Run("should return [Push, Push, Push] when target is [1, 2, 3] and n is 3", func(t *testing.T) {
		target := []int{1, 2, 3}
		n := 3
		expected := []string{"Push", "Push", "Push"}

		assert.Equal(t, expected, buildArray(target, n))
	})

	t.Run("should return [Push, Push] when target is [1, 2] and n is 4", func(t *testing.T) {
		target := []int{1, 2}
		n := 4
		expected := []string{"Push", "Push"}

		assert.Equal(t, expected, buildArray(target, n))
	})
}
