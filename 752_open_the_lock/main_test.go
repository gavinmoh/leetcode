package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpenLock(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		deadends := []string{"0201", "0101", "0102", "1212", "2002"}
		target := "0202"
		expected := 6

		assert.Equal(t, expected, openLock(deadends, target))
	})

	t.Run("test case 2", func(t *testing.T) {
		deadends := []string{"8888"}
		target := "0009"
		expected := 1

		assert.Equal(t, expected, openLock(deadends, target))
	})

	t.Run("test case 3", func(t *testing.T) {
		deadends := []string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}
		target := "8888"
		expected := -1

		assert.Equal(t, expected, openLock(deadends, target))
	})

	t.Run("test case 4", func(t *testing.T) {
		deadends := []string{"0000"}
		target := "8888"
		expected := -1

		assert.Equal(t, expected, openLock(deadends, target))
	})

	t.Run("test case 5", func(t *testing.T) {
		deadends := []string{"0201", "0101", "0102", "1212", "2002"}
		target := "0000"
		expected := 0

		assert.Equal(t, expected, openLock(deadends, target))
	})
}
