package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumDeleteSum(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		s1 := "sea"
		s2 := "eat"
		expected := 231

		assert.Equal(t, expected, minimumDeleteSum(s1, s2))
	})

	t.Run("test case 2", func(t *testing.T) {
		s1 := "delete"
		s2 := "leet"
		expected := 403

		assert.Equal(t, expected, minimumDeleteSum(s1, s2))
	})
}
