package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumCost(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		source := "abcd"
		target := "acbe"
		original := []byte{'a', 'b', 'c', 'c', 'e', 'd'}
		changed := []byte{'b', 'c', 'b', 'e', 'b', 'e'}
		cost := []int{2, 5, 5, 1, 2, 20}
		expected := int64(28)

		assert.Equal(t, expected, minimumCost(source, target, original, changed, cost))
	})

	t.Run("test case 2", func(t *testing.T) {
		source := "aaaa"
		target := "bbbb"
		original := []byte{'a', 'c'}
		changed := []byte{'c', 'b'}
		cost := []int{1, 2}
		expected := int64(12)

		assert.Equal(t, expected, minimumCost(source, target, original, changed, cost))
	})

	t.Run("test case 3", func(t *testing.T) {
		source := "abcd"
		target := "abce"
		original := []byte{'a'}
		changed := []byte{'e'}
		cost := []int{10000}
		expected := int64(-1)

		assert.Equal(t, expected, minimumCost(source, target, original, changed, cost))
	})
}
